package promotion

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoumprangedeleteAPIResponse 删除活动范围 API返回值
// taobao.ump.range.delete
//
// 去指先前指定在某项活动中，某些类型的物品参加或者不参加活动的设置
type TaobaoumprangedeleteAPIResponse struct {
	model.CommonResponse
	TaobaoumprangedeleteAPIResponseModel
}

// TaobaoumprangedeleteAPIResponseModel is 删除活动范围 成功返回结果
type TaobaoumprangedeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"ump_range_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 调用是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
