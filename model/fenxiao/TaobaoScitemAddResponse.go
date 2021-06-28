package fenxiao

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
发布后端商品 APIResponse
taobao.scitem.add

发布后端商品
*/
type TaobaoScitemAddAPIResponse struct {
    model.CommonResponse
    TaobaoScitemAddResponse
}

type TaobaoScitemAddResponse struct {
    XMLName xml.Name `xml:"scitem_add_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 后台商品信息
    
    ScItem   *ScItem `json:"sc_item,omitempty" xml:"sc_item,omitempty"`

    
}
