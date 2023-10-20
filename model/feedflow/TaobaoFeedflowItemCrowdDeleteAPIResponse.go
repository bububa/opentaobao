package feedflow

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFeedflowItemCrowdDeleteAPIResponse 删除单品人群 API返回值
// taobao.feedflow.item.crowd.delete
//
// 删除单品人群
type TaobaoFeedflowItemCrowdDeleteAPIResponse struct {
	model.CommonResponse
	TaobaoFeedflowItemCrowdDeleteAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoFeedflowItemCrowdDeleteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoFeedflowItemCrowdDeleteAPIResponseModel).Reset()
}

// TaobaoFeedflowItemCrowdDeleteAPIResponseModel is 删除单品人群 成功返回结果
type TaobaoFeedflowItemCrowdDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"feedflow_item_crowd_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果对象
	Result *TaobaoFeedflowItemCrowdDeleteResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoFeedflowItemCrowdDeleteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoFeedflowItemCrowdDeleteAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoFeedflowItemCrowdDeleteAPIResponse)
	},
}

// GetTaobaoFeedflowItemCrowdDeleteAPIResponse 从 sync.Pool 获取 TaobaoFeedflowItemCrowdDeleteAPIResponse
func GetTaobaoFeedflowItemCrowdDeleteAPIResponse() *TaobaoFeedflowItemCrowdDeleteAPIResponse {
	return poolTaobaoFeedflowItemCrowdDeleteAPIResponse.Get().(*TaobaoFeedflowItemCrowdDeleteAPIResponse)
}

// ReleaseTaobaoFeedflowItemCrowdDeleteAPIResponse 将 TaobaoFeedflowItemCrowdDeleteAPIResponse 保存到 sync.Pool
func ReleaseTaobaoFeedflowItemCrowdDeleteAPIResponse(v *TaobaoFeedflowItemCrowdDeleteAPIResponse) {
	v.Reset()
	poolTaobaoFeedflowItemCrowdDeleteAPIResponse.Put(v)
}
