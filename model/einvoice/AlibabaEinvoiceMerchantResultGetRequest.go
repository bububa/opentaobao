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
    _serialNo   string
    // 电商平台代码。淘宝：taobao，天猫：tmall
    _platformCode   string
    // 电商平台对应的订单号
    _platformTid   string
    // 收款方税务登记证号
    _payeeRegisterNo   string
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
func (r *AlibabaEinvoiceMerchantResultGetRequest) SetSerialNo(_serialNo string) error {
    r._serialNo = _serialNo
    r.Set("serial_no", _serialNo)
    return nil
}

// SerialNo Getter
func (r AlibabaEinvoiceMerchantResultGetRequest) GetSerialNo() string {
    return r._serialNo
}
// PlatformCode Setter
// 电商平台代码。淘宝：taobao，天猫：tmall
func (r *AlibabaEinvoiceMerchantResultGetRequest) SetPlatformCode(_platformCode string) error {
    r._platformCode = _platformCode
    r.Set("platform_code", _platformCode)
    return nil
}

// PlatformCode Getter
func (r AlibabaEinvoiceMerchantResultGetRequest) GetPlatformCode() string {
    return r._platformCode
}
// PlatformTid Setter
// 电商平台对应的订单号
func (r *AlibabaEinvoiceMerchantResultGetRequest) SetPlatformTid(_platformTid string) error {
    r._platformTid = _platformTid
    r.Set("platform_tid", _platformTid)
    return nil
}

// PlatformTid Getter
func (r AlibabaEinvoiceMerchantResultGetRequest) GetPlatformTid() string {
    return r._platformTid
}
// PayeeRegisterNo Setter
// 收款方税务登记证号
func (r *AlibabaEinvoiceMerchantResultGetRequest) SetPayeeRegisterNo(_payeeRegisterNo string) error {
    r._payeeRegisterNo = _payeeRegisterNo
    r.Set("payee_register_no", _payeeRegisterNo)
    return nil
}

// PayeeRegisterNo Getter
func (r AlibabaEinvoiceMerchantResultGetRequest) GetPayeeRegisterNo() string {
    return r._payeeRegisterNo
}
