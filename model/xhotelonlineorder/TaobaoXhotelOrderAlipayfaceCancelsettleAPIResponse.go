package xhotelonlineorder

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelOrderAlipayfaceCancelsettleAPIResponse 信用住订单取消结算接口 API返回值
// taobao.xhotel.order.alipayface.cancelsettle
//
// 信用住订单由于客人为出现等原因，最终取消结算,一定要在结算后2个小时之内调用，否则不会成功。
type TaobaoXhotelOrderAlipayfaceCancelsettleAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelOrderAlipayfaceCancelsettleAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoXhotelOrderAlipayfaceCancelsettleAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoXhotelOrderAlipayfaceCancelsettleAPIResponseModel).Reset()
}

// TaobaoXhotelOrderAlipayfaceCancelsettleAPIResponseModel is 信用住订单取消结算接口 成功返回结果
type TaobaoXhotelOrderAlipayfaceCancelsettleAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_order_alipayface_cancelsettle_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回信息
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoXhotelOrderAlipayfaceCancelsettleAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = ""
}

var poolTaobaoXhotelOrderAlipayfaceCancelsettleAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelOrderAlipayfaceCancelsettleAPIResponse)
	},
}

// GetTaobaoXhotelOrderAlipayfaceCancelsettleAPIResponse 从 sync.Pool 获取 TaobaoXhotelOrderAlipayfaceCancelsettleAPIResponse
func GetTaobaoXhotelOrderAlipayfaceCancelsettleAPIResponse() *TaobaoXhotelOrderAlipayfaceCancelsettleAPIResponse {
	return poolTaobaoXhotelOrderAlipayfaceCancelsettleAPIResponse.Get().(*TaobaoXhotelOrderAlipayfaceCancelsettleAPIResponse)
}

// ReleaseTaobaoXhotelOrderAlipayfaceCancelsettleAPIResponse 将 TaobaoXhotelOrderAlipayfaceCancelsettleAPIResponse 保存到 sync.Pool
func ReleaseTaobaoXhotelOrderAlipayfaceCancelsettleAPIResponse(v *TaobaoXhotelOrderAlipayfaceCancelsettleAPIResponse) {
	v.Reset()
	poolTaobaoXhotelOrderAlipayfaceCancelsettleAPIResponse.Put(v)
}
