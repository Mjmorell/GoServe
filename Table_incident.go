package goserve

//Incident is for the Incident table
type Incident struct {
	Active                    string `json:"active,omitempty"`
	ActivityDue               string `json:"activity_due,omitempty"`
	AdditionalAssigneeList    string `json:"additional_assignee_list,omitempty"`
	Approval                  string `json:"approval,omitempty"`
	ApprovalHistory           string `json:"approval_history,omitempty"`
	ApprovalSet               string `json:"approval_set,omitempty"`
	AssignedTo                string `json:"assigned_to,omitempty"`
	AssignmentGroup           string `json:"assignment_group,omitempty"`
	BusinessDuration          string `json:"business_duration,omitempty"`
	BusinessService           string `json:"business_service,omitempty"`
	BusinessStc               string `json:"business_stc,omitempty"`
	CalendarDuration          string `json:"calendar_duration,omitempty"`
	CalendarStc               string `json:"calendar_stc,omitempty"`
	CallerID                  string `json:"caller_id,omitempty"`
	Category                  string `json:"category,omitempty"`
	CausedBy                  string `json:"caused_by,omitempty"`
	CloseCode                 string `json:"close_code,omitempty"`
	CloseNotes                string `json:"close_notes,omitempty"`
	ClosedAt                  string `json:"closed_at,omitempty"`
	ClosedBy                  string `json:"closed_by,omitempty"`
	CmdbCi                    string `json:"cmdb_ci,omitempty"`
	Comments                  string `json:"comments,omitempty"`
	Company                   string `json:"company,omitempty"`
	ContactType               string `json:"contact_type,omitempty"`
	Contract                  string `json:"contract,omitempty"`
	CorrelationDisplay        string `json:"correlation_display,omitempty"`
	CorrelationID             string `json:"correlation_id,omitempty"`
	DeliveryPlan              string `json:"delivery_plan,omitempty"`
	DeliveryTask              string `json:"delivery_task,omitempty"`
	Description               string `json:"description,omitempty"`
	DueDate                   string `json:"due_date,omitempty"`
	Effort                    string `json:"effort,omitempty"`
	EndDate                   string `json:"end_date,omitempty"`
	Escalation                string `json:"escalation,omitempty"`
	ExpectedStart             string `json:"expected_start,omitempty"`
	FollowUp                  string `json:"follow_up,omitempty"`
	GroupList                 string `json:"group_list,omitempty"`
	Impact                    string `json:"impact,omitempty"`
	IncidentState             string `json:"incident_state,omitempty"`
	Knowledge                 string `json:"knowledge,omitempty"`
	Location                  string `json:"location,omitempty"`
	MadeSLA                   string `json:"made_sla,omitempty"`
	Notify                    string `json:"notify,omitempty"`
	Number                    string `json:"number,omitempty"`
	OpenedAt                  string `json:"opened_at,omitempty"`
	OpenedBy                  string `json:"opened_by,omitempty"`
	Order                     string `json:"order,omitempty"`
	Parent                    string `json:"parent,omitempty"`
	Priority                  string `json:"priority,omitempty"`
	ProblemID                 string `json:"problem_id,omitempty"`
	ReassignmentCount         string `json:"reassignment_count,omitempty"`
	RemainingEffort           string `json:"remaining_effort,omitempty"`
	RequestedFor              string `json:"requested_for,omitempty"`
	Rfc                       string `json:"rfc,omitempty"`
	Severity                  string `json:"severity,omitempty"`
	ShortDescription          string `json:"short_description,omitempty"`
	Skills                    string `json:"skills,omitempty"`
	SLADue                    string `json:"sla_due,omitempty"`
	StartDate                 string `json:"start_date,omitempty"`
	State                     string `json:"state,omitempty"`
	Subcategory               string `json:"subcategory,omitempty"`
	SysClassName              string `json:"sys_class_name,omitempty"`
	SysCreatedBy              string `json:"sys_created_by,omitempty"`
	SysCreatedOn              string `json:"sys_created_on,omitempty"`
	SysDomain                 string `json:"sys_domain,omitempty"`
	SysID                     string `json:"sys_id,omitempty"`
	SysModCount               string `json:"sys_mod_count,omitempty"`
	SysTags                   string `json:"sys_tags,omitempty"`
	SysUpdatedBy              string `json:"sys_updated_by,omitempty"`
	SysUpdatedOn              string `json:"sys_updated_on,omitempty"`
	TimeWorked                string `json:"time_worked,omitempty"`
	Type                      string `json:"type,omitempty"`
	UAccountable              string `json:"u_accountable,omitempty"`
	UAssignmentDepartment     string `json:"u_assignment_department,omitempty"`
	UAssignmentDivision       string `json:"u_assignment_division,omitempty"`
	UAssignmentOrganization   string `json:"u_assignment_organization,omitempty"`
	UAssignmentTeam           string `json:"u_assignment_team,omitempty"`
	UBaselineDueDate          string `json:"u_baseline_due_date,omitempty"`
	UBenefitingCostCenter     string `json:"u_benefiting_cost_center,omitempty"`
	UBenefitingDepartment     string `json:"u_benefiting_department,omitempty"`
	UBudgetYear               string `json:"u_budget_year,omitempty"`
	UCausedByTask             string `json:"u_caused_by_task,omitempty"`
	UDemandRequest            string `json:"u_demand_request,omitempty"`
	UDisposition              string `json:"u_disposition,omitempty"`
	UDivision                 string `json:"u_division,omitempty"`
	UDueDateChangeCount       string `json:"u_due_date_change_count,omitempty"`
	UEmailNotification        string `json:"u_email_notification,omitempty"`
	UEscalatedAt              string `json:"u_escalated_at,omitempty"`
	UEscalatedBy              string `json:"u_escalated_by,omitempty"`
	UExternalSysID            string `json:"u_external_sys_id,omitempty"`
	UFinalMetric              string `json:"u_final_metric,omitempty"`
	UFinalMetricCode          string `json:"u_final_metric_code,omitempty"`
	UFinalPerformance         string `json:"u_final_performance,omitempty"`
	UFundingDepartment        string `json:"u_funding_department,omitempty"`
	UJiraID                   string `json:"u_jira_id,omitempty"`
	UOrganization             string `json:"u_organization,omitempty"`
	UOwner                    string `json:"u_owner,omitempty"`
	UProduct                  string `json:"u_product,omitempty"`
	URank                     string `json:"u_rank,omitempty"`
	UReleasePackage           string `json:"u_release_package,omitempty"`
	URequestedFor             string `json:"u_requested_for,omitempty"`
	URequestedForOrganization string `json:"u_requested_for_organization,omitempty"`
	UResolvedAt               string `json:"u_resolved_at,omitempty"`
	UResolvedBy               string `json:"u_resolved_by,omitempty"`
	UResolvedDuration         string `json:"u_resolved_duration,omitempty"`
	USlaBreach                string `json:"u_sla_breach,omitempty"`
	USlaType                  string `json:"u_sla_type,omitempty"`
	USubstate                 string `json:"u_substate,omitempty"`
	UTaskScheduler            string `json:"u_task_scheduler,omitempty"`
	UTemplate                 string `json:"u_template,omitempty"`
	UUrl                      string `json:"u_url,omitempty"`
	UValidated                string `json:"u_validated,omitempty"`
	UVendor                   string `json:"u_vendor,omitempty"`
	UponApproval              string `json:"upon_approval,omitempty"`
	UponReject                string `json:"upon_reject,omitempty"`
	Urgency                   string `json:"urgency,omitempty"`
	UserInput                 string `json:"user_input,omitempty"`
	Variables                 string `json:"variables,omitempty"`
	WorkEffort                string `json:"work_effort,omitempty"`
	WorkEnd                   string `json:"work_end,omitempty"`
	WorkNotes                 string `json:"work_notes,omitempty"`
	WorkStart                 string `json:"work_start,omitempty"`
}
