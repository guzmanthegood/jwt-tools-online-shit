package apix

// Client interface
type Client interface {
	GetGroup(code string) (*Group, error)
}

// IamGroupResponseStruct : iam group response struct
// https://mholt.github.io/json-to-go/ generated
type IamGroupResponseStruct struct {
	Admin struct {
		Groups struct {
			Edges []struct {
				Node struct {
					GroupData Group `json:"groupData"`
				} `json:"node"`
			} `json:"edges"`
		} `json:"groups"`
	} `json:"admin"`
}

// Group : iam group struct
type Group struct {
	ID   string `json:"id"`
	Code string `json:"code"`
	Type string `json:"type"`
}
