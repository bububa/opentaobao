package einvoice

// InvoiceApplyDtlDto 结构体
type InvoiceApplyDtlDto struct {
	// 该申请单 请求开票的结果，拆单的场景下可能有多笔发票请求  当apply_status为creating_inv, inv_failed, inv_success, inv_part_success 时返回该字段
	CreateInvResults []InvoiceCreateSimpleResultDto `json:"create_inv_results,omitempty" xml:"create_inv_results>invoice_create_simple_result_dto,omitempty"`
	// 申请明细列表
	InvoiceItems []InvoiceApplyItemsDto `json:"invoice_items,omitempty" xml:"invoice_items>invoice_apply_items_dto,omitempty"`
	// 发票申请模式，可选值：  pre_apply_url: URL预申请模式；适用于扫码开票（一单一码）的业务场景：业务前台提交开票金额等信息，请求阿里发票平台生成一个对应的发票申请页面URL。用户可在该页面中填写抬头等信息，然后提交正式的发票申请。  normal(默认为此模式): 正式提交用户的发票申请，商户根据此发票申请自动或审核开票。
	ApplyMode string `json:"apply_mode,omitempty" xml:"apply_mode,omitempty"`
	// 合计实付金额（申请开票的总金额，含税），格式为2位小数。开红票时传正数。  需满足公式：开票总金额(invoiceAmount) = 各项明细的交易金额(amount)之和 - 各项明细的优惠金额(discount)之和。  当指定auto_create_invoice=true或商户配置为自动开票时该字段必填。
	ApplyAmount string `json:"apply_amount,omitempty" xml:"apply_amount,omitempty"`
	// 发票申请ID
	ApplyId string `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	// 申请状态，可选值：  applying: 申请中，初始状态  cancelled: 申请已取消  confirmed: 商户已确认，待开/待录入发票  creating_inv: 开票中，待发票结果回传  inv_failed: 开票失败  inv_success: 开票成功  inv_part_success: 部分成功（拆单场景下存在。举例：发票申请拆单之后有10张票，其中有1张开票成功时，此时申请状态即为inv_part_success，当10张票全部成功申请状态则为inv_success）
	ApplyStatus string `json:"apply_status,omitempty" xml:"apply_status,omitempty"`
	// 申请创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 申请最近修改时间
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 发票备注
	InvoiceMemo string `json:"invoice_memo,omitempty" xml:"invoice_memo,omitempty"`
	// 发票(开票)类型，可选值：  blue: 蓝票  red: 红票
	InvoiceType string `json:"invoice_type,omitempty" xml:"invoice_type,omitempty"`
	// 原发票代码；  申请红票时必传
	NormalInvoiceCode string `json:"normal_invoice_code,omitempty" xml:"normal_invoice_code,omitempty"`
	// 原发票号码；  申请红票时必传
	NormalInvoiceNo string `json:"normal_invoice_no,omitempty" xml:"normal_invoice_no,omitempty"`
	// 外部业务方创建入驻工单的唯一幂等ID, 由业务方自己生成。  由字母数字组成
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 纳税人识别号
	PayeeRegisterNo string `json:"payee_register_no,omitempty" xml:"payee_register_no,omitempty"`
	// 购方地址，  专票必填。
	PayerAddress string `json:"payer_address,omitempty" xml:"payer_address,omitempty"`
	// 购方银行账号，  专票必填。
	PayerBankAccountId string `json:"payer_bank_account_id,omitempty" xml:"payer_bank_account_id,omitempty"`
	// 购方开户行名称，  专票必填。
	PayerBankName string `json:"payer_bank_name,omitempty" xml:"payer_bank_name,omitempty"`
	// 购方电子邮箱
	PayerEmail string `json:"payer_email,omitempty" xml:"payer_email,omitempty"`
	// 购方开票备注。不会显示在票面
	PayerMemo string `json:"payer_memo,omitempty" xml:"payer_memo,omitempty"`
	// 购方抬头
	PayerName string `json:"payer_name,omitempty" xml:"payer_name,omitempty"`
	// 购方联系电话，  专票必填。
	PayerPhone string `json:"payer_phone,omitempty" xml:"payer_phone,omitempty"`
	// 购方方税务登记证号，  开企业抬头时必填，  专票必填。
	PayerRegisterNo string `json:"payer_register_no,omitempty" xml:"payer_register_no,omitempty"`
	// 购方UID
	PayerUid string `json:"payer_uid,omitempty" xml:"payer_uid,omitempty"`
	// 业务前台的业务标记。  提供给业务前台打上特殊的业务标识，解决前台一些特殊场景，阿里发票不关心该字段的业务含义。
	PlatformBizFlag string `json:"platform_biz_flag,omitempty" xml:"platform_biz_flag,omitempty"`
	// 业务平台code, 由发票中台分配
	PlatformCode string `json:"platform_code,omitempty" xml:"platform_code,omitempty"`
	// 业务平台发票申请对应的订单号
	PlatformTid string `json:"platform_tid,omitempty" xml:"platform_tid,omitempty"`
	// 业务平台商户ID
	PlatformUserId string `json:"platform_user_id,omitempty" xml:"platform_user_id,omitempty"`
	// 红字发票信息表编号。  专票冲红时需要，商家跟税局申请
	RedNoticeNo string `json:"red_notice_no,omitempty" xml:"red_notice_no,omitempty"`
	// 抬头类型。可选值：  0：个人  1：企业
	BusinessType int64 `json:"business_type,omitempty" xml:"business_type,omitempty"`
	// 请求开票的销方信息  间联开票模式下，该字段无值。若调用方需要获取开具发票的销方信息，可调用查询发票详情接口
	CreateInvPayeeInfo *InvoiceCreatePayeeInfoDto `json:"create_inv_payee_info,omitempty" xml:"create_inv_payee_info,omitempty"`
	// 申请开票类型，可选值：  0: 电票  1：纸质普票  2：纸质专票
	InvoiceKind int64 `json:"invoice_kind,omitempty" xml:"invoice_kind,omitempty"`
	// 当前申请单是否为已终结状态。true: 是，false: 否。  主要用于区分inv_part_success状态下是终态还是中间态。
	IsFinally bool `json:"is_finally,omitempty" xml:"is_finally,omitempty"`
}
