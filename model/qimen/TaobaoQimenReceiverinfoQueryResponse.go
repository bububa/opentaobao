package qimen

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
OAID 收件人信息解密接口 APIResponse
taobao.qimen.receiverinfo.query

WMS 调用该接口，通过 OAID 查询解密后的收件人信息
*/
type TaobaoQimenReceiverinfoQueryAPIResponse struct {
    model.CommonResponse
    TaobaoQimenReceiverinfoQueryResponse
}

type TaobaoQimenReceiverinfoQueryResponse struct {
    XMLName xml.Name `xml:"qimen_receiverinfo_query_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 
    
    Response   *Response `json:"response,omitempty" xml:"response,omitempty"`

    
}
