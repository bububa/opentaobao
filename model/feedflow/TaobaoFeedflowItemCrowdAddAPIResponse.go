package feedflow

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFeedflowItemCrowdAddAPIResponse 单品单元下，新增定向人群 API返回值
// taobao.feedflow.item.crowd.add
//
// 单品单元下，新增定向人群
type TaobaoFeedflowItemCrowdAddAPIResponse struct {
	model.CommonResponse
	TaobaoFeedflowItemCrowdAddAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoFeedflowItemCrowdAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoFeedflowItemCrowdAddAPIResponseModel).Reset()
}

// TaobaoFeedflowItemCrowdAddAPIResponseModel is 单品单元下，新增定向人群 成功返回结果
type TaobaoFeedflowItemCrowdAddAPIResponseModel struct {
	XMLName xml.Name `xml:"feedflow_item_crowd_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果对象
	Result *TaobaoFeedflowItemCrowdAddResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoFeedflowItemCrowdAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoFeedflowItemCrowdAddAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoFeedflowItemCrowdAddAPIResponse)
	},
}

// GetTaobaoFeedflowItemCrowdAddAPIResponse 从 sync.Pool 获取 TaobaoFeedflowItemCrowdAddAPIResponse
func GetTaobaoFeedflowItemCrowdAddAPIResponse() *TaobaoFeedflowItemCrowdAddAPIResponse {
	return poolTaobaoFeedflowItemCrowdAddAPIResponse.Get().(*TaobaoFeedflowItemCrowdAddAPIResponse)
}

// ReleaseTaobaoFeedflowItemCrowdAddAPIResponse 将 TaobaoFeedflowItemCrowdAddAPIResponse 保存到 sync.Pool
func ReleaseTaobaoFeedflowItemCrowdAddAPIResponse(v *TaobaoFeedflowItemCrowdAddAPIResponse) {
	v.Reset()
	poolTaobaoFeedflowItemCrowdAddAPIResponse.Put(v)
}
