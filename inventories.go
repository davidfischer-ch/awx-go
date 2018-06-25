package awx

import (
	"net/http"
	"time"
)

type InventoriesService struct {
	client *Client
}

type Related struct {
	CreatedBy                    string `json:"created_by"`
	ModifiedBy                   string `json:"modified_by"`
	JobTemplates                 string `json:"job_templates"`
	VariableData                 string `json:"variable_data"`
	RootGroups                   string `json:"root_groups"`
	ObjectRoles                  string `json:"object_roles"`
	AdHocCommands                string `json:"ad_hoc_commands"`
	Script                       string `json:"script"`
	Tree                         string `json:"tree"`
	AccessList                   string `json:"access_list"`
	ActivityStream               string `json:"activity_stream"`
	InstanceGroups               string `json:"instance_groups"`
	Hosts                        string `json:"hosts"`
	Groups                       string `json:"groups"`
	Copy                         string `json:"copy"`
	UpdateInventorySources       string `json:"update_inventory_sources"`
	InventorySources             string `json:"inventory_sources"`
	Organization                 string `json:"organization"`
	Labels                       string `json:"labels"`
	Inventory                    string `json:"inventory"`
	Project                      string `json:"project"`
	Credential                   string `json:"credential"`
	ExtraCredentials             string `json:"extra_credentials"`
	Credentials                  string `json:"credentials"`
	NotificationTemplatesError   string `json:"notification_templates_error"`
	NotificationTemplatesSuccess string `json:"notification_templates_success"`
	Jobs                         string `json:"jobs"`
	NotificationTemplatesAny     string `json:"notification_templates_any"`
	Launch                       string `json:"launch"`
	Schedules                    string `json:"schedules"`
	SurveySpec                   string `json:"survey_spec"`
}

type OrgnizationSummary struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ByUserSummary struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type ApplyRole struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ObjectRoles struct {
	UseRole     *ApplyRole `json:"use_role"`
	AdminRole   *ApplyRole `json:"admin_role"`
	AdhocRole   *ApplyRole `json:"adhoc_role"`
	UpdateRole  *ApplyRole `json:"update_role"`
	ReadRole    *ApplyRole `json:"read_role"`
	ExecuteRole *ApplyRole `json:"execute_role"`
}

type UserCapabilities struct {
	Edit     bool `json:"edit"`
	Start    bool `json:"start"`
	Schedule bool `json:"schedule"`
	Copy     bool `json:"copy"`
	Adhoc    bool `json:"adhoc"`
	Delete   bool `json:"delete"`
}

type Project struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
	ScmType     string `json:"scm_type"`
}

type Inventory struct {
	ID                           int         `json:"id"`
	Type                         string      `json:"type"`
	URL                          string      `json:"url"`
	Related                      *Related    `json:"related"`
	SummaryFields                *Summary    `json:"summary_fields"`
	Created                      time.Time   `json:"created"`
	Modified                     time.Time   `json:"modified"`
	Name                         string      `json:"name"`
	Description                  string      `json:"description"`
	Organization                 int         `json:"organization"`
	OrganizationID               int         `json:"organization_id"`
	Kind                         string      `json:"kind"`
	HostFilter                   interface{} `json:"host_filter"`
	Variables                    string      `json:"variables"`
	HasActiveFailures            bool        `json:"has_active_failures"`
	TotalHosts                   int         `json:"total_hosts"`
	HostsWithActiveFailures      int         `json:"hosts_with_active_failures"`
	TotalGroups                  int         `json:"total_groups"`
	GroupsWithActiveFailures     int         `json:"groups_with_active_failures"`
	HasInventorySources          bool        `json:"has_inventory_sources"`
	TotalInventorySources        int         `json:"total_inventory_sources"`
	InventorySourcesWithFailures int         `json:"inventory_sources_with_failures"`
	InsightsCredential           interface{} `json:"insights_credential"`
	PendingDeletion              bool        `json:"pending_deletion"`
}

type Credential struct {
	Description      string `json:"description"`
	CredentialTypeID int    `json:"credential_type_id"`
	ID               int    `json:"id"`
	Kind             string `json:"kind"`
	Name             string `json:"name"`
}

type Labels struct {
	Count   int           `json:"count"`
	Results []interface{} `json:"results"`
}

type Summary struct {
	Organization     *OrgnizationSummary `json:"organization"`
	CreatedBy        *ByUserSummary      `json:"created_by"`
	ModifiedBy       *ByUserSummary      `json:"modified_by"`
	ObjectRoles      *ObjectRoles        `json:"object_roles"`
	UserCapabilities *UserCapabilities   `json:"user_capabilities"`
	Project          *Project            `json:"project"`
	Inventory        *Inventory          `json:"inventory"`
	RecentJobs       []interface{}       `json:"recent_jobs"`
	Credentials      []Credential        `json:"credentials"`
	Labels           *Labels             `json:"labels"`
}

type ListInventoriesResponse struct {
	Pagination
	Results []*Inventory `json:"results"`
}

func (this *InventoriesService) ListInventories(params map[string]string) ([]*Inventory, *http.Response, error) {
	result := new(ListInventoriesResponse)
	endpoint := "/api/v2/inventories/"
	resp, err := this.client.Requester.GetJSON(endpoint, result, params)
	if err != nil {
		return nil, resp, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, resp, err
	}

	return result.Results, resp, nil
}
