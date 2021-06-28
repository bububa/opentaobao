package shop

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取商品 APIResponse
alibaba.data.item.get

获取商品信息，作为客户端Weex鉴权的虚拟api
*/
type AlibabaDataItemGetAPIResponse struct {
    model.CommonResponse
    AlibabaDataItemGetResponse
}

type AlibabaDataItemGetResponse struct {
    XMLName xml.Name `xml:"alibaba_data_item_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 获取商品信息，作为客户端Weex鉴权的虚拟api
    
    Unnamed   string `json:"unnamed,omitempty" xml:"unnamed,omitempty"`

    
}
