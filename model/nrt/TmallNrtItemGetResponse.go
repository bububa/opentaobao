package nrt

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
家装新零售商品信息查询 APIResponse
tmall.nrt.item.get

查询新零售商品信息
*/
type TmallNrtItemGetAPIResponse struct {
    model.CommonResponse
    TmallNrtItemGetResponse
}

type TmallNrtItemGetResponse struct {
    XMLName xml.Name `xml:"tmall_nrt_item_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    TmallNrtItemGet   *TmallNrtItemGetResultDo `json:"tmall_nrt_item_get,omitempty" xml:"tmall_nrt_item_get,omitempty"`

    
}
