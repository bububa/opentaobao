package feedflow

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFeedflowItemCrowdModifybindAPIResponse 修改人群出价或状态 API返回值
// taobao.feedflow.item.crowd.modifybind
//
// 修改人群出价或状态
type TaobaoFeedflowItemCrowdModifybindAPIResponse struct {
	model.CommonResponse
	TaobaoFeedflowItemCrowdModifybindAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoFeedflowItemCrowdModifybindAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoFeedflowItemCrowdModifybindAPIResponseModel).Reset()
}

// TaobaoFeedflowItemCrowdModifybindAPIResponseModel is 修改人群出价或状态 成功返回结果
type TaobaoFeedflowItemCrowdModifybindAPIResponseModel struct {
	XMLName xml.Name `xml:"feedflow_item_crowd_modifybind_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果对象
	Result *TaobaoFeedflowItemCrowdModifybindResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoFeedflowItemCrowdModifybindAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoFeedflowItemCrowdModifybindAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoFeedflowItemCrowdModifybindAPIResponse)
	},
}

// GetTaobaoFeedflowItemCrowdModifybindAPIResponse 从 sync.Pool 获取 TaobaoFeedflowItemCrowdModifybindAPIResponse
func GetTaobaoFeedflowItemCrowdModifybindAPIResponse() *TaobaoFeedflowItemCrowdModifybindAPIResponse {
	return poolTaobaoFeedflowItemCrowdModifybindAPIResponse.Get().(*TaobaoFeedflowItemCrowdModifybindAPIResponse)
}

// ReleaseTaobaoFeedflowItemCrowdModifybindAPIResponse 将 TaobaoFeedflowItemCrowdModifybindAPIResponse 保存到 sync.Pool
func ReleaseTaobaoFeedflowItemCrowdModifybindAPIResponse(v *TaobaoFeedflowItemCrowdModifybindAPIResponse) {
	v.Reset()
	poolTaobaoFeedflowItemCrowdModifybindAPIResponse.Put(v)
}
