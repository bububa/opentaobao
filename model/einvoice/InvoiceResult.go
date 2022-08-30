package einvoice

// InvoiceResult 结构体
type InvoiceResult struct {
	// 电子发票明细，erp开票默认不返回，如果erp需要获取阿里发票平台自动开票的结果，需要先找阿里小二开通权限
	InvoiceItems []InvoiceItem `json:"invoice_items,omitempty" xml:"invoice_items>invoice_item,omitempty"`
	// 防伪码
	AntiFakeCode string `json:"anti_fake_code,omitempty" xml:"anti_fake_code,omitempty"`
	// 发票密文，密码区的字符串
	Ciphertext string `json:"ciphertext,omitempty" xml:"ciphertext,omitempty"`
	// 税控设备编号(新版电子发票有)
	DeviceNo string `json:"device_no,omitempty" xml:"device_no,omitempty"`
	// erp自定义单据号
	ErpTid string `json:"erp_tid,omitempty" xml:"erp_tid,omitempty"`
	// 文件类型(pdf,jpg,png)
	FileDataType string `json:"file_data_type,omitempty" xml:"file_data_type,omitempty"`
	// 发票PDF的下载地址(仅在单个查询接口上显示，批量查询不显示)
	FilePath string `json:"file_path,omitempty" xml:"file_path,omitempty"`
	// 开票金额
	InvoiceAmount string `json:"invoice_amount,omitempty" xml:"invoice_amount,omitempty"`
	// 发票代码
	InvoiceCode string `json:"invoice_code,omitempty" xml:"invoice_code,omitempty"`
	// 开票日期
	InvoiceDate string `json:"invoice_date,omitempty" xml:"invoice_date,omitempty"`
	// 发票号码
	InvoiceNo string `json:"invoice_no,omitempty" xml:"invoice_no,omitempty"`
	// 电商平台代码。淘宝：taobao，天猫：tmall
	PlatformCode string `json:"platform_code,omitempty" xml:"platform_code,omitempty"`
	// 电商平台订单号
	PlatformTid string `json:"platform_tid,omitempty" xml:"platform_tid,omitempty"`
	// 开票流水号，唯一标志开票请求。如果两次请求流水号相同，则表示重复请求。
	SerialNo string `json:"serial_no,omitempty" xml:"serial_no,omitempty"`
	// 开票状态 (waiting = 开票中) 、(create_success = 开票成功)、(create_failed = 开票失败)
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 错误编码
	BizErrorCode string `json:"biz_error_code,omitempty" xml:"biz_error_code,omitempty"`
	// 错误信息
	BizErrorMsg string `json:"biz_error_msg,omitempty" xml:"biz_error_msg,omitempty"`
	// 发票类型，blue=蓝票，red=红票
	InvoiceType string `json:"invoice_type,omitempty" xml:"invoice_type,omitempty"`
	// 原蓝票发票代码，invoiceType=red时有值
	NormalInvoiceCode string `json:"normal_invoice_code,omitempty" xml:"normal_invoice_code,omitempty"`
	// 原蓝票发票号码，invoiceType=red时有值
	NormalInvoiceNo string `json:"normal_invoice_no,omitempty" xml:"normal_invoice_no,omitempty"`
	// 开票人，erp开票不返回，用来erp获取自动开票结果
	PayeeOperator string `json:"payee_operator,omitempty" xml:"payee_operator,omitempty"`
	// 收款人，erp开票不返回，用来erp获取自动开票结果
	PayeeReceiver string `json:"payee_receiver,omitempty" xml:"payee_receiver,omitempty"`
	// 复核人，erp开票不返回，用来erp获取自动开票结果
	PayeeChecker string `json:"payee_checker,omitempty" xml:"payee_checker,omitempty"`
	// 购买方抬头，erp开票不返回，用来erp获取自动开票结果
	PayerName string `json:"payer_name,omitempty" xml:"payer_name,omitempty"`
	// 购买方税号，erp开票不返回，用来erp获取自动开票结果
	PayerRegisterNo string `json:"payer_register_no,omitempty" xml:"payer_register_no,omitempty"`
	// 购买方企业电话，erp开票不返回，用来erp获取自动开票结果
	PayerPhone string `json:"payer_phone,omitempty" xml:"payer_phone,omitempty"`
	// 购买方企业地址，erp开票不返回，用来erp获取自动开票结果
	PayerAddress string `json:"payer_address,omitempty" xml:"payer_address,omitempty"`
	// 购买方企业银行及账号，erp开票不返回，用来erp获取自动开票结果
	PayerBankaccount string `json:"payer_bankaccount,omitempty" xml:"payer_bankaccount,omitempty"`
	// 销售方税号
	PayeeRegisterNo string `json:"payee_register_no,omitempty" xml:"payee_register_no,omitempty"`
	// 开票时间，时分秒格式（注意：2019-04-11之后开具的发票才返回）
	InvoiceTime string `json:"invoice_time,omitempty" xml:"invoice_time,omitempty"`
	// 二维码
	QrCode string `json:"qr_code,omitempty" xml:"qr_code,omitempty"`
	// 发票种类，0=电子发票，1=纸质发票，2=纸质专票
	InvoiceKind int64 `json:"invoice_kind,omitempty" xml:"invoice_kind,omitempty"`
}
