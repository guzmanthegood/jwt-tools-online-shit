package auth

import (
	"net/http"

	"jwt-tools-online-shit/utils"

	"github.com/go-chi/render"
)

// GetValidGroups : return valid groups
func GetValidGroups(w http.ResponseWriter, r *http.Request) {
	a := NewAuth(r.Context())

	api := r.URL.Query().Get("api")
	rsc := r.URL.Query().Get("rsc")
	per := r.URL.Query().Get("per")

	if api == "" {
		render.Render(w, r, utils.EmptyField("api"))
		return
	}
	if rsc == "" {
		render.Render(w, r, utils.EmptyField("rsc"))
		return
	}
	if per == "" {
		render.Render(w, r, utils.EmptyField("per"))
		return
	}

	validGroups, err := a.GetAllValidGroups(api, rsc, per)
	if err != nil {
		render.Render(w, r, utils.BadRequest(err))
		return
	}

	render.JSON(w, r, ValidGroupsResponse{Status: 200, ValidGroups: validGroups})
}

// HasPermission : return true if has permission
func HasPermission(w http.ResponseWriter, r *http.Request) {
	a := NewAuth(r.Context())

	api := r.URL.Query().Get("api")
	rsc := r.URL.Query().Get("rsc")
	per := r.URL.Query().Get("per")
	grp := r.URL.Query().Get("grp")

	if api == "" {
		render.Render(w, r, utils.EmptyField("api"))
		return
	}
	if rsc == "" {
		render.Render(w, r, utils.EmptyField("rsc"))
		return
	}
	if per == "" {
		render.Render(w, r, utils.EmptyField("per"))
		return
	}
	if grp == "" {
		render.Render(w, r, utils.EmptyField("grp"))
		return
	}

	ok, validGroups, gr, err := a.HasPermission(api, rsc, per, grp)
	if err != nil {
		render.Render(w, r, utils.BadRequest(err))
		return
	}
	render.JSON(w, r, HasPermissionResponse{Status: 200, ValidGroups: validGroups, HasPermission: ok, Group: gr})
}
