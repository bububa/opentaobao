package fenxiao

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFenxiaoDealerRequisitionorderCloseAPIResponse 供应商/分销商关闭采购申请/经销采购单 API返回值
// taobao.fenxiao.dealer.requisitionorder.close
//
// 供应商或分销商关闭采购申请/经销采购单
type TaobaoFenxiaoDealerRequisitionorderCloseAPIResponse struct {
	model.CommonResponse
	TaobaoFenxiaoDealerRequisitionorderCloseAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoFenxiaoDealerRequisitionorderCloseAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoFenxiaoDealerRequisitionorderCloseAPIResponseModel).Reset()
}

// TaobaoFenxiaoDealerRequisitionorderCloseAPIResponseModel is 供应商/分销商关闭采购申请/经销采购单 成功返回结果
type TaobaoFenxiaoDealerRequisitionorderCloseAPIResponseModel struct {
	XMLName xml.Name `xml:"fenxiao_dealer_requisitionorder_close_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 操作是否成功。true：成功；false：失败。
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoFenxiaoDealerRequisitionorderCloseAPIResponseModel) Reset() {
	m.RequestId = ""
	m.IsSuccess = false
}

var poolTaobaoFenxiaoDealerRequisitionorderCloseAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoFenxiaoDealerRequisitionorderCloseAPIResponse)
	},
}

// GetTaobaoFenxiaoDealerRequisitionorderCloseAPIResponse 从 sync.Pool 获取 TaobaoFenxiaoDealerRequisitionorderCloseAPIResponse
func GetTaobaoFenxiaoDealerRequisitionorderCloseAPIResponse() *TaobaoFenxiaoDealerRequisitionorderCloseAPIResponse {
	return poolTaobaoFenxiaoDealerRequisitionorderCloseAPIResponse.Get().(*TaobaoFenxiaoDealerRequisitionorderCloseAPIResponse)
}

// ReleaseTaobaoFenxiaoDealerRequisitionorderCloseAPIResponse 将 TaobaoFenxiaoDealerRequisitionorderCloseAPIResponse 保存到 sync.Pool
func ReleaseTaobaoFenxiaoDealerRequisitionorderCloseAPIResponse(v *TaobaoFenxiaoDealerRequisitionorderCloseAPIResponse) {
	v.Reset()
	poolTaobaoFenxiaoDealerRequisitionorderCloseAPIResponse.Put(v)
}
