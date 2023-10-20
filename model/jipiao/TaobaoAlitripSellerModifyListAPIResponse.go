package jipiao

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripSellerModifyListAPIResponse 【机票代理商订单】改签订单列表 API返回值
// taobao.alitrip.seller.modify.list
//
// 提供机票代理商查询改签订单列表
type TaobaoAlitripSellerModifyListAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripSellerModifyListAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAlitripSellerModifyListAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAlitripSellerModifyListAPIResponseModel).Reset()
}

// TaobaoAlitripSellerModifyListAPIResponseModel is 【机票代理商订单】改签订单列表 成功返回结果
type TaobaoAlitripSellerModifyListAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_seller_modify_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 改签订单列表
	OrderList []SyncOrderDo `json:"order_list,omitempty" xml:"order_list>sync_order_do,omitempty"`
	// 查出总记录数
	TotalItem int64 `json:"total_item,omitempty" xml:"total_item,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAlitripSellerModifyListAPIResponseModel) Reset() {
	m.RequestId = ""
	m.OrderList = m.OrderList[:0]
	m.TotalItem = 0
}

var poolTaobaoAlitripSellerModifyListAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripSellerModifyListAPIResponse)
	},
}

// GetTaobaoAlitripSellerModifyListAPIResponse 从 sync.Pool 获取 TaobaoAlitripSellerModifyListAPIResponse
func GetTaobaoAlitripSellerModifyListAPIResponse() *TaobaoAlitripSellerModifyListAPIResponse {
	return poolTaobaoAlitripSellerModifyListAPIResponse.Get().(*TaobaoAlitripSellerModifyListAPIResponse)
}

// ReleaseTaobaoAlitripSellerModifyListAPIResponse 将 TaobaoAlitripSellerModifyListAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAlitripSellerModifyListAPIResponse(v *TaobaoAlitripSellerModifyListAPIResponse) {
	v.Reset()
	poolTaobaoAlitripSellerModifyListAPIResponse.Put(v)
}
