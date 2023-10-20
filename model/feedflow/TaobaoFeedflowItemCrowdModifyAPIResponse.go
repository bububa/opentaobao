package feedflow

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFeedflowItemCrowdModifyAPIResponse 覆盖单元下同类型定向人群 API返回值
// taobao.feedflow.item.crowd.modify
//
// 覆盖单元下同类型定向人群
type TaobaoFeedflowItemCrowdModifyAPIResponse struct {
	model.CommonResponse
	TaobaoFeedflowItemCrowdModifyAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoFeedflowItemCrowdModifyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoFeedflowItemCrowdModifyAPIResponseModel).Reset()
}

// TaobaoFeedflowItemCrowdModifyAPIResponseModel is 覆盖单元下同类型定向人群 成功返回结果
type TaobaoFeedflowItemCrowdModifyAPIResponseModel struct {
	XMLName xml.Name `xml:"feedflow_item_crowd_modify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果对象
	Result *TaobaoFeedflowItemCrowdModifyResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoFeedflowItemCrowdModifyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoFeedflowItemCrowdModifyAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoFeedflowItemCrowdModifyAPIResponse)
	},
}

// GetTaobaoFeedflowItemCrowdModifyAPIResponse 从 sync.Pool 获取 TaobaoFeedflowItemCrowdModifyAPIResponse
func GetTaobaoFeedflowItemCrowdModifyAPIResponse() *TaobaoFeedflowItemCrowdModifyAPIResponse {
	return poolTaobaoFeedflowItemCrowdModifyAPIResponse.Get().(*TaobaoFeedflowItemCrowdModifyAPIResponse)
}

// ReleaseTaobaoFeedflowItemCrowdModifyAPIResponse 将 TaobaoFeedflowItemCrowdModifyAPIResponse 保存到 sync.Pool
func ReleaseTaobaoFeedflowItemCrowdModifyAPIResponse(v *TaobaoFeedflowItemCrowdModifyAPIResponse) {
	v.Reset()
	poolTaobaoFeedflowItemCrowdModifyAPIResponse.Put(v)
}
