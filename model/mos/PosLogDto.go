package mos

// PosLogDto 结构体
type PosLogDto struct {
	// IP地址
	IpAddr string `json:"ip_addr,omitempty" xml:"ip_addr,omitempty"`
	// 当前操作步骤
	OperStep string `json:"oper_step,omitempty" xml:"oper_step,omitempty"`
	// 外部订单号
	OutTradeNo string `json:"out_trade_no,omitempty" xml:"out_trade_no,omitempty"`
	// pos机类型:POS("大POS"), SHOPPE(" 单品开票"), COUNTER("专柜自收银"), CLOUDPOS("云POS");
	PosType string `json:"pos_type,omitempty" xml:"pos_type,omitempty"`
	// pos机版本号
	Version string `json:"version,omitempty" xml:"version,omitempty"`
	// 门店号
	StoreNo string `json:"store_no,omitempty" xml:"store_no,omitempty"`
	// 专柜号
	CounterId string `json:"counter_id,omitempty" xml:"counter_id,omitempty"`
	// 联网:ON/OFF
	NetStat string `json:"net_stat,omitempty" xml:"net_stat,omitempty"`
	// 日志级别:INFO/DEBUG/ERROR
	LogLevel string `json:"log_level,omitempty" xml:"log_level,omitempty"`
	// 操作步骤结果描述
	OperResult string `json:"oper_result,omitempty" xml:"oper_result,omitempty"`
	// pos日志上传时间(本地时间)
	UploadTime string `json:"upload_time,omitempty" xml:"upload_time,omitempty"`
	// pos日志发生时间(本地时间)
	HappenTime string `json:"happen_time,omitempty" xml:"happen_time,omitempty"`
	// 云pos序列号
	Sn string `json:"sn,omitempty" xml:"sn,omitempty"`
	// 请求状态码
	HttpStatus int64 `json:"http_status,omitempty" xml:"http_status,omitempty"`
	// 小票号码
	ReceiptNo string `json:"receipt_no,omitempty" xml:"receipt_no,omitempty"`
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误源
	ErrSource string `json:"err_source,omitempty" xml:"err_source,omitempty"`
	// 扩展数据
	Extension string `json:"extension,omitempty" xml:"extension,omitempty"`
	// 日志数据类型:SYSTEM系统/BUSINESS 业务
	DataType string `json:"data_type,omitempty" xml:"data_type,omitempty"`
	// 错误描述
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// MAC地址
	MacAddr string `json:"mac_addr,omitempty" xml:"mac_addr,omitempty"`
	// 请求地址
	ApiUrl string `json:"api_url,omitempty" xml:"api_url,omitempty"`
	// 收银机号
	PosNo string `json:"pos_no,omitempty" xml:"pos_no,omitempty"`
	// OPER("步骤"), REMOTESERVICE(" 远程服务调用"), TRANS("传输日志"), DEFAULT("默认");
	LogType string `json:"log_type,omitempty" xml:"log_type,omitempty"`
	// 收银员
	Cashier string `json:"cashier,omitempty" xml:"cashier,omitempty"`
	// 请求内容
	RequestContent string `json:"request_content,omitempty" xml:"request_content,omitempty"`
	// 告警级别
	BizAlarmLevel string `json:"biz_alarm_level,omitempty" xml:"biz_alarm_level,omitempty"`
}
