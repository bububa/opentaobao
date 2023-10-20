package feedflow

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFeedflowItemCreativeDeleteAPIResponse 信息流删除创意 API返回值
// taobao.feedflow.item.creative.delete
//
// 信息流删除创意
type TaobaoFeedflowItemCreativeDeleteAPIResponse struct {
	model.CommonResponse
	TaobaoFeedflowItemCreativeDeleteAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoFeedflowItemCreativeDeleteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoFeedflowItemCreativeDeleteAPIResponseModel).Reset()
}

// TaobaoFeedflowItemCreativeDeleteAPIResponseModel is 信息流删除创意 成功返回结果
type TaobaoFeedflowItemCreativeDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"feedflow_item_creative_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果对下
	Result *TaobaoFeedflowItemCreativeDeleteResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoFeedflowItemCreativeDeleteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoFeedflowItemCreativeDeleteAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoFeedflowItemCreativeDeleteAPIResponse)
	},
}

// GetTaobaoFeedflowItemCreativeDeleteAPIResponse 从 sync.Pool 获取 TaobaoFeedflowItemCreativeDeleteAPIResponse
func GetTaobaoFeedflowItemCreativeDeleteAPIResponse() *TaobaoFeedflowItemCreativeDeleteAPIResponse {
	return poolTaobaoFeedflowItemCreativeDeleteAPIResponse.Get().(*TaobaoFeedflowItemCreativeDeleteAPIResponse)
}

// ReleaseTaobaoFeedflowItemCreativeDeleteAPIResponse 将 TaobaoFeedflowItemCreativeDeleteAPIResponse 保存到 sync.Pool
func ReleaseTaobaoFeedflowItemCreativeDeleteAPIResponse(v *TaobaoFeedflowItemCreativeDeleteAPIResponse) {
	v.Reset()
	poolTaobaoFeedflowItemCreativeDeleteAPIResponse.Put(v)
}
