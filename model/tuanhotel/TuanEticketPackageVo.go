package tuanhotel

// TuanEticketPackageVo 
type TuanEticketPackageVo struct {

    // 核销放行码商名称
    
    PassMerchantUserName   string `json:"pass_merchant_user_name,omitempty" xml:"pass_merchant_user_name,omitempty"`
    

    // 门店分账是否支持资金到门店：0不支持，1支持
    
    IsMoneyToStore   int64 `json:"is_money_to_store,omitempty" xml:"is_money_to_store,omitempty"`
    

    // 分账方式 0-不分账，1-门店分账 2-宝贝分账
    
    BillType   int64 `json:"bill_type,omitempty" xml:"bill_type,omitempty"`
    

    // 核销帐号名称
    
    SendMerchantUserName   string `json:"send_merchant_user_name,omitempty" xml:"send_merchant_user_name,omitempty"`
    

    // 宝贝分账是否支持子账号核销
    
    IsSubAccount   int64 `json:"is_sub_account,omitempty" xml:"is_sub_account,omitempty"`
    

    // 系统自动生成
    
    Online   int64 `json:"online,omitempty" xml:"online,omitempty"`
    

    // 核销帐号id。电子凭证下，返回。关于发行码商和核销放行的说明 淘宝：0 阿里旅行代金券：2395482105 第三方码商（浙江深大：1077204890、孟宽测试：112414636、子畵码商：436160339、银科环企：2811880569）
    
    SendMerchantUserId   int64 `json:"send_merchant_user_id,omitempty" xml:"send_merchant_user_id,omitempty"`
    

    // 是否分次预约，0-否，1-是
    
    MultipleTimes   int64 `json:"multiple_times,omitempty" xml:"multiple_times,omitempty"`
    

    // 宝贝分账信息，宝贝分账必填。分账多个账号之间以逗号隔开。数据格式（名称:比例数值或金额:类型
    
    BillInfos   string `json:"bill_infos,omitempty" xml:"bill_infos,omitempty"`
    

    // 核销放行码商id
    
    PassMerchantUserId   int64 `json:"pass_merchant_user_id,omitempty" xml:"pass_merchant_user_id,omitempty"`
    

    // 分账描述
    
    BillTypeDesc   string `json:"bill_type_desc,omitempty" xml:"bill_type_desc,omitempty"`
    

}
