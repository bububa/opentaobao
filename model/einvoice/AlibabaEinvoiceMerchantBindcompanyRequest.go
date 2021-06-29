package einvoice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
发票中台-跨平台绑定已入驻税号与商户 API请求
alibaba.einvoice.merchant.bindcompany

税号在阿里发票平台入驻成功后，允许业务方通过本接口跨业务平台绑定入驻税号和业务平台商户，绑定成功后该商户可以使用该税号的盘进行开票。绑定成功后，可以使用同平台授权、取消授权税号适用商户接口来变更税号和商户关系。
*/
type AlibabaEinvoiceMerchantBindcompanyRequest struct {
    model.Params
    // 业务方发起首次绑定门店的唯一幂等ID, 由业务方生成。只能由字母和数字组成。
    outerId   string
    // 业务平台商户ID
    merchantUserId   string
    // 激活码
    activationCode   string
    // 业务平台code, 由阿里发票分配
    platformCode   string
    // 税务登记号
    payeeRegisterNo   string
    // 业务平台门店名称
    merchantName   string
    // 税号已入驻的原业务平台code
    sourcePlatformCode   string
}

// 初始化AlibabaEinvoiceMerchantBindcompanyRequest对象
func NewAlibabaEinvoiceMerchantBindcompanyRequest() *AlibabaEinvoiceMerchantBindcompanyRequest{
    return &AlibabaEinvoiceMerchantBindcompanyRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaEinvoiceMerchantBindcompanyRequest) GetApiMethodName() string {
    return "alibaba.einvoice.merchant.bindcompany"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaEinvoiceMerchantBindcompanyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OuterId Setter
// 业务方发起首次绑定门店的唯一幂等ID, 由业务方生成。只能由字母和数字组成。
func (r *AlibabaEinvoiceMerchantBindcompanyRequest) SetOuterId(outerId string) error {
    r.outerId = outerId
    r.Set("outer_id", outerId)
    return nil
}

// OuterId Getter
func (r AlibabaEinvoiceMerchantBindcompanyRequest) GetOuterId() string {
    return r.outerId
}
// MerchantUserId Setter
// 业务平台商户ID
func (r *AlibabaEinvoiceMerchantBindcompanyRequest) SetMerchantUserId(merchantUserId string) error {
    r.merchantUserId = merchantUserId
    r.Set("merchant_user_id", merchantUserId)
    return nil
}

// MerchantUserId Getter
func (r AlibabaEinvoiceMerchantBindcompanyRequest) GetMerchantUserId() string {
    return r.merchantUserId
}
// ActivationCode Setter
// 激活码
func (r *AlibabaEinvoiceMerchantBindcompanyRequest) SetActivationCode(activationCode string) error {
    r.activationCode = activationCode
    r.Set("activation_code", activationCode)
    return nil
}

// ActivationCode Getter
func (r AlibabaEinvoiceMerchantBindcompanyRequest) GetActivationCode() string {
    return r.activationCode
}
// PlatformCode Setter
// 业务平台code, 由阿里发票分配
func (r *AlibabaEinvoiceMerchantBindcompanyRequest) SetPlatformCode(platformCode string) error {
    r.platformCode = platformCode
    r.Set("platform_code", platformCode)
    return nil
}

// PlatformCode Getter
func (r AlibabaEinvoiceMerchantBindcompanyRequest) GetPlatformCode() string {
    return r.platformCode
}
// PayeeRegisterNo Setter
// 税务登记号
func (r *AlibabaEinvoiceMerchantBindcompanyRequest) SetPayeeRegisterNo(payeeRegisterNo string) error {
    r.payeeRegisterNo = payeeRegisterNo
    r.Set("payee_register_no", payeeRegisterNo)
    return nil
}

// PayeeRegisterNo Getter
func (r AlibabaEinvoiceMerchantBindcompanyRequest) GetPayeeRegisterNo() string {
    return r.payeeRegisterNo
}
// MerchantName Setter
// 业务平台门店名称
func (r *AlibabaEinvoiceMerchantBindcompanyRequest) SetMerchantName(merchantName string) error {
    r.merchantName = merchantName
    r.Set("merchant_name", merchantName)
    return nil
}

// MerchantName Getter
func (r AlibabaEinvoiceMerchantBindcompanyRequest) GetMerchantName() string {
    return r.merchantName
}
// SourcePlatformCode Setter
// 税号已入驻的原业务平台code
func (r *AlibabaEinvoiceMerchantBindcompanyRequest) SetSourcePlatformCode(sourcePlatformCode string) error {
    r.sourcePlatformCode = sourcePlatformCode
    r.Set("source_platform_code", sourcePlatformCode)
    return nil
}

// SourcePlatformCode Getter
func (r AlibabaEinvoiceMerchantBindcompanyRequest) GetSourcePlatformCode() string {
    return r.sourcePlatformCode
}
