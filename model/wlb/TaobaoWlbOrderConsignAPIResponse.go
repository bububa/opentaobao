package wlb

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbOrderConsignAPIResponse 物流宝订单已发货通知接口 API返回值
// taobao.wlb.order.consign
//
// 如果erp导入淘宝交易订单到物流宝，当物流宝订单已发货的时候，erp需要调用该接口来通知物流订单和淘宝交易订单已发货
type TaobaoWlbOrderConsignAPIResponse struct {
	model.CommonResponse
	TaobaoWlbOrderConsignAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoWlbOrderConsignAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoWlbOrderConsignAPIResponseModel).Reset()
}

// TaobaoWlbOrderConsignAPIResponseModel is 物流宝订单已发货通知接口 成功返回结果
type TaobaoWlbOrderConsignAPIResponseModel struct {
	XMLName xml.Name `xml:"wlb_order_consign_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 修改时间
	ModifyTime string `json:"modify_time,omitempty" xml:"modify_time,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoWlbOrderConsignAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ModifyTime = ""
}

var poolTaobaoWlbOrderConsignAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoWlbOrderConsignAPIResponse)
	},
}

// GetTaobaoWlbOrderConsignAPIResponse 从 sync.Pool 获取 TaobaoWlbOrderConsignAPIResponse
func GetTaobaoWlbOrderConsignAPIResponse() *TaobaoWlbOrderConsignAPIResponse {
	return poolTaobaoWlbOrderConsignAPIResponse.Get().(*TaobaoWlbOrderConsignAPIResponse)
}

// ReleaseTaobaoWlbOrderConsignAPIResponse 将 TaobaoWlbOrderConsignAPIResponse 保存到 sync.Pool
func ReleaseTaobaoWlbOrderConsignAPIResponse(v *TaobaoWlbOrderConsignAPIResponse) {
	v.Reset()
	poolTaobaoWlbOrderConsignAPIResponse.Put(v)
}
