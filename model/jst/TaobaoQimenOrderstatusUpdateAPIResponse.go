package jst

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenOrderstatusUpdateAPIResponse 订单状态更新接口 API返回值
// taobao.qimen.orderstatus.update
//
// 星盘和ISV，可以通过此接口，来更新订单状态。此接口应用于使用阿里星盘分单，且使用商家系统（非阿里掌柜）接单/拒单的模式下更新订单状态。
type TaobaoQimenOrderstatusUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoQimenOrderstatusUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoQimenOrderstatusUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoQimenOrderstatusUpdateAPIResponseModel).Reset()
}

// TaobaoQimenOrderstatusUpdateAPIResponseModel is 订单状态更新接口 成功返回结果
type TaobaoQimenOrderstatusUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_orderstatus_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// resultCode
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// success
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoQimenOrderstatusUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Message = ""
	m.ResultCode = ""
	m.IsSuccess = false
}

var poolTaobaoQimenOrderstatusUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQimenOrderstatusUpdateAPIResponse)
	},
}

// GetTaobaoQimenOrderstatusUpdateAPIResponse 从 sync.Pool 获取 TaobaoQimenOrderstatusUpdateAPIResponse
func GetTaobaoQimenOrderstatusUpdateAPIResponse() *TaobaoQimenOrderstatusUpdateAPIResponse {
	return poolTaobaoQimenOrderstatusUpdateAPIResponse.Get().(*TaobaoQimenOrderstatusUpdateAPIResponse)
}

// ReleaseTaobaoQimenOrderstatusUpdateAPIResponse 将 TaobaoQimenOrderstatusUpdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoQimenOrderstatusUpdateAPIResponse(v *TaobaoQimenOrderstatusUpdateAPIResponse) {
	v.Reset()
	poolTaobaoQimenOrderstatusUpdateAPIResponse.Put(v)
}
