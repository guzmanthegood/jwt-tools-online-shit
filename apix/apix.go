package apix

// apix is a client package that connects to the TravelgateX API
// using graphQL client to wrap with cache.

import (
	"context"
	"fmt"
	"log"

	"github.com/machinebox/graphql"
)

const (
	defaultURL    = "https://api.travelgatex.com/"
	defaultAPIKey = "bb7de66f-7a3c-464a-4f72-2cce1cd48205"
)

// NewClient returns new TravelgateX GateWay API client
func NewClient(url string, apiKey string) Client {
	cli := graphql.NewClient(url)
	cli.Log = func(s string) { log.Println(s) }
	return apiClient{cli, apiKey}
}

// NewDefaultClient return new api.travelgatex.com API client
func NewDefaultClient() Client {
	return NewClient(defaultURL, defaultAPIKey)
}

type apiClient struct {
	cli    *graphql.Client
	apiKey string
}

// GetGroup returns IAM group basic info
func (a apiClient) GetGroup(code string) (*Group, error) {
	req := graphql.NewRequest(`
		query($codes: [ID!]){
			admin{
				groups(codes:$codes){
					edges{
						node{
							groupData{
								id
								code
								type
							}
						}
					}
				}
			}
		}
	`)

	res := IamGroupResponseStruct{}
	req.Var("codes", []string{code})
	req.Header.Add("Authorization", "Apikey "+a.apiKey)

	ctx := context.Background()
	if err := a.cli.Run(ctx, req, &res); err != nil {
		return nil, err
	}
	if len(res.Admin.Groups.Edges) == 0 {
		return nil, fmt.Errorf("IAM Group '%v' not found", code)
	}

	return &res.Admin.Groups.Edges[0].Node.GroupData, nil
}
