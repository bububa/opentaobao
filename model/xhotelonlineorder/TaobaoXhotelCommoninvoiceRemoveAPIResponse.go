package xhotelonlineorder

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
常用发票信息删除接口 API返回值 
taobao.xhotel.commoninvoice.remove

常用发票信息删除接口
*/
type TaobaoXhotelCommoninvoiceRemoveAPIResponse struct {
    model.CommonResponse
    TaobaoXhotelCommoninvoiceRemoveAPIResponseModel
}

// 常用发票信息删除接口 成功返回结果
type TaobaoXhotelCommoninvoiceRemoveAPIResponseModel struct {
    XMLName xml.Name `xml:"xhotel_commoninvoice_remove_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // success
    Issuccess   bool `json:"issuccess,omitempty" xml:"issuccess,omitempty"`
    // errorCode
    Errcode   string `json:"errcode,omitempty" xml:"errcode,omitempty"`
    // errorMsg
    Errmsg   string `json:"errmsg,omitempty" xml:"errmsg,omitempty"`
}
