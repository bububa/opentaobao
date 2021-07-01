package xhotelonlineorder

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
常用发票信息更新接口 API返回值 
taobao.xhotel.commoninvoice.update

常用发票信息更新接口(根据用户id,发票抬头和发票属性或发票id进行更新,没有则添加)
*/
type TaobaoXhotelCommoninvoiceUpdateAPIResponse struct {
    model.CommonResponse
    TaobaoXhotelCommoninvoiceUpdateAPIResponseModel
}

// 常用发票信息更新接口 成功返回结果
type TaobaoXhotelCommoninvoiceUpdateAPIResponseModel struct {
    XMLName xml.Name `xml:"xhotel_commoninvoice_update_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // success
    Issuccess   bool `json:"issuccess,omitempty" xml:"issuccess,omitempty"`
    // errorCode
    Errcode   string `json:"errcode,omitempty" xml:"errcode,omitempty"`
    // errorMsg
    Errmsg   string `json:"errmsg,omitempty" xml:"errmsg,omitempty"`
}
