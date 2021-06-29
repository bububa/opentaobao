package einvoice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
ERP开票结果获取 API请求
alibaba.einvoice.create.result.get

ERP开票结果获取
*/
type AlibabaEinvoiceCreateResultGetRequest struct {
    model.Params
    // 流水号 (serial_no)和(platform_code,platform_tid)必须填写其中一组,serial_no优先级更高
    _serialNo   string
    // 电商平台代码。淘宝：taobao，天猫：tmall
    _platformCode   string
    // 电商平台对应的订单号
    _platformTid   string
    // 收款方税务登记证号
    _payeeRegisterNo   string
    // 外部平台店铺名称，需要在阿里发票平台配置，只有当platform_code不为TB和TM时，这个字段才生效。注意：后台配置的店铺平台必须和入参platform_code一致
    _outShopName   string
}

// 初始化AlibabaEinvoiceCreateResultGetRequest对象
func NewAlibabaEinvoiceCreateResultGetRequest() *AlibabaEinvoiceCreateResultGetRequest{
    return &AlibabaEinvoiceCreateResultGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaEinvoiceCreateResultGetRequest) GetApiMethodName() string {
    return "alibaba.einvoice.create.result.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaEinvoiceCreateResultGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SerialNo Setter
// 流水号 (serial_no)和(platform_code,platform_tid)必须填写其中一组,serial_no优先级更高
func (r *AlibabaEinvoiceCreateResultGetRequest) SetSerialNo(_serialNo string) error {
    r._serialNo = _serialNo
    r.Set("serial_no", _serialNo)
    return nil
}

// SerialNo Getter
func (r AlibabaEinvoiceCreateResultGetRequest) GetSerialNo() string {
    return r._serialNo
}
// PlatformCode Setter
// 电商平台代码。淘宝：taobao，天猫：tmall
func (r *AlibabaEinvoiceCreateResultGetRequest) SetPlatformCode(_platformCode string) error {
    r._platformCode = _platformCode
    r.Set("platform_code", _platformCode)
    return nil
}

// PlatformCode Getter
func (r AlibabaEinvoiceCreateResultGetRequest) GetPlatformCode() string {
    return r._platformCode
}
// PlatformTid Setter
// 电商平台对应的订单号
func (r *AlibabaEinvoiceCreateResultGetRequest) SetPlatformTid(_platformTid string) error {
    r._platformTid = _platformTid
    r.Set("platform_tid", _platformTid)
    return nil
}

// PlatformTid Getter
func (r AlibabaEinvoiceCreateResultGetRequest) GetPlatformTid() string {
    return r._platformTid
}
// PayeeRegisterNo Setter
// 收款方税务登记证号
func (r *AlibabaEinvoiceCreateResultGetRequest) SetPayeeRegisterNo(_payeeRegisterNo string) error {
    r._payeeRegisterNo = _payeeRegisterNo
    r.Set("payee_register_no", _payeeRegisterNo)
    return nil
}

// PayeeRegisterNo Getter
func (r AlibabaEinvoiceCreateResultGetRequest) GetPayeeRegisterNo() string {
    return r._payeeRegisterNo
}
// OutShopName Setter
// 外部平台店铺名称，需要在阿里发票平台配置，只有当platform_code不为TB和TM时，这个字段才生效。注意：后台配置的店铺平台必须和入参platform_code一致
func (r *AlibabaEinvoiceCreateResultGetRequest) SetOutShopName(_outShopName string) error {
    r._outShopName = _outShopName
    r.Set("out_shop_name", _outShopName)
    return nil
}

// OutShopName Getter
func (r AlibabaEinvoiceCreateResultGetRequest) GetOutShopName() string {
    return r._outShopName
}
