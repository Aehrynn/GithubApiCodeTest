package Structs

type Organisation struct {
	Login            string      `json:"login"`
	ID               int         `json:"id"`
	NodeID           string      `json:"node_id"`
	Name             string      `json:"name"`
	URL              string      `json:"url"`
	ReposURL         string      `json:"repos_url"`
	EventsURL        string      `json:"events_url"`
	HooksURL         string      `json:"hooks_url"`
	IssuesURL        string      `json:"issues_url"`
	MembersURL       string      `json:"members_url"`
	PublicMembersURL string      `json:"public_members_url"`
	AvatarURL        string      `json:"avatar_url"`
	Description      interface{} `json:"description"`
}
