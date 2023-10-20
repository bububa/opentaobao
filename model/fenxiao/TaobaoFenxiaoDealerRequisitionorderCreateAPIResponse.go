package fenxiao

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFenxiaoDealerRequisitionorderCreateAPIResponse 创建经销采购申请 API返回值
// taobao.fenxiao.dealer.requisitionorder.create
//
// 创建经销采购申请
type TaobaoFenxiaoDealerRequisitionorderCreateAPIResponse struct {
	model.CommonResponse
	TaobaoFenxiaoDealerRequisitionorderCreateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoFenxiaoDealerRequisitionorderCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoFenxiaoDealerRequisitionorderCreateAPIResponseModel).Reset()
}

// TaobaoFenxiaoDealerRequisitionorderCreateAPIResponseModel is 创建经销采购申请 成功返回结果
type TaobaoFenxiaoDealerRequisitionorderCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"fenxiao_dealer_requisitionorder_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 经销采购申请编号
	DealerOrderId int64 `json:"dealer_order_id,omitempty" xml:"dealer_order_id,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoFenxiaoDealerRequisitionorderCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.DealerOrderId = 0
}

var poolTaobaoFenxiaoDealerRequisitionorderCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoFenxiaoDealerRequisitionorderCreateAPIResponse)
	},
}

// GetTaobaoFenxiaoDealerRequisitionorderCreateAPIResponse 从 sync.Pool 获取 TaobaoFenxiaoDealerRequisitionorderCreateAPIResponse
func GetTaobaoFenxiaoDealerRequisitionorderCreateAPIResponse() *TaobaoFenxiaoDealerRequisitionorderCreateAPIResponse {
	return poolTaobaoFenxiaoDealerRequisitionorderCreateAPIResponse.Get().(*TaobaoFenxiaoDealerRequisitionorderCreateAPIResponse)
}

// ReleaseTaobaoFenxiaoDealerRequisitionorderCreateAPIResponse 将 TaobaoFenxiaoDealerRequisitionorderCreateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoFenxiaoDealerRequisitionorderCreateAPIResponse(v *TaobaoFenxiaoDealerRequisitionorderCreateAPIResponse) {
	v.Reset()
	poolTaobaoFenxiaoDealerRequisitionorderCreateAPIResponse.Put(v)
}
