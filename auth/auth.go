package auth

import (
	"context"
	"fmt"

	"jwt-tools-online-shit/apix"
	"jwt-tools-online-shit/utils"

	authorization "github.com/travelgateX/go-jwt-tools"
)

// Valid Permissions
var validPermissions = []string{"c", "r", "u", "d", "x", "a", "f"}

// Auth wrapper struct
type Auth struct {
	authorization.User
}

// NewAuth creates new authentication object
func NewAuth(ctx context.Context) Auth {
	user, _ := authorization.UserFromContext(ctx)
	return Auth{*user}
}

// GetAllValidGroups returns all groups of the cartesian product of all permissions
func (a *Auth) GetAllValidGroups(api string, rsc string, per string) (gr []string, err error) {
	// check if is valid permission
	if !utils.Contains(validPermissions, per) {
		return nil, fmt.Errorf("permission: '%v' its not valid. Valid options: %v", per, validPermissions)
	}

	// check all permissions
	for _, p := range permissionsCombination(api, rsc) {
		gr = utils.JoinArray(gr, a.validGroups(p[0], p[1], "a"))
		gr = utils.JoinArray(gr, a.validGroups(p[0], p[1], per))
	}
	return gr, nil
}

// HasPermission return true if bearer has permission in the searched group
// * asks group type to IAM *
func (a *Auth) HasPermission(api string, rsc string, per string, grp string) (bool, []string, *GroupInfo, error) {
	validGroups, err := a.GetAllValidGroups(api, rsc, per)
	if err != nil {
		return false, validGroups, nil, err
	}

	// recover IAM group info
	iam := apix.NewCachedClient(apix.NewDefaultClient())
	iamGroup, err := iam.GetGroup(grp)
	if err != nil {
		return false, validGroups, nil, err
	}

	// IamGroup To GroupInfo
	gr := GroupInfo{
		ID:   iamGroup.ID,
		Code: iamGroup.Code,
		Type: iamGroup.Type,
	}

	// returns true if specific group is found
	if utils.Contains(validGroups, grp) {
		return true, validGroups, &gr, nil
	}

	// return true if contains "all" and IAM group isn't ROOT or TEAM type
	if utils.Contains(validGroups, "all") {
		ok := (gr.Type != "ROOT" && gr.Type != "TEAM")
		return ok, validGroups, &gr, nil
	}

	return false, validGroups, &gr, nil
}

func (a *Auth) validGroups(product string, object string, per string) []string {
	g, _ := a.Permissions.CheckPermission(product, object, authorization.Permission(per))
	return g
}

func permissionsCombination(product, object string) [][]string {
	return [][]string{
		[]string{product, object},
		[]string{product, "all"},
		[]string{"all", object},
		[]string{"all", "all"},
	}
}
