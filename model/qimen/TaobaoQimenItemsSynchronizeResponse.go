package qimen

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商品同步接口 (批量) APIResponse
taobao.qimen.items.synchronize

ERP调用奇门的接口,批量同步商品信息给WMS
*/
type TaobaoQimenItemsSynchronizeAPIResponse struct {
    model.CommonResponse
    TaobaoQimenItemsSynchronizeResponse
}

type TaobaoQimenItemsSynchronizeResponse struct {
    XMLName xml.Name `xml:"qimen_items_synchronize_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 
    
    Response   *ItemsSynchronizeResponse `json:"response,omitempty" xml:"response,omitempty"`

    
}
