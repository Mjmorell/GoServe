package goserve

//SysAudit is for the SysAudit table
type SysAudit struct {
	Documentkey        string `json:"documentkey,omitempty"`
	Fieldname          string `json:"fieldname,omitempty"`
	InternalCheckpoint string `json:"internal_checkpoint,omitempty"`
	Newvalue           string `json:"newvalue,omitempty"`
	Oldvalue           string `json:"oldvalue,omitempty"`
	Reason             string `json:"reason,omitempty"`
	RecordCheckpoint   string `json:"record_checkpoint,omitempty"`
	SysCreatedBy       string `json:"sys_created_by,omitempty"`
	SysCreatedOn       string `json:"sys_created_on,omitempty"`
	SysID              string `json:"sys_id,omitempty"`
	Tablename          string `json:"tablename,omitempty"`
	User               string `json:"user,omitempty"`
	Variables          string `json:"variables,omitempty"`
}
