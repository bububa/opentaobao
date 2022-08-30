package ascp

// Process 结构体
type Process struct {
	// 包裹确认出库的时候必填
	Packages []Package `json:"packages,omitempty" xml:"packages>package,omitempty"`
	// 作业状态，ACCEPT(接单)、PRINT(打印)、PICK(拣货)、CHECK(复核)、PACKAGE(打包)、CONFIRM(确认出库)
	ProcessStatus string `json:"process_status,omitempty" xml:"process_status,omitempty"`
	// 创建发货单时的配编码
	LogisticsCode string `json:"logistics_code,omitempty" xml:"logistics_code,omitempty"`
	// 当前状态操作员编码
	OperatorCode string `json:"operator_code,omitempty" xml:"operator_code,omitempty"`
	// 当前状态操作员姓名
	OperatorName string `json:"operator_name,omitempty" xml:"operator_name,omitempty"`
	// 操作内容
	OperateInfo string `json:"operate_info,omitempty" xml:"operate_info,omitempty"`
	// 备注
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 运单号
	ExpressCode string `json:"express_code,omitempty" xml:"express_code,omitempty"`
	// 操作时间（时间戳）
	OperateTime int64 `json:"operate_time,omitempty" xml:"operate_time,omitempty"`
}
