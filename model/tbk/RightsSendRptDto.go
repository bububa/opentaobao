package tbk

// RightsSendRptDto 结构体
type RightsSendRptDto struct {
	// 渠道关系id的发放数据
	RelationRptList []RightsSendRelationRptDto `json:"relation_rpt_list,omitempty" xml:"relation_rpt_list>rights_send_relation_rpt_dto,omitempty"`
	// pid的发放数据
	PidRptList []RightsSendRelationRptDto `json:"pid_rpt_list,omitempty" xml:"pid_rpt_list>rights_send_relation_rpt_dto,omitempty"`
}
