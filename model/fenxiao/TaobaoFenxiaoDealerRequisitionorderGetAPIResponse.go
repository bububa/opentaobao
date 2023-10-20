package fenxiao

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFenxiaoDealerRequisitionorderGetAPIResponse 批量查询采购申请/经销采购单 API返回值
// taobao.fenxiao.dealer.requisitionorder.get
//
// 批量查询采购申请/经销采购单，目前支持供应商和分销商查询
type TaobaoFenxiaoDealerRequisitionorderGetAPIResponse struct {
	model.CommonResponse
	TaobaoFenxiaoDealerRequisitionorderGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoFenxiaoDealerRequisitionorderGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoFenxiaoDealerRequisitionorderGetAPIResponseModel).Reset()
}

// TaobaoFenxiaoDealerRequisitionorderGetAPIResponseModel is 批量查询采购申请/经销采购单 成功返回结果
type TaobaoFenxiaoDealerRequisitionorderGetAPIResponseModel struct {
	XMLName xml.Name `xml:"fenxiao_dealer_requisitionorder_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 采购申请/经销采购单结果列表
	DealerOrders []DealerOrder `json:"dealer_orders,omitempty" xml:"dealer_orders>dealer_order,omitempty"`
	// 按查询条件查到的记录总数
	TotalResults int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoFenxiaoDealerRequisitionorderGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.DealerOrders = m.DealerOrders[:0]
	m.TotalResults = 0
}

var poolTaobaoFenxiaoDealerRequisitionorderGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoFenxiaoDealerRequisitionorderGetAPIResponse)
	},
}

// GetTaobaoFenxiaoDealerRequisitionorderGetAPIResponse 从 sync.Pool 获取 TaobaoFenxiaoDealerRequisitionorderGetAPIResponse
func GetTaobaoFenxiaoDealerRequisitionorderGetAPIResponse() *TaobaoFenxiaoDealerRequisitionorderGetAPIResponse {
	return poolTaobaoFenxiaoDealerRequisitionorderGetAPIResponse.Get().(*TaobaoFenxiaoDealerRequisitionorderGetAPIResponse)
}

// ReleaseTaobaoFenxiaoDealerRequisitionorderGetAPIResponse 将 TaobaoFenxiaoDealerRequisitionorderGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoFenxiaoDealerRequisitionorderGetAPIResponse(v *TaobaoFenxiaoDealerRequisitionorderGetAPIResponse) {
	v.Reset()
	poolTaobaoFenxiaoDealerRequisitionorderGetAPIResponse.Put(v)
}
