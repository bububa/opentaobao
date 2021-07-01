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
type AlibabaEinvoiceMerchantDeleteAPIRequest struct {
    model.Params
    // 验证码，商户首次绑定已入驻税号接口返回的taxToken
    _taxToken   string
    // 业务方发起删除商户的唯一幂等ID, 由业务方生成。只能由字母和数字组成。
    _outerId   string
    // 业务平台商户ID
    _merchantUserId   string
    // 业务平台code, 由阿里发票分配
    _platformCode   string
    // 税号
    _payeeRegisterNo   string
}

// 初始化AlibabaEinvoiceMerchantDeleteAPIRequest对象
func NewAlibabaEinvoiceMerchantDeleteRequest() *AlibabaEinvoiceMerchantDeleteAPIRequest{
    return &AlibabaEinvoiceMerchantDeleteAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaEinvoiceMerchantDeleteAPIRequest) GetApiMethodName() string {
    return "alibaba.einvoice.merchant.delete"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaEinvoiceMerchantDeleteAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TaxToken Setter
// 验证码，商户首次绑定已入驻税号接口返回的taxToken
func (r *AlibabaEinvoiceMerchantDeleteAPIRequest) SetTaxToken(_taxToken string) error {
    r._taxToken = _taxToken
    r.Set("tax_token", _taxToken)
    return nil
}

// TaxToken Getter
func (r AlibabaEinvoiceMerchantDeleteAPIRequest) GetTaxToken() string {
    return r._taxToken
}
// OuterId Setter
// 业务方发起删除商户的唯一幂等ID, 由业务方生成。只能由字母和数字组成。
func (r *AlibabaEinvoiceMerchantDeleteAPIRequest) SetOuterId(_outerId string) error {
    r._outerId = _outerId
    r.Set("outer_id", _outerId)
    return nil
}

// OuterId Getter
func (r AlibabaEinvoiceMerchantDeleteAPIRequest) GetOuterId() string {
    return r._outerId
}
// MerchantUserId Setter
// 业务平台商户ID
func (r *AlibabaEinvoiceMerchantDeleteAPIRequest) SetMerchantUserId(_merchantUserId string) error {
    r._merchantUserId = _merchantUserId
    r.Set("merchant_user_id", _merchantUserId)
    return nil
}

// MerchantUserId Getter
func (r AlibabaEinvoiceMerchantDeleteAPIRequest) GetMerchantUserId() string {
    return r._merchantUserId
}
// PlatformCode Setter
// 业务平台code, 由阿里发票分配
func (r *AlibabaEinvoiceMerchantDeleteAPIRequest) SetPlatformCode(_platformCode string) error {
    r._platformCode = _platformCode
    r.Set("platform_code", _platformCode)
    return nil
}

// PlatformCode Getter
func (r AlibabaEinvoiceMerchantDeleteAPIRequest) GetPlatformCode() string {
    return r._platformCode
}
// PayeeRegisterNo Setter
// 税号
func (r *AlibabaEinvoiceMerchantDeleteAPIRequest) SetPayeeRegisterNo(_payeeRegisterNo string) error {
    r._payeeRegisterNo = _payeeRegisterNo
    r.Set("payee_register_no", _payeeRegisterNo)
    return nil
}

// PayeeRegisterNo Getter
func (r AlibabaEinvoiceMerchantDeleteAPIRequest) GetPayeeRegisterNo() string {
    return r._payeeRegisterNo
}
