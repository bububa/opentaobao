package einvoice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商家自研ERP开票结果获取 API请求
alibaba.einvoice.merchant.result.get

商家自研ERP开票结果获取
*/
type AlibabaEinvoiceMerchantResultGetRequest struct {
    model.Params
    // 流水号 (serial_no)和(platform_code,platform_tid)必须填写其中一组,serial_no优先级更高
    serialNo   string
    // 电商平台代码。淘宝：taobao，天猫：tmall
    platformCode   string
    // 电商平台对应的订单号
    platformTid   string
    // 收款方税务登记证号
    payeeRegisterNo   string
}

// 初始化AlibabaEinvoiceMerchantResultGetRequest对象
func NewAlibabaEinvoiceMerchantResultGetRequest() *AlibabaEinvoiceMerchantResultGetRequest{
    return &AlibabaEinvoiceMerchantResultGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaEinvoiceMerchantResultGetRequest) GetApiMethodName() string {
    return "alibaba.einvoice.merchant.result.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaEinvoiceMerchantResultGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SerialNo Setter
// 流水号 (serial_no)和(platform_code,platform_tid)必须填写其中一组,serial_no优先级更高
func (r *AlibabaEinvoiceMerchantResultGetRequest) SetSerialNo(serialNo string) error {
    r.serialNo = serialNo
    r.Set("serial_no", serialNo)
    return nil
}

// SerialNo Getter
func (r AlibabaEinvoiceMerchantResultGetRequest) GetSerialNo() string {
    return r.serialNo
}
// PlatformCode Setter
// 电商平台代码。淘宝：taobao，天猫：tmall
func (r *AlibabaEinvoiceMerchantResultGetRequest) SetPlatformCode(platformCode string) error {
    r.platformCode = platformCode
    r.Set("platform_code", platformCode)
    return nil
}

// PlatformCode Getter
func (r AlibabaEinvoiceMerchantResultGetRequest) GetPlatformCode() string {
    return r.platformCode
}
// PlatformTid Setter
// 电商平台对应的订单号
func (r *AlibabaEinvoiceMerchantResultGetRequest) SetPlatformTid(platformTid string) error {
    r.platformTid = platformTid
    r.Set("platform_tid", platformTid)
    return nil
}

// PlatformTid Getter
func (r AlibabaEinvoiceMerchantResultGetRequest) GetPlatformTid() string {
    return r.platformTid
}
// PayeeRegisterNo Setter
// 收款方税务登记证号
func (r *AlibabaEinvoiceMerchantResultGetRequest) SetPayeeRegisterNo(payeeRegisterNo string) error {
    r.payeeRegisterNo = payeeRegisterNo
    r.Set("payee_register_no", payeeRegisterNo)
    return nil
}

// PayeeRegisterNo Getter
func (r AlibabaEinvoiceMerchantResultGetRequest) GetPayeeRegisterNo() string {
    return r.payeeRegisterNo
}
