package goserve

import (
	"encoding/json"
)

type Request struct {
	Status                   string      `json:"__status,omitempty"`
	Active                   bool        `json:"active,string,omitempty"`
	ActivityDue              string      `json:"activity_due,omitempty"`
	AdditionalAssigneeList   string      `json:"additional_assignee_list,omitempty"`
	Approval                 string      `json:"approval,omitempty"`
	ApprovalHistory          string      `json:"approval_history,omitempty"`
	ApprovalSet              string      `json:"approval_set,omitempty"`
	AssignedTo               string      `json:"assigned_to,omitempty"`
	AssignmentGroup          string      `json:"assignment_group,omitempty"`
	BackoutPlan              string      `json:"backout_plan,omitempty"`
	BusinessDuration         string      `json:"business_duration,omitempty"`
	BusinessService          string      `json:"business_service,omitempty"`
	CabDate                  string      `json:"cab_date,omitempty"` //
	CabDelegate              string      `json:"cab_delegate,omitempty"`
	CabRecommendation        string      `json:"cab_recommendation,omitempty"`
	CabRequired              bool        `json:"cab_required,string,omitempty"`
	CalendarDuration         string      `json:"calendar_duration,omitempty"`
	Category                 string      `json:"category,omitempty"`
	ChangePlan               string      `json:"change_plan,omitempty"`
	CloseCode                string      `json:"close_code,omitempty"`
	CloseNotes               string      `json:"close_notes,omitempty"`
	ClosedAt                 string      `json:"closed_at,omitempty"` //
	ClosedBy                 string      `json:"closed_by,omitempty"`
	CmdbCi                   string      `json:"cmdb_ci,omitempty"`
	Comments                 string      `json:"comments,omitempty"`
	CommentsAndWorkNotes     string      `json:"comments_and_work_notes,omitempty"`
	Company                  string      `json:"company,omitempty"`
	ConflictLastRun          string      `json:"conflict_last_run,omitempty"`
	ConflictStatus           string      `json:"conflict_status,omitempty"`
	ContactType              string      `json:"contact_type,omitempty"`
	CorrelationDisplay       string      `json:"correlation_display,omitempty"`
	CorrelationID            string      `json:"correlation_id,omitempty"`
	DeliveryPlan             string      `json:"delivery_plan,omitempty"`
	DeliveryTask             string      `json:"delivery_task,omitempty"`
	Description              string      `json:"description,omitempty"`
	DueDate                  string      `json:"due_date,omitempty"` //
	EndDate                  string      `json:"end_date,omitempty"` //
	Escalation               json.Number `json:"escalation,omitempty"`
	ExpectedStart            string      `json:"expected_start,omitempty"` //
	FollowUp                 string      `json:"follow_up,omitempty"`
	GroupList                string      `json:"group_list,omitempty"`
	Impact                   json.Number `json:"impact,omitempty"`
	ImplementationPlan       string      `json:"implementation_plan,omitempty"`
	Justification            string      `json:"justification,omitempty"`
	Knowledge                bool        `json:"knowledge,string,omitempty"`
	Location                 string      `json:"location,omitempty"`
	MadeSLA                  bool        `json:"made_sla,string,omitempty"`
	Number                   string      `json:"number,omitempty"`
	OnHold                   bool        `json:"on_hold,string,omitempty"`
	OnHoldReason             string      `json:"on_hold_reason,omitempty"`
	OpenedAt                 string      `json:"opened_at,omitempty"` //
	OpenedBy                 string      `json:"opened_by,omitempty"`
	Order                    string      `json:"order,omitempty"`
	OutsideMaintSchedule     bool        `json:"outside_maintenance_schedule,string,omitempty"`
	Parent                   string      `json:"parent,omitempty"`
	Phase                    string      `json:"phase,omitempty"`
	PhaseState               string      `json:"phase_state,omitempty"`
	Priority                 json.Number `json:"priority,omitempty"`
	ProductionSystem         bool        `json:"production_system,string,omitempty"`
	Reason                   string      `json:"reason,omitempty"`
	ReassignmentCount        json.Number `json:"reassignment_count,omitempty"`
	RequestedBy              string      `json:"requested_by,omitempty"`
	RequestedByDate          string      `json:"requested_by_date,omitempty"` //
	ReviewComments           string      `json:"review_comments,omitempty"`
	ReviewDate               string      `json:"review_date,omitempty"` //
	ReviewStatus             string      `json:"review_status,omitempty"`
	Risk                     json.Number `json:"risk,omitempty"`
	RiskImpactAnalysis       string      `json:"risk_impact_analysis,omitempty"`
	Scope                    json.Number `json:"scope,omitempty"`
	ShortDescription         string      `json:"short_description,omitempty"`
	SLADue                   string      `json:"sla_due,omitempty"`
	StartDate                string      `json:"start_date,omitempty"` //
	State                    string      `json:"state,omitempty"`
	Substate                 string      `json:"u_substate,omitempty"`
	StdChangeProducerVersion string      `json:"std_change_producer_version,omitempty"`
	SysClassName             string      `json:"sys_class_name,omitempty"`
	SysCreatedBy             string      `json:"sys_created_by,omitempty"`
	SysCreatedOn             string      `json:"sys_created_on,omitempty"` //
	SysDomain                string      `json:"sys_domain,omitempty"`
	SysDomainPath            string      `json:"sys_domain_path,omitempty"`
	SysID                    string      `json:"sys_id,omitempty"`
	SysModCount              json.Number `json:"sys_mod_count,omitempty"`
	SysTags                  string      `json:"sys_tags,omitempty"`
	SysUpdatedBy             string      `json:"sys_updated_by,omitempty"`
	SysUpdatedOn             string      `json:"sys_updated_on,omitempty"` //
	TestPlan                 string      `json:"test_plan,omitempty"`
	TimeWorked               string      `json:"time_worked,omitempty"`
	Type                     string      `json:"type,omitempty"`
	UponApproval             string      `json:"upon_approval,omitempty"`
	UponReject               string      `json:"upon_reject,omitempty"`
	Urgency                  string      `json:"urgency,omitempty"`
	UserInput                string      `json:"user_input,omitempty"`
	WatchList                string      `json:"watch_list,omitempty"`
	WorkEnd                  string      `json:"work_end,omitempty"`
	WorkNotes                string      `json:"work_notes,omitempty"`
	WorkNotesList            string      `json:"work_notes_list,omitempty"`
	WorkStart                string      `json:"work_start,omitempty"`
}
