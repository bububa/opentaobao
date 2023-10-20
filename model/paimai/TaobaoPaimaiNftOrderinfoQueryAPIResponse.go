package paimai

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPaimaiNftOrderinfoQueryAPIResponse 查询订单类型 API返回值
// taobao.paimai.nft.orderinfo.query
//
// 查询订单类型
type TaobaoPaimaiNftOrderinfoQueryAPIResponse struct {
	model.CommonResponse
	TaobaoPaimaiNftOrderinfoQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoPaimaiNftOrderinfoQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoPaimaiNftOrderinfoQueryAPIResponseModel).Reset()
}

// TaobaoPaimaiNftOrderinfoQueryAPIResponseModel is 查询订单类型 成功返回结果
type TaobaoPaimaiNftOrderinfoQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"paimai_nft_orderinfo_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 订单信息
	OrderInfoList []NftTradeOrderDto `json:"order_info_list,omitempty" xml:"order_info_list>nft_trade_order_dto,omitempty"`
	// 错误码
	FailCode string `json:"fail_code,omitempty" xml:"fail_code,omitempty"`
	// 错误信息
	FailMessage string `json:"fail_message,omitempty" xml:"fail_message,omitempty"`
	// 是否成功
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoPaimaiNftOrderinfoQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.OrderInfoList = m.OrderInfoList[:0]
	m.FailCode = ""
	m.FailMessage = ""
	m.Result = false
}

var poolTaobaoPaimaiNftOrderinfoQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoPaimaiNftOrderinfoQueryAPIResponse)
	},
}

// GetTaobaoPaimaiNftOrderinfoQueryAPIResponse 从 sync.Pool 获取 TaobaoPaimaiNftOrderinfoQueryAPIResponse
func GetTaobaoPaimaiNftOrderinfoQueryAPIResponse() *TaobaoPaimaiNftOrderinfoQueryAPIResponse {
	return poolTaobaoPaimaiNftOrderinfoQueryAPIResponse.Get().(*TaobaoPaimaiNftOrderinfoQueryAPIResponse)
}

// ReleaseTaobaoPaimaiNftOrderinfoQueryAPIResponse 将 TaobaoPaimaiNftOrderinfoQueryAPIResponse 保存到 sync.Pool
func ReleaseTaobaoPaimaiNftOrderinfoQueryAPIResponse(v *TaobaoPaimaiNftOrderinfoQueryAPIResponse) {
	v.Reset()
	poolTaobaoPaimaiNftOrderinfoQueryAPIResponse.Put(v)
}
