package auth

// ValidGroupsResponse : validGroups response
type ValidGroupsResponse struct {
	Status      int      `json:"status"`
	ValidGroups []string `json:"validGroups,omitempty"`
}

// HasPermissionResponse : haspermission response
type HasPermissionResponse struct {
	Status        int        `json:"status"`
	HasPermission bool       `json:"hasPermission"`
	ValidGroups   []string   `json:"validGroups,omitempty"`
	Group         *GroupInfo `json:"group,omitempty"`
}

// GroupInfo IAM group info
type GroupInfo struct {
	ID   string `json:"id,omitempty"`
	Code string `json:"code,omitempty"`
	Type string `json:"type,omitempty"`
}
