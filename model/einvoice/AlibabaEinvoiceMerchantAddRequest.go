package einvoice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
发票中台-同平台授权税号适用商户 APIRequest
alibaba.einvoice.merchant.add

适用于以下场景：
业务税号入驻成功后，需要将税号授权给同平台下其他商户，使得其他商户也具备开票能力
*/
type AlibabaEinvoiceMerchantAddRequest struct {
    model.Params

    // 验证码，门店绑定已入驻税号接口返回的taxToken
    taxToken   string 

    // 业务方发起新增门店的唯一幂等ID, 由业务方生成。只能由字母和数字组成。
    outerId   string 

    // 业务平台门店ID
    merchantUserId   string 

    // 业务平台code, 由阿里发票分配
    platformCode   string 

    // 税务登记号
    payeeRegisterNo   string 

    // 业务平台门店名称
    merchantName   string 

    // 税盘列表
    deviceIds   []string 

}

func NewAlibabaEinvoiceMerchantAddRequest() *AlibabaEinvoiceMerchantAddRequest{
    return &AlibabaEinvoiceMerchantAddRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaEinvoiceMerchantAddRequest) GetApiMethodName() string {
    return "alibaba.einvoice.merchant.add"
}

func (r AlibabaEinvoiceMerchantAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaEinvoiceMerchantAddRequest) SetTaxToken(taxToken string) error {
    r.taxToken = taxToken
    r.Set("tax_token", taxToken)
    return nil
}

func (r AlibabaEinvoiceMerchantAddRequest) GetTaxToken() string {
    return r.taxToken
}

func (r *AlibabaEinvoiceMerchantAddRequest) SetOuterId(outerId string) error {
    r.outerId = outerId
    r.Set("outer_id", outerId)
    return nil
}

func (r AlibabaEinvoiceMerchantAddRequest) GetOuterId() string {
    return r.outerId
}

func (r *AlibabaEinvoiceMerchantAddRequest) SetMerchantUserId(merchantUserId string) error {
    r.merchantUserId = merchantUserId
    r.Set("merchant_user_id", merchantUserId)
    return nil
}

func (r AlibabaEinvoiceMerchantAddRequest) GetMerchantUserId() string {
    return r.merchantUserId
}

func (r *AlibabaEinvoiceMerchantAddRequest) SetPlatformCode(platformCode string) error {
    r.platformCode = platformCode
    r.Set("platform_code", platformCode)
    return nil
}

func (r AlibabaEinvoiceMerchantAddRequest) GetPlatformCode() string {
    return r.platformCode
}

func (r *AlibabaEinvoiceMerchantAddRequest) SetPayeeRegisterNo(payeeRegisterNo string) error {
    r.payeeRegisterNo = payeeRegisterNo
    r.Set("payee_register_no", payeeRegisterNo)
    return nil
}

func (r AlibabaEinvoiceMerchantAddRequest) GetPayeeRegisterNo() string {
    return r.payeeRegisterNo
}

func (r *AlibabaEinvoiceMerchantAddRequest) SetMerchantName(merchantName string) error {
    r.merchantName = merchantName
    r.Set("merchant_name", merchantName)
    return nil
}

func (r AlibabaEinvoiceMerchantAddRequest) GetMerchantName() string {
    return r.merchantName
}

func (r *AlibabaEinvoiceMerchantAddRequest) SetDeviceIds(deviceIds []string) error {
    r.deviceIds = deviceIds
    r.Set("device_ids", deviceIds)
    return nil
}

func (r AlibabaEinvoiceMerchantAddRequest) GetDeviceIds() []string {
    return r.deviceIds
}

