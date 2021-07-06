package qt

// QtReport 结构体
type QtReport struct {
	// 创建日期
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 送检人昵称
	Nick string `json:"nick,omitempty" xml:"nick,omitempty"`
	// 质检服务商名称
	SpName string `json:"sp_name,omitempty" xml:"sp_name,omitempty"`
	// 出具报告时间。
	GmtReport string `json:"gmt_report,omitempty" xml:"gmt_report,omitempty"`
	// 不合格原因
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 提交检查时间
	GmtSubmit string `json:"gmt_submit,omitempty" xml:"gmt_submit,omitempty"`
	// 质检编号
	QtCode string `json:"qt_code,omitempty" xml:"qt_code,omitempty"`
	// 质检项名称
	QtName string `json:"qt_name,omitempty" xml:"qt_name,omitempty"`
	// 样品的链接
	ItemUrl string `json:"item_url,omitempty" xml:"item_url,omitempty"`
	// 质检标准
	QtStandard string `json:"qt_standard,omitempty" xml:"qt_standard,omitempty"`
	// 质检报告地址
	ReportUrl string `json:"report_url,omitempty" xml:"report_url,omitempty"`
	// 扩展属性通过冒号分隔属性与属性值. 分号分隔不同属性
	ExtAttr string `json:"ext_attr,omitempty" xml:"ext_attr,omitempty"`
	// 质检有效期限
	GmtExpiry string `json:"gmt_expiry,omitempty" xml:"gmt_expiry,omitempty"`
	// 质检报告在淘宝的Id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 质检类型 0,全部 1,聚划算认证2,消保打标
	QtType int64 `json:"qt_type,omitempty" xml:"qt_type,omitempty"`
	// 样本宝贝id
	NumIid int64 `json:"num_iid,omitempty" xml:"num_iid,omitempty"`
	// 质检报告状态<br/> 0:已提交申请 1:已收到样品 2:已出检测结果 3.已出具报告
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 是否合格
	IsPassed bool `json:"is_passed,omitempty" xml:"is_passed,omitempty"`
}
