package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
新增活动详情 API返回值 
taobao.ump.detail.add

增加活动详情。活动详情里面包括活动的范围（店铺，商品），活动的参数（比如具体的折扣），参与类型（全部，部分，部分不参加）等信息。当参与类型为部分或部分不参加的时候需要和taobao.ump.range.add来配合使用。
*/
type TaobaoUmpDetailAddAPIResponse struct {
    model.CommonResponse
    TaobaoUmpDetailAddAPIResponseModel
}

// 新增活动详情 成功返回结果
type TaobaoUmpDetailAddAPIResponseModel struct {
    XMLName xml.Name `xml:"ump_detail_add_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 活动详情的id
    DetailId   int64 `json:"detail_id,omitempty" xml:"detail_id,omitempty"`
}
