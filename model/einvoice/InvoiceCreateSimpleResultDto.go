package einvoice

// InvoiceCreateSimpleResultDto 结构体
type InvoiceCreateSimpleResultDto struct {
	// 错误码
	BizErrorCode string `json:"biz_error_code,omitempty" xml:"biz_error_code,omitempty"`
	// 错误描述
	BizErrorMsg string `json:"biz_error_msg,omitempty" xml:"biz_error_msg,omitempty"`
	// 开票状态
	CreateStatus string `json:"create_status,omitempty" xml:"create_status,omitempty"`
	// 错误类型
	ErrorType string `json:"error_type,omitempty" xml:"error_type,omitempty"`
	// 实际开票金额
	InvoiceAmount string `json:"invoice_amount,omitempty" xml:"invoice_amount,omitempty"`
	// 销售方名称
	PayeeName string `json:"payee_name,omitempty" xml:"payee_name,omitempty"`
	// 销方税号
	PayeeRegisterNo string `json:"payee_register_no,omitempty" xml:"payee_register_no,omitempty"`
	// 税控产品
	ProductCode string `json:"product_code,omitempty" xml:"product_code,omitempty"`
	// 单张票序列号
	SingleSerialNo string `json:"single_serial_no,omitempty" xml:"single_serial_no,omitempty"`
	// 合计金额（不含税）
	SumPrice string `json:"sum_price,omitempty" xml:"sum_price,omitempty"`
	// 合计税额，格式为2位小数
	SumTax string `json:"sum_tax,omitempty" xml:"sum_tax,omitempty"`
	// 防伪码
	AntiFakeCode string `json:"anti_fake_code,omitempty" xml:"anti_fake_code,omitempty"`
	// 发票板式文件的下载地址。  need_download_url=true时返回。默认不生成。
	DownloadUrl string `json:"download_url,omitempty" xml:"download_url,omitempty"`
	// 发票板式文件下载地址(download_url)和预览图地址(invoice_image_url)的失效时间。  过期之后板式文件下载地址将无法访问，可以重新调用此接口，重新生成访问链接。  格式为：yyyy-MM-dd HH:mm:ss
	DownloadUrlExpires string `json:"download_url_expires,omitempty" xml:"download_url_expires,omitempty"`
	// 发票代码
	InvoiceCode string `json:"invoice_code,omitempty" xml:"invoice_code,omitempty"`
	// 开票日期；  当开票成功时必选
	InvoiceDate string `json:"invoice_date,omitempty" xml:"invoice_date,omitempty"`
	// 发票预览图的下载地址。  need_download_url=true时返回。默认不生成。
	InvoiceImageUrl string `json:"invoice_image_url,omitempty" xml:"invoice_image_url,omitempty"`
	// 发票号码
	InvoiceNo string `json:"invoice_no,omitempty" xml:"invoice_no,omitempty"`
	// 红票申请ID
	RedApplyId string `json:"red_apply_id,omitempty" xml:"red_apply_id,omitempty"`
	// 红票申请状态，定义同create_status
	RedCreateStatus string `json:"red_create_status,omitempty" xml:"red_create_status,omitempty"`
	// 开票发票类型，可选值：  0: 电票  1：纸质普票  2：纸质专票
	InvoiceKind int64 `json:"invoice_kind,omitempty" xml:"invoice_kind,omitempty"`
}
