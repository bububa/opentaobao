package xhotel

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelDataServiceOrderDetailAPIResponse 服务订单详情 API返回值
// taobao.xhotel.data.service.order.detail
//
// 服务订单详情top接口构建
type TaobaoXhotelDataServiceOrderDetailAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelDataServiceOrderDetailAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoXhotelDataServiceOrderDetailAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoXhotelDataServiceOrderDetailAPIResponseModel).Reset()
}

// TaobaoXhotelDataServiceOrderDetailAPIResponseModel is 服务订单详情 成功返回结果
type TaobaoXhotelDataServiceOrderDetailAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_data_service_order_detail_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaoXhotelDataServiceOrderDetailResultSet `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoXhotelDataServiceOrderDetailAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoXhotelDataServiceOrderDetailAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelDataServiceOrderDetailAPIResponse)
	},
}

// GetTaobaoXhotelDataServiceOrderDetailAPIResponse 从 sync.Pool 获取 TaobaoXhotelDataServiceOrderDetailAPIResponse
func GetTaobaoXhotelDataServiceOrderDetailAPIResponse() *TaobaoXhotelDataServiceOrderDetailAPIResponse {
	return poolTaobaoXhotelDataServiceOrderDetailAPIResponse.Get().(*TaobaoXhotelDataServiceOrderDetailAPIResponse)
}

// ReleaseTaobaoXhotelDataServiceOrderDetailAPIResponse 将 TaobaoXhotelDataServiceOrderDetailAPIResponse 保存到 sync.Pool
func ReleaseTaobaoXhotelDataServiceOrderDetailAPIResponse(v *TaobaoXhotelDataServiceOrderDetailAPIResponse) {
	v.Reset()
	poolTaobaoXhotelDataServiceOrderDetailAPIResponse.Put(v)
}
