package nrt

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
家装新零售商品信息查询 API返回值 
tmall.nrt.item.get

查询新零售商品信息
*/
type TmallNrtItemGetAPIResponse struct {
    model.CommonResponse
    TmallNrtItemGetAPIResponseModel
}

// 家装新零售商品信息查询 成功返回结果
type TmallNrtItemGetAPIResponseModel struct {
    XMLName xml.Name `xml:"tmall_nrt_item_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回结果
    TmallNrtItemGet   *TmallNrtItemGetResultDo `json:"tmall_nrt_item_get,omitempty" xml:"tmall_nrt_item_get,omitempty"`
}
