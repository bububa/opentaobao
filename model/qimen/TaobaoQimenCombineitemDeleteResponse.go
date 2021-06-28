package qimen

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
组合货品删除接口 APIResponse
taobao.qimen.combineitem.delete

组合货品删除
*/
type TaobaoQimenCombineitemDeleteAPIResponse struct {
    model.CommonResponse
    TaobaoQimenCombineitemDeleteResponse
}

type TaobaoQimenCombineitemDeleteResponse struct {
    XMLName xml.Name `xml:"qimen_combineitem_delete_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 
    
    Response   *ResponseDO `json:"response,omitempty" xml:"response,omitempty"`

    
}
