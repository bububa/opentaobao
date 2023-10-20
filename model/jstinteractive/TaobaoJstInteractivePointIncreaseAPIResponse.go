package jstinteractive

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJstInteractivePointIncreaseAPIResponse 互动积分发放接口 API返回值
// taobao.jst.interactive.point.increase
//
// 向用户发放互动积分
type TaobaoJstInteractivePointIncreaseAPIResponse struct {
	model.CommonResponse
	TaobaoJstInteractivePointIncreaseAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoJstInteractivePointIncreaseAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoJstInteractivePointIncreaseAPIResponseModel).Reset()
}

// TaobaoJstInteractivePointIncreaseAPIResponseModel is 互动积分发放接口 成功返回结果
type TaobaoJstInteractivePointIncreaseAPIResponseModel struct {
	XMLName xml.Name `xml:"jst_interactive_point_increase_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 用户积分总额
	Amount int64 `json:"amount,omitempty" xml:"amount,omitempty"`
	// 调用是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoJstInteractivePointIncreaseAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Amount = 0
	m.IsSuccess = false
}

var poolTaobaoJstInteractivePointIncreaseAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoJstInteractivePointIncreaseAPIResponse)
	},
}

// GetTaobaoJstInteractivePointIncreaseAPIResponse 从 sync.Pool 获取 TaobaoJstInteractivePointIncreaseAPIResponse
func GetTaobaoJstInteractivePointIncreaseAPIResponse() *TaobaoJstInteractivePointIncreaseAPIResponse {
	return poolTaobaoJstInteractivePointIncreaseAPIResponse.Get().(*TaobaoJstInteractivePointIncreaseAPIResponse)
}

// ReleaseTaobaoJstInteractivePointIncreaseAPIResponse 将 TaobaoJstInteractivePointIncreaseAPIResponse 保存到 sync.Pool
func ReleaseTaobaoJstInteractivePointIncreaseAPIResponse(v *TaobaoJstInteractivePointIncreaseAPIResponse) {
	v.Reset()
	poolTaobaoJstInteractivePointIncreaseAPIResponse.Put(v)
}
