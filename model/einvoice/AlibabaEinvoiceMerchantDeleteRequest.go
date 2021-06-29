package einvoice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
发票中台-同平台取消授权税号适用商户 API请求
alibaba.einvoice.merchant.delete

税号授权给同平台下其他商户使用后，可以使用此接口取消授权，被取消授权的商户失去开票能力
*/
type AlibabaEinvoiceMerchantDeleteRequest struct {
    model.Params
    // 验证码，商户首次绑定已入驻税号接口返回的taxToken
    taxToken   string
    // 业务方发起删除商户的唯一幂等ID, 由业务方生成。只能由字母和数字组成。
    outerId   string
    // 业务平台商户ID
    merchantUserId   string
    // 业务平台code, 由阿里发票分配
    platformCode   string
    // 税号
    payeeRegisterNo   string
}

// 初始化AlibabaEinvoiceMerchantDeleteRequest对象
func NewAlibabaEinvoiceMerchantDeleteRequest() *AlibabaEinvoiceMerchantDeleteRequest{
    return &AlibabaEinvoiceMerchantDeleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaEinvoiceMerchantDeleteRequest) GetApiMethodName() string {
    return "alibaba.einvoice.merchant.delete"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaEinvoiceMerchantDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TaxToken Setter
// 验证码，商户首次绑定已入驻税号接口返回的taxToken
func (r *AlibabaEinvoiceMerchantDeleteRequest) SetTaxToken(taxToken string) error {
    r.taxToken = taxToken
    r.Set("tax_token", taxToken)
    return nil
}

// TaxToken Getter
func (r AlibabaEinvoiceMerchantDeleteRequest) GetTaxToken() string {
    return r.taxToken
}
// OuterId Setter
// 业务方发起删除商户的唯一幂等ID, 由业务方生成。只能由字母和数字组成。
func (r *AlibabaEinvoiceMerchantDeleteRequest) SetOuterId(outerId string) error {
    r.outerId = outerId
    r.Set("outer_id", outerId)
    return nil
}

// OuterId Getter
func (r AlibabaEinvoiceMerchantDeleteRequest) GetOuterId() string {
    return r.outerId
}
// MerchantUserId Setter
// 业务平台商户ID
func (r *AlibabaEinvoiceMerchantDeleteRequest) SetMerchantUserId(merchantUserId string) error {
    r.merchantUserId = merchantUserId
    r.Set("merchant_user_id", merchantUserId)
    return nil
}

// MerchantUserId Getter
func (r AlibabaEinvoiceMerchantDeleteRequest) GetMerchantUserId() string {
    return r.merchantUserId
}
// PlatformCode Setter
// 业务平台code, 由阿里发票分配
func (r *AlibabaEinvoiceMerchantDeleteRequest) SetPlatformCode(platformCode string) error {
    r.platformCode = platformCode
    r.Set("platform_code", platformCode)
    return nil
}

// PlatformCode Getter
func (r AlibabaEinvoiceMerchantDeleteRequest) GetPlatformCode() string {
    return r.platformCode
}
// PayeeRegisterNo Setter
// 税号
func (r *AlibabaEinvoiceMerchantDeleteRequest) SetPayeeRegisterNo(payeeRegisterNo string) error {
    r.payeeRegisterNo = payeeRegisterNo
    r.Set("payee_register_no", payeeRegisterNo)
    return nil
}

// PayeeRegisterNo Getter
func (r AlibabaEinvoiceMerchantDeleteRequest) GetPayeeRegisterNo() string {
    return r.payeeRegisterNo
}
