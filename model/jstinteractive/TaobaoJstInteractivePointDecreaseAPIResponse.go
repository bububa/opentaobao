package jstinteractive

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJstInteractivePointDecreaseAPIResponse 互动积分扣减接口 API返回值
// taobao.jst.interactive.point.decrease
//
// 扣减用户的互动积分
type TaobaoJstInteractivePointDecreaseAPIResponse struct {
	model.CommonResponse
	TaobaoJstInteractivePointDecreaseAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoJstInteractivePointDecreaseAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoJstInteractivePointDecreaseAPIResponseModel).Reset()
}

// TaobaoJstInteractivePointDecreaseAPIResponseModel is 互动积分扣减接口 成功返回结果
type TaobaoJstInteractivePointDecreaseAPIResponseModel struct {
	XMLName xml.Name `xml:"jst_interactive_point_decrease_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 用户积分总额
	Amount int64 `json:"amount,omitempty" xml:"amount,omitempty"`
	// 调用是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoJstInteractivePointDecreaseAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Amount = 0
	m.IsSuccess = false
}

var poolTaobaoJstInteractivePointDecreaseAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoJstInteractivePointDecreaseAPIResponse)
	},
}

// GetTaobaoJstInteractivePointDecreaseAPIResponse 从 sync.Pool 获取 TaobaoJstInteractivePointDecreaseAPIResponse
func GetTaobaoJstInteractivePointDecreaseAPIResponse() *TaobaoJstInteractivePointDecreaseAPIResponse {
	return poolTaobaoJstInteractivePointDecreaseAPIResponse.Get().(*TaobaoJstInteractivePointDecreaseAPIResponse)
}

// ReleaseTaobaoJstInteractivePointDecreaseAPIResponse 将 TaobaoJstInteractivePointDecreaseAPIResponse 保存到 sync.Pool
func ReleaseTaobaoJstInteractivePointDecreaseAPIResponse(v *TaobaoJstInteractivePointDecreaseAPIResponse) {
	v.Reset()
	poolTaobaoJstInteractivePointDecreaseAPIResponse.Put(v)
}
