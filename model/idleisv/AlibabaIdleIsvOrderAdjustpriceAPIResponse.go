package idleisv

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleIsvOrderAdjustpriceAPIResponse 闲鱼服务商订单价格修改接口 API返回值
// alibaba.idle.isv.order.adjustprice
//
// 闲鱼用户通过授权的服务商修改订单价格和邮费
type AlibabaIdleIsvOrderAdjustpriceAPIResponse struct {
	model.CommonResponse
	AlibabaIdleIsvOrderAdjustpriceAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIdleIsvOrderAdjustpriceAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIdleIsvOrderAdjustpriceAPIResponseModel).Reset()
}

// AlibabaIdleIsvOrderAdjustpriceAPIResponseModel is 闲鱼服务商订单价格修改接口 成功返回结果
type AlibabaIdleIsvOrderAdjustpriceAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_isv_order_adjustprice_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回结果
	Result *AlibabaIdleIsvOrderAdjustpriceResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIdleIsvOrderAdjustpriceAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaIdleIsvOrderAdjustpriceAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIdleIsvOrderAdjustpriceAPIResponse)
	},
}

// GetAlibabaIdleIsvOrderAdjustpriceAPIResponse 从 sync.Pool 获取 AlibabaIdleIsvOrderAdjustpriceAPIResponse
func GetAlibabaIdleIsvOrderAdjustpriceAPIResponse() *AlibabaIdleIsvOrderAdjustpriceAPIResponse {
	return poolAlibabaIdleIsvOrderAdjustpriceAPIResponse.Get().(*AlibabaIdleIsvOrderAdjustpriceAPIResponse)
}

// ReleaseAlibabaIdleIsvOrderAdjustpriceAPIResponse 将 AlibabaIdleIsvOrderAdjustpriceAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIdleIsvOrderAdjustpriceAPIResponse(v *AlibabaIdleIsvOrderAdjustpriceAPIResponse) {
	v.Reset()
	poolAlibabaIdleIsvOrderAdjustpriceAPIResponse.Put(v)
}
