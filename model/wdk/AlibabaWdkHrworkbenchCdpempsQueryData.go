package wdk

// AlibabaWdkHrworkbenchCdpempsQueryData 结构体
type AlibabaWdkHrworkbenchCdpempsQueryData struct {
	// 工号
	WorkNo string `json:"work_no,omitempty" xml:"work_no,omitempty"`
	// 商家编号
	MerchantCode string `json:"merchant_code,omitempty" xml:"merchant_code,omitempty"`
	// 消息摘要
	Md5 string `json:"md5,omitempty" xml:"md5,omitempty"`
	// 员工id
	EmpId int64 `json:"emp_id,omitempty" xml:"emp_id,omitempty"`
}
