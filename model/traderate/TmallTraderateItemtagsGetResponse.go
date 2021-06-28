package traderate

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
通过商品ID获取标签列表 APIResponse
tmall.traderate.itemtags.get

通过商品ID获取标签详细信息
*/
type TmallTraderateItemtagsGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"tmall_traderate_itemtags_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 标签列表
    
    Tags   []TmallRateTagDetail `json:"tags,omitempty" xml:"