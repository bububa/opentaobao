package einvoice

// InvoiceApplyDto 
type InvoiceApplyDto struct {

    // 合计实付金额（申请开票的总金额，含税），格式为2位小数。开红票时传正数。需满足公式：开票总金额(invoiceAmount) = 各项明细的交易金额(amount)之和 - 各项明细的优惠金额(discount)之和
    
    ApplyAmount   string `json:"apply_amount,omitempty" xml:"apply_amount,omitempty"`
    

    // 请求来源：order: 下单
    
    ApplySource   string `json:"apply_source,omitempty" xml:"apply_source,omitempty"`
    

    // 当前申请单是否自动开票。当业务前台传入了该字段时，以前台传入的值为准。当前台未传值时，中台会读取商户在中台维护的自动开票配置。true: 申请单会自动转开票请求，调用税控进行开票。false: 申请单数据会在中台落地，状态为申请中。不会发起开票请求。适用于商户需要人工审核之后，再确认开票的场景。
    
    AutoCreateInvoice   bool `json:"auto_create_invoice,omitempty" xml:"auto_create_invoice,omitempty"`
    

    // 抬头类型。可选值：0：个人1：企业；当apply_mode=pre_apply_url时可选
    
    BusinessType   int64 `json:"business_type,omitempty" xml:"business_type,omitempty"`
    

    // 开票明细
    
    InvoiceItems   []InvoiceApplyItemsDto `json:"invoice_items,omitempty" xml:"invoice_items,omitempty"`
    

    // 开票发票类型可选值：0: 电票1：纸质普票2：纸质专票
    
    InvoiceKind   int64 `json:"invoice_kind,omitempty" xml:"invoice_kind,omitempty"`
    

    // 发票(开票)类型，可选值：blue: 蓝票red: 红票
    
    InvoiceType   string `json:"invoice_type,omitempty" xml:"invoice_type,omitempty"`
    

    // 征税方式，0普通征收，1减按征收，2差额征收
    
    LevyType   int64 `json:"levy_type,omitempty" xml:"levy_type,omitempty"`
    

    // 原发票代码(开红票时使用)
    
    NormalInvoiceCode   string `json:"normal_invoice_code,omitempty" xml:"normal_invoice_code,omitempty"`
    

    // 原发票号码(开红票时使用)
    
    NormalInvoiceNo   string `json:"normal_invoice_no,omitempty" xml:"normal_invoice_no,omitempty"`
    

    // 外部业务方发起开票申请的唯一幂等ID,?由调用平台生成。只能由字母和数字组成。
    
    OuterId   string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
    

    // 购方地址，专票必填。
    
    PayerAddress   string `json:"payer_address,omitempty" xml:"payer_address,omitempty"`
    

    // 购方银行账号，专票必填。
    
    PayerBankAccountId   string `json:"payer_bank_account_id,omitempty" xml:"payer_bank_account_id,omitempty"`
    

    // 购方开户行名称，专票必填。
    
    PayerBankName   string `json:"payer_bank_name,omitempty" xml:"payer_bank_name,omitempty"`
    

    // 购方电子邮箱，需满足邮箱格式。  格式要求：\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*
    
    PayerEmail   string `json:"payer_email,omitempty" xml:"payer_email,omitempty"`
    

    // 购方收票物流信息，用于纸票场景。
    
    PayerLogisticsInfo   *PayerLogisticsInfoDto `json:"payer_logistics_info,omitempty" xml:"payer_logistics_info,omitempty"`
    

    // 购方开票备注。不会显示在票面
    
    PayerMemo   string `json:"payer_memo,omitempty" xml:"payer_memo,omitempty"`
    

    // 购方抬头；当apply_mode=pre_apply_url时可选
    
    PayerName   string `json:"payer_name,omitempty" xml:"payer_name,omitempty"`
    

    // 购方联系电话，专票必填。
    
    PayerPhone   string `json:"payer_phone,omitempty" xml:"payer_phone,omitempty"`
    

    // 购方方税务登记证号，由大写字母或数字组成，长度要求15~20位。开企业抬头时必填，专票必填。
    
    PayerRegisterNo   string `json:"payer_register_no,omitempty" xml:"payer_register_no,omitempty"`
    

    // 购方UID
    
    PayerUid   string `json:"payer_uid,omitempty" xml:"payer_uid,omitempty"`
    

    // 购方联系电话，专票必填。
    
    PhoneNumber   string `json:"phone_number,omitempty" xml:"phone_number,omitempty"`
    

    // 业务前台的业务标记。提供给业务前台打上特殊的业务标识，解决前台一些特殊场景，中台不关心该字段的业务含义。
    
    PlatformBizFlag   string `json:"platform_biz_flag,omitempty" xml:"platform_biz_flag,omitempty"`
    

    // 请求来源平台Code, 由发票中台分配
    
    PlatformCode   string `json:"platform_code,omitempty" xml:"platform_code,omitempty"`
    

    // 业务平台发票申请对应的订单号。
    
    PlatformTid   string `json:"platform_tid,omitempty" xml:"platform_tid,omitempty"`
    

    // 业务平台商户ID/用户ID
    
    PlatformUserId   string `json:"platform_user_id,omitempty" xml:"platform_user_id,omitempty"`
    

    // 红字通知单号，专票冲红时需要，商家跟税局申请
    
    RedNoticeNo   string `json:"red_notice_no,omitempty" xml:"red_notice_no,omitempty"`
    

    // 来源标识
    
    SourceFlag   string `json:"source_flag,omitempty" xml:"source_flag,omitempty"`
    

    // 特殊票种标识，可选值：02: 农产品收购票
    
    SpecialFlag   string `json:"special_flag,omitempty" xml:"special_flag,omitempty"`
    

    // 交易时间
    
    TradeTime   string `json:"trade_time,omitempty" xml:"trade_time,omitempty"`
    

    // 业务来源平台, 由发票中台分配
    
    SourcePlatformCode   string `json:"source_platform_code,omitempty" xml:"source_platform_code,omitempty"`
    

    // 销方税务登记证号，长度要求15~20位。  传了此参数，则阿里发票平台会使用传入的销方税号进行开票。  未传则阿里发票平台会自动选择商户入驻的税号进行开票。
    
    PayeeRegisterNo   string `json:"payee_register_no,omitempty" xml:"payee_register_no,omitempty"`
    

    // 发票申请模式，可选值：  pre_apply_url: URL预申请模式；适用于扫码开票（一单一码）的业务场景：业务前台提交开票金额等信息，请求阿里发票平台生成一个对应的发票申请页面URL。用户可在该页面中填写抬头等信息，然后提交正式的发票申请。  normal(默认为此模式): 正式提交用户的发票申请，商户根据此发票申请自动或审核开票。
    
    ApplyMode   string `json:"apply_mode,omitempty" xml:"apply_mode,omitempty"`
    

    // 请求开票的销方信息。 传了此参数，则使用传入的销方信息进行开票。 未传则会以商户维护在阿里发票平台的销方信息为准。 业务前台请根据商户在哪维护销方信息进行选择，推荐后者。
    
    CreateInvPayeeInfo   *InvoiceCreatePayeeInfoDto `json:"create_inv_payee_info,omitempty" xml:"create_inv_payee_info,omitempty"`
    

    // 指定的开票税控设备ID 传了此参数，则使用传入的设备ID进行开票。 未传则会使用商户维护在阿里发票平台的默认设备开票。 业务前台请根据商户在哪维护税控设备进行选择，推荐后者。
    
    DeviceId   string `json:"device_id,omitempty" xml:"device_id,omitempty"`
    

    // 发票备注，会显示在票面
    
    InvoiceMemo   string `json:"invoice_memo,omitempty" xml:"invoice_memo,omitempty"`
    

}
