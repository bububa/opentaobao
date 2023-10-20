package feedflow

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFeedflowItemTargetValidlistAPIResponse 获取有权限的定向列表 API返回值
// taobao.feedflow.item.target.validlist
//
// 获取有权限的定向列表
type TaobaoFeedflowItemTargetValidlistAPIResponse struct {
	model.CommonResponse
	TaobaoFeedflowItemTargetValidlistAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoFeedflowItemTargetValidlistAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoFeedflowItemTargetValidlistAPIResponseModel).Reset()
}

// TaobaoFeedflowItemTargetValidlistAPIResponseModel is 获取有权限的定向列表 成功返回结果
type TaobaoFeedflowItemTargetValidlistAPIResponseModel struct {
	XMLName xml.Name `xml:"feedflow_item_target_validlist_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果对象
	Result *TaobaoFeedflowItemTargetValidlistResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoFeedflowItemTargetValidlistAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoFeedflowItemTargetValidlistAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoFeedflowItemTargetValidlistAPIResponse)
	},
}

// GetTaobaoFeedflowItemTargetValidlistAPIResponse 从 sync.Pool 获取 TaobaoFeedflowItemTargetValidlistAPIResponse
func GetTaobaoFeedflowItemTargetValidlistAPIResponse() *TaobaoFeedflowItemTargetValidlistAPIResponse {
	return poolTaobaoFeedflowItemTargetValidlistAPIResponse.Get().(*TaobaoFeedflowItemTargetValidlistAPIResponse)
}

// ReleaseTaobaoFeedflowItemTargetValidlistAPIResponse 将 TaobaoFeedflowItemTargetValidlistAPIResponse 保存到 sync.Pool
func ReleaseTaobaoFeedflowItemTargetValidlistAPIResponse(v *TaobaoFeedflowItemTargetValidlistAPIResponse) {
	v.Reset()
	poolTaobaoFeedflowItemTargetValidlistAPIResponse.Put(v)
}
