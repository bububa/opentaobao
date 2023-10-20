package wlb

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbOrderJzwithinsConsignAPIResponse 家装发货接口 API返回值
// taobao.wlb.order.jzwithins.consign
//
// 为支持家装类目的商家，对绑定家装物流服务的订单可以在商家的ERP中发货、批量发货，因此开发带安装服务商的发货接口
type TaobaoWlbOrderJzwithinsConsignAPIResponse struct {
	model.CommonResponse
	TaobaoWlbOrderJzwithinsConsignAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoWlbOrderJzwithinsConsignAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoWlbOrderJzwithinsConsignAPIResponseModel).Reset()
}

// TaobaoWlbOrderJzwithinsConsignAPIResponseModel is 家装发货接口 成功返回结果
type TaobaoWlbOrderJzwithinsConsignAPIResponseModel struct {
	XMLName xml.Name `xml:"wlb_order_jzwithins_consign_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 发货返回信息，如果发货错误则报出对应错误
	ResultInfo string `json:"result_info,omitempty" xml:"result_info,omitempty"`
	// 发货成功或者失败
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoWlbOrderJzwithinsConsignAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultInfo = ""
	m.IsSuccess = false
}

var poolTaobaoWlbOrderJzwithinsConsignAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoWlbOrderJzwithinsConsignAPIResponse)
	},
}

// GetTaobaoWlbOrderJzwithinsConsignAPIResponse 从 sync.Pool 获取 TaobaoWlbOrderJzwithinsConsignAPIResponse
func GetTaobaoWlbOrderJzwithinsConsignAPIResponse() *TaobaoWlbOrderJzwithinsConsignAPIResponse {
	return poolTaobaoWlbOrderJzwithinsConsignAPIResponse.Get().(*TaobaoWlbOrderJzwithinsConsignAPIResponse)
}

// ReleaseTaobaoWlbOrderJzwithinsConsignAPIResponse 将 TaobaoWlbOrderJzwithinsConsignAPIResponse 保存到 sync.Pool
func ReleaseTaobaoWlbOrderJzwithinsConsignAPIResponse(v *TaobaoWlbOrderJzwithinsConsignAPIResponse) {
	v.Reset()
	poolTaobaoWlbOrderJzwithinsConsignAPIResponse.Put(v)
}
