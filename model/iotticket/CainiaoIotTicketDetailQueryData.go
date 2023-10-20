package iotticket

// CainiaoIotTicketDetailQueryData 结构体
type CainiaoIotTicketDetailQueryData struct {
	// 图片列表
	Images []Images `json:"images,omitempty" xml:"images>images,omitempty"`
	// 操作记录
	OperateLogList []OperateLogList `json:"operate_log_list,omitempty" xml:"operate_log_list>operate_log_list,omitempty"`
	// 客户地址
	CustomerAddress string `json:"customer_address,omitempty" xml:"customer_address,omitempty"`
	// 事件类型描述
	EventTypeDesc string `json:"event_type_desc,omitempty" xml:"event_type_desc,omitempty"`
	// 创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 客户寄回设备邮编
	CustomerMailNo string `json:"customer_mail_no,omitempty" xml:"customer_mail_no,omitempty"`
	// 客户名称
	CustomerName string `json:"customer_name,omitempty" xml:"customer_name,omitempty"`
	// 设备bar code
	DeviceBarCode string `json:"device_bar_code,omitempty" xml:"device_bar_code,omitempty"`
	// 购买日期
	DevicePurchaseDate string `json:"device_purchase_date,omitempty" xml:"device_purchase_date,omitempty"`
	// 客户联系方式
	CustomerPhone string `json:"customer_phone,omitempty" xml:"customer_phone,omitempty"`
	// 工单详情描述
	TicketDescription string `json:"ticket_description,omitempty" xml:"ticket_description,omitempty"`
	// 服务商设备寄回邮寄编号
	SpMailNo string `json:"sp_mail_no,omitempty" xml:"sp_mail_no,omitempty"`
	// 服务状态描述：待受理；待发起维修方案；待确认维修方案；维修服务中；待确认完成；待评价；已完结；已撤销
	StatusDesc string `json:"status_desc,omitempty" xml:"status_desc,omitempty"`
	// 维修方案信息
	MaintenanceInfo *MaintenanceInfo `json:"maintenance_info,omitempty" xml:"maintenance_info,omitempty"`
	// 上门人员信息
	RepairmanInfo *RepairmanInfo `json:"repairman_info,omitempty" xml:"repairman_info,omitempty"`
	// 工单Id
	TicketId int64 `json:"ticket_id,omitempty" xml:"ticket_id,omitempty"`
}
