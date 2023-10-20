package fenxiao

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFenxiaoDealerRequisitionorderRemarkUpdateAPIResponse 修改经销采购单备注 API返回值
// taobao.fenxiao.dealer.requisitionorder.remark.update
//
// 供应商修改经销采购单备注
type TaobaoFenxiaoDealerRequisitionorderRemarkUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoFenxiaoDealerRequisitionorderRemarkUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoFenxiaoDealerRequisitionorderRemarkUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoFenxiaoDealerRequisitionorderRemarkUpdateAPIResponseModel).Reset()
}

// TaobaoFenxiaoDealerRequisitionorderRemarkUpdateAPIResponseModel is 修改经销采购单备注 成功返回结果
type TaobaoFenxiaoDealerRequisitionorderRemarkUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"fenxiao_dealer_requisitionorder_remark_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 操作是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoFenxiaoDealerRequisitionorderRemarkUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.IsSuccess = false
}

var poolTaobaoFenxiaoDealerRequisitionorderRemarkUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoFenxiaoDealerRequisitionorderRemarkUpdateAPIResponse)
	},
}

// GetTaobaoFenxiaoDealerRequisitionorderRemarkUpdateAPIResponse 从 sync.Pool 获取 TaobaoFenxiaoDealerRequisitionorderRemarkUpdateAPIResponse
func GetTaobaoFenxiaoDealerRequisitionorderRemarkUpdateAPIResponse() *TaobaoFenxiaoDealerRequisitionorderRemarkUpdateAPIResponse {
	return poolTaobaoFenxiaoDealerRequisitionorderRemarkUpdateAPIResponse.Get().(*TaobaoFenxiaoDealerRequisitionorderRemarkUpdateAPIResponse)
}

// ReleaseTaobaoFenxiaoDealerRequisitionorderRemarkUpdateAPIResponse 将 TaobaoFenxiaoDealerRequisitionorderRemarkUpdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoFenxiaoDealerRequisitionorderRemarkUpdateAPIResponse(v *TaobaoFenxiaoDealerRequisitionorderRemarkUpdateAPIResponse) {
	v.Reset()
	poolTaobaoFenxiaoDealerRequisitionorderRemarkUpdateAPIResponse.Put(v)
}
