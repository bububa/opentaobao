package qimen

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
组合商品接口 APIResponse
taobao.qimen.combineitem.synchronize

ERP调用奇门的接口,将商品信息同步给WMS
*/
type TaobaoQimenCombineitemSynchronizeAPIResponse struct {
    model.CommonResponse
    TaobaoQimenCombineitemSynchronizeResponse
}

type TaobaoQimenCombineitemSynchronizeResponse struct {
    XMLName xml.Name `xml:"qimen_combineitem_synchronize_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 
    
    Response   *Response `json:"response,omitempty" xml:"response,omitempty"`

    
}
