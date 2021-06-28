package wdk

// CreateContractTemplateRequest 
/* model for simplify = false
type CreateContractTemplateRequest struct {

    // 提货券配置列表
    
    CouponConfigList  struct {
        Couponconfiglist  []Couponconfiglist `json:"couponconfiglist,omitempty"`
    } `json:"coupon_config_list,omitempty"`
    

    // 乙方签约日期-日
    
    SecondSignDay   string `json:"second_sign_day,omitempty"`
    

    // 乙方签约日期-月
    
    SecondSignMonth   string `json:"second_sign_month,omitempty"`
    

    // 乙方签约日期-年
    
    SecondSignYear   string `json:"second_sign_year,omitempty"`
    

    // 乙方法人
    
    SecondLegalPerson   string `json:"second_legal_person,omitempty"`
    

    // 甲方签约日期-日
    
    FirstSignDay   string `json:"first_sign_day,omitempty"`
    

    // 甲方签约日期-月
    
    FirstSignMonth   string `json:"first_sign_month,omitempty"`
    

    // 甲方签约日期-年
    
    FirstSignYear   string `json:"first_sign_year,omitempty"`
    

    // 甲方法人
    
    FirstLegalPerson   string `json:"first_legal_person,omitempty"`
    

    // 乙方银行账号
    
    SecondAccountNo   string `json:"second_account_no,omitempty"`
    

    // 乙方开户行
    
    SecondBankName   string `json:"second_bank_name,omitempty"`
    

    // 乙方账户名称
    
    SecondAccountName   string `json:"second_account_name,omitempty"`
    

    // 甲方银行账号
    
    FirstAccountNo   string `json:"first_account_no,omitempty"`
    

    // 甲方开户行
    
    FirstBankName   string `json:"first_bank_name,omitempty"`
    

    // 甲方账户名称
    
    FirstAccountName   string `json:"first_account_name,omitempty"`
    

    // 甲方为提供乙方的折扣数
    
    Discount   string `json:"discount,omitempty"`
    

    // 乙方支付甲方款项超时天数
    
    PayTimeOutDay   string `json:"pay_time_out_day,omitempty"`
    

    // 发票类型简写
    
    InvoiceTypeSimpleName   string `json:"invoice_type_simple_name,omitempty"`
    

    // 保证金（单位：元）
    
    Bond   string `json:"bond,omitempty"`
    

    // 最高赊销额度（单位：元）
    
    MaxCreditAmount   string `json:"max_credit_amount,omitempty"`
    

    // 合同有效期结束时间-日
    
    ContractEndDay   string `json:"contract_end_day,omitempty"`
    

    // 合同有效期结束时间-月
    
    ContractEndMonth   string `json:"contract_end_month,omitempty"`
    

    // 合同有效期结束时间-年
    
    ContractEndYear   string `json:"contract_end_year,omitempty"`
    

    // 合同有效期开始时间-日
    
    ContractStartDay   string `json:"contract_start_day,omitempty"`
    

    // 合同有效期开始时间-月
    
    ContractStartMonth   string `json:"contract_start_month,omitempty"`
    

    // 合同有效期开始时间-年
    
    ContractStartYear   string `json:"contract_start_year,omitempty"`
    

    // 券使用的结束时间-日
    
    CouponEndDay   string `json:"coupon_end_day,omitempty"`
    

    // 券使用的结束时间-月
    
    CouponEndMonth   string `json:"coupon_end_month,omitempty"`
    

    // 券使用的结束时间-年
    
    CouponEndYear   string `json:"coupon_end_year,omitempty"`
    

    // 券使用的开始时间-日
    
    CouponStartDay   string `json:"coupon_start_day,omitempty"`
    

    // 券使用的开始时间-月
    
    CouponStartMonth   string `json:"coupon_start_month,omitempty"`
    

    // 券使用的开始时间-年
    
    CouponStartYear   string `json:"coupon_start_year,omitempty"`
    

    // 门店详细地址
    
    PoiDetailAddress   string `json:"poi_detail_address,omitempty"`
    

    // 门店所在街道
    
    PoiStreet   string `json:"poi_street,omitempty"`
    

    // 门店所在区
    
    PoiArea   string `json:"poi_area,omitempty"`
    

    // 门店所在市
    
    PoiCity   string `json:"poi_city,omitempty"`
    

    // 不提供的商品信息，传入格式同providerItemInfos
    
    NotProviderItemInfos  struct {
        String  []string `json:"string,omitempty"`
    } `json:"not_provider_item_infos,omitempty"`
    

    // 提供的商品信息，需要传入一个String的List，每个字符串描述一个商品
    
    ProviderItemInfos  struct {
        String  []string `json:"string,omitempty"`
    } `json:"provider_item_infos,omitempty"`
    

    // 大润发超市名称
    
    PoiName   string `json:"poi_name,omitempty"`
    

    // 乙方公司地址
    
    SecondCompanyAddress   string `json:"second_company_address,omitempty"`
    

    // 乙方公司名称
    
    SecondCompanyName   string `json:"second_company_name,omitempty"`
    

    // 甲方公司地址
    
    FirstCompanyAddress   string `json:"first_company_address,omitempty"`
    

    // 甲方公司名称
    
    FirstCompanyName   string `json:"first_company_name,omitempty"`
    

}
*/

// CreateContractTemplateRequest 
type CreateContractTemplateRequest struct {

    // 提货券配置列表
    CouponConfigList   []Couponconfiglist `json:"coupon_config_list,omitempty"`

    // 乙方签约日期-日
    SecondSignDay   string `json:"second_sign_day,omitempty"`

    // 乙方签约日期-月
    SecondSignMonth   string `json:"second_sign_month,omitempty"`

    // 乙方签约日期-年
    SecondSignYear   string `json:"second_sign_year,omitempty"`

    // 乙方法人
    SecondLegalPerson   string `json:"second_legal_person,omitempty"`

    // 甲方签约日期-日
    FirstSignDay   string `json:"first_sign_day,omitempty"`

    // 甲方签约日期-月
    FirstSignMonth   string `json:"first_sign_month,omitempty"`

    // 甲方签约日期-年
    FirstSignYear   string `json:"first_sign_year,omitempty"`

    // 甲方法人
    FirstLegalPerson   string `json:"first_legal_person,omitempty"`

    // 乙方银行账号
    SecondAccountNo   string `json:"second_account_no,omitempty"`

    // 乙方开户行
    SecondBankName   string `json:"second_bank_name,omitempty"`

    // 乙方账户名称
    SecondAccountName   string `json:"second_account_name,omitempty"`

    // 甲方银行账号
    FirstAccountNo   string `json:"first_account_no,omitempty"`

    // 甲方开户行
    FirstBankName   string `json:"first_bank_name,omitempty"`

    // 甲方账户名称
    FirstAccountName   string `json:"first_account_name,omitempty"`

    // 甲方为提供乙方的折扣数
    Discount   string `json:"discount,omitempty"`

    // 乙方支付甲方款项超时天数
    PayTimeOutDay   string `json:"pay_time_out_day,omitempty"`

    // 发票类型简写
    InvoiceTypeSimpleName   string `json:"invoice_type_simple_name,omitempty"`

    // 保证金（单位：元）
    Bond   string `json:"bond,omitempty"`

    // 最高赊销额度（单位：元）
    MaxCreditAmount   string `json:"max_credit_amount,omitempty"`

    // 合同有效期结束时间-日
    ContractEndDay   string `json:"contract_end_day,omitempty"`

    // 合同有效期结束时间-月
    ContractEndMonth   string `json:"contract_end_month,omitempty"`

    // 合同有效期结束时间-年
    ContractEndYear   string `json:"contract_end_year,omitempty"`

    // 合同有效期开始时间-日
    ContractStartDay   string `json:"contract_start_day,omitempty"`

    // 合同有效期开始时间-月
    ContractStartMonth   string `json:"contract_start_month,omitempty"`

    // 合同有效期开始时间-年
    ContractStartYear   string `json:"contract_start_year,omitempty"`

    // 券使用的结束时间-日
    CouponEndDay   string `json:"coupon_end_day,omitempty"`

    // 券使用的结束时间-月
    CouponEndMonth   string `json:"coupon_end_month,omitempty"`

    // 券使用的结束时间-年
    CouponEndYear   string `json:"coupon_end_year,omitempty"`

    // 券使用的开始时间-日
    CouponStartDay   string `json:"coupon_start_day,omitempty"`

    // 券使用的开始时间-月
    CouponStartMonth   string `json:"coupon_start_month,omitempty"`

    // 券使用的开始时间-年
    CouponStartYear   string `json:"coupon_start_year,omitempty"`

    // 门店详细地址
    PoiDetailAddress   string `json:"poi_detail_address,omitempty"`

    // 门店所在街道
    PoiStreet   string `json:"poi_street,omitempty"`

    // 门店所在区
    PoiArea   string `json:"poi_area,omitempty"`

    // 门店所在市
    PoiCity   string `json:"poi_city,omitempty"`

    // 不提供的商品信息，传入格式同providerItemInfos
    NotProviderItemInfos   []string `json:"not_provider_item_infos,omitempty"`

    // 提供的商品信息，需要传入一个String的List，每个字符串描述一个商品
    ProviderItemInfos   []string `json:"provider_item_infos,omitempty"`

    // 大润发超市名称
    PoiName   string `json:"poi_name,omitempty"`

    // 乙方公司地址
    SecondCompanyAddress   string `json:"second_company_address,omitempty"`

    // 乙方公司名称
    SecondCompanyName   string `json:"second_company_name,omitempty"`

    // 甲方公司地址
    FirstCompanyAddress   string `json:"first_company_address,omitempty"`

    // 甲方公司名称
    FirstCompanyName   string `json:"first_company_name,omitempty"`

}
