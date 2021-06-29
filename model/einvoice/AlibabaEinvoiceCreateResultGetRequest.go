package einvoice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
ERP开票结果获取 APIRequest
alibaba.einvoice.create.result.get

ERP开票结果获取
*/
type AlibabaEinvoiceCreateResultGetRequest struct {
    model.Params

    // 流水号 (serial_no)和(platform_code,platform_tid)必须填写其中一组,serial_no优先级更高
    serialNo   string 

    // 电商平台代码。淘宝：taobao，天猫：tmall
    platformCode   string 

    // 电商平台对应的订单号
    platformTid   string 

    // 收款方税务登记证号
    payeeRegisterNo   string 

    // 外部平台店铺名称，需要在阿里发票平台配置，只有当platform_code不为TB和TM时，这个字段才生效。注意：后台配置的店铺平台必须和入参platform_code一致
    outShopName   string 

}

func NewAlibabaEinvoiceCreateResultGetRequest() *AlibabaEinvoiceCreateResultGetRequest{
    return &AlibabaEinvoiceCreateResultGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaEinvoiceCreateResultGetRequest) GetApiMethodName() string {
    return "alibaba.einvoice.create.result.get"
}

func (r AlibabaEinvoiceCreateResultGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaEinvoiceCreateResultGetRequest) SetSerialNo(serialNo string) error {
    r.serialNo = serialNo
    r.Set("serial_no", serialNo)
    return nil
}

func (r AlibabaEinvoiceCreateResultGetRequest) GetSerialNo() string {
    return r.serialNo
}

func (r *AlibabaEinvoiceCreateResultGetRequest) SetPlatformCode(platformCode string) error {
    r.platformCode = platformCode
    r.Set("platform_code", platformCode)
    return nil
}

func (r AlibabaEinvoiceCreateResultGetRequest) GetPlatformCode() string {
    return r.platformCode
}

func (r *AlibabaEinvoiceCreateResultGetRequest) SetPlatformTid(platformTid string) error {
    r.platformTid = platformTid
    r.Set("platform_tid", platformTid)
    return nil
}

func (r AlibabaEinvoiceCreateResultGetRequest) GetPlatformTid() string {
    return r.platformTid
}

func (r *AlibabaEinvoiceCreateResultGetRequest) SetPayeeRegisterNo(payeeRegisterNo string) error {
    r.payeeRegisterNo = payeeRegisterNo
    r.Set("payee_register_no", payeeRegisterNo)
    return nil
}

func (r AlibabaEinvoiceCreateResultGetRequest) GetPayeeRegisterNo() string {
    return r.payeeRegisterNo
}

func (r *AlibabaEinvoiceCreateResultGetRequest) SetOutShopName(outShopName string) error {
    r.outShopName = outShopName
    r.Set("out_shop_name", outShopName)
    return nil
}

func (r AlibabaEinvoiceCreateResultGetRequest) GetOutShopName() string {
    return r.outShopName
}

