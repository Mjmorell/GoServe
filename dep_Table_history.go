package goserve

//History is for the History table
type History struct {
	DateFormat         string `json:"date_format,omitempty"`
	Domain             string `json:"domain,omitempty"`
	ID                 string `json:"id,omitempty"`
	InitialValues      string `json:"initial_values,omitempty"`
	InternalCheckpoint string `json:"internal_checkpoint,omitempty"`
	Language           string `json:"language,omitempty"`
	LastUpdateRecorded string `json:"last_update_recorded,omitempty"`
	LineTable          string `json:"line_table,omitempty"`
	LoadTime           string `json:"load_time,omitempty"`
	Refresh            string `json:"refresh,omitempty"`
	SysCreatedBy       string `json:"sys_created_by,omitempty"`
	SysCreatedOn       string `json:"sys_created_on,omitempty"`
	SysID              string `json:"sys_id,omitempty"`
	SysModCount        string `json:"sys_mod_count,omitempty"`
	SysUpdatedBy       string `json:"sys_updated_by,omitempty"`
	SysUpdatedOn       string `json:"sys_updated_on,omitempty"`
	Table              string `json:"table,omitempty"`
	Timezone           string `json:"timezone,omitempty"`
	Updates            string `json:"updates,omitempty"`
	Variables          string `json:"variables,omitempty"`
}
