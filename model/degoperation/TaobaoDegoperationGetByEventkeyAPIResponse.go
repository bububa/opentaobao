package degoperation

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoDegoperationGetByEventkeyAPIResponse 通用用户抽奖次数限制 API返回值
// taobao.degoperation.get.by.eventkey
//
// 通用用户抽奖次数限制
type TaobaoDegoperationGetByEventkeyAPIResponse struct {
	model.CommonResponse
	TaobaoDegoperationGetByEventkeyAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoDegoperationGetByEventkeyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoDegoperationGetByEventkeyAPIResponseModel).Reset()
}

// TaobaoDegoperationGetByEventkeyAPIResponseModel is 通用用户抽奖次数限制 成功返回结果
type TaobaoDegoperationGetByEventkeyAPIResponseModel struct {
	XMLName xml.Name `xml:"degoperation_get_by_eventkey_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *BonusResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoDegoperationGetByEventkeyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoDegoperationGetByEventkeyAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoDegoperationGetByEventkeyAPIResponse)
	},
}

// GetTaobaoDegoperationGetByEventkeyAPIResponse 从 sync.Pool 获取 TaobaoDegoperationGetByEventkeyAPIResponse
func GetTaobaoDegoperationGetByEventkeyAPIResponse() *TaobaoDegoperationGetByEventkeyAPIResponse {
	return poolTaobaoDegoperationGetByEventkeyAPIResponse.Get().(*TaobaoDegoperationGetByEventkeyAPIResponse)
}

// ReleaseTaobaoDegoperationGetByEventkeyAPIResponse 将 TaobaoDegoperationGetByEventkeyAPIResponse 保存到 sync.Pool
func ReleaseTaobaoDegoperationGetByEventkeyAPIResponse(v *TaobaoDegoperationGetByEventkeyAPIResponse) {
	v.Reset()
	poolTaobaoDegoperationGetByEventkeyAPIResponse.Put(v)
}
