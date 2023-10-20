package xhotelonlineorder

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelOrderDetailSearchAPIResponse 订单详情查询 API返回值
// taobao.xhotel.order.detail.search
//
// 提供订单详情查询
type TaobaoXhotelOrderDetailSearchAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelOrderDetailSearchAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoXhotelOrderDetailSearchAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoXhotelOrderDetailSearchAPIResponseModel).Reset()
}

// TaobaoXhotelOrderDetailSearchAPIResponseModel is 订单详情查询 成功返回结果
type TaobaoXhotelOrderDetailSearchAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_order_detail_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误编号
	Error string `json:"error,omitempty" xml:"error,omitempty"`
	// 订单详情对象
	TopOrderDetail *TopOrderDetail `json:"top_order_detail,omitempty" xml:"top_order_detail,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoXhotelOrderDetailSearchAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Error = ""
	m.TopOrderDetail = nil
}

var poolTaobaoXhotelOrderDetailSearchAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelOrderDetailSearchAPIResponse)
	},
}

// GetTaobaoXhotelOrderDetailSearchAPIResponse 从 sync.Pool 获取 TaobaoXhotelOrderDetailSearchAPIResponse
func GetTaobaoXhotelOrderDetailSearchAPIResponse() *TaobaoXhotelOrderDetailSearchAPIResponse {
	return poolTaobaoXhotelOrderDetailSearchAPIResponse.Get().(*TaobaoXhotelOrderDetailSearchAPIResponse)
}

// ReleaseTaobaoXhotelOrderDetailSearchAPIResponse 将 TaobaoXhotelOrderDetailSearchAPIResponse 保存到 sync.Pool
func ReleaseTaobaoXhotelOrderDetailSearchAPIResponse(v *TaobaoXhotelOrderDetailSearchAPIResponse) {
	v.Reset()
	poolTaobaoXhotelOrderDetailSearchAPIResponse.Put(v)
}
