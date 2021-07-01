package happytrip

// TravelHeadDto 结构体
type TravelHeadDto struct {
	// 单据创建人用户id，阿里工号
	CreateUserId string `json:"create_user_id,omitempty" xml:"create_user_id,omitempty"`
	// 外部差旅申请单id，用作数据同步主键
	OuterTravelHeadId string `json:"outer_travel_head_id,omitempty" xml:"outer_travel_head_id,omitempty"`
	// 同行人列表
	PeerStaffList []PeerStaff `json:"peer_staff_list,omitempty" xml:"peer_staff_list>peer_staff,omitempty"`
	// 单据审批状态， APPROVING：审批中,     APPROVED：已审批,     REJECTED：已拒绝
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 单据实际申请人用户id，阿里工号
	SubmitUserId string `json:"submit_user_id,omitempty" xml:"submit_user_id,omitempty"`
	// 出差原因说明
	TravelPurpose string `json:"travel_purpose,omitempty" xml:"travel_purpose,omitempty"`
	// 差旅类型，   ORDINARY：出差（内部项目&会议等），      RECRUITING：出差（外部合作&交流等），      CONFERENCE：外部招待，     MARKETING：公司大型活动，     PUBLIC_RELATIONSHIP：入职&候选人面试，   Training_Lecture：其它
	TravelType string `json:"travel_type,omitempty" xml:"travel_type,omitempty"`
	// 差旅行程列表
	LineList []TravelLineDto `json:"line_list,omitempty" xml:"line_list>travel_line_dto,omitempty"`
	// 申请单提交时间
	SubmitDate string `json:"submit_date,omitempty" xml:"submit_date,omitempty"`
}
