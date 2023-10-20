package logistic

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoRdcAligeniusWarehouseResendLogisticsMsgPostAPIResponse 补发单物流信息回传 API返回值
// taobao.rdc.aligenius.warehouse.resend.logistics.msg.post
//
// 补发单erp物流信息回传平台
type TaobaoRdcAligeniusWarehouseResendLogisticsMsgPostAPIResponse struct {
	model.CommonResponse
	TaobaoRdcAligeniusWarehouseResendLogisticsMsgPostAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoRdcAligeniusWarehouseResendLogisticsMsgPostAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoRdcAligeniusWarehouseResendLogisticsMsgPostAPIResponseModel).Reset()
}

// TaobaoRdcAligeniusWarehouseResendLogisticsMsgPostAPIResponseModel is 补发单物流信息回传 成功返回结果
type TaobaoRdcAligeniusWarehouseResendLogisticsMsgPostAPIResponseModel struct {
	XMLName xml.Name `xml:"rdc_aligenius_warehouse_resend_logistics_msg_post_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaoRdcAligeniusWarehouseResendLogisticsMsgPostResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoRdcAligeniusWarehouseResendLogisticsMsgPostAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoRdcAligeniusWarehouseResendLogisticsMsgPostAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoRdcAligeniusWarehouseResendLogisticsMsgPostAPIResponse)
	},
}

// GetTaobaoRdcAligeniusWarehouseResendLogisticsMsgPostAPIResponse 从 sync.Pool 获取 TaobaoRdcAligeniusWarehouseResendLogisticsMsgPostAPIResponse
func GetTaobaoRdcAligeniusWarehouseResendLogisticsMsgPostAPIResponse() *TaobaoRdcAligeniusWarehouseResendLogisticsMsgPostAPIResponse {
	return poolTaobaoRdcAligeniusWarehouseResendLogisticsMsgPostAPIResponse.Get().(*TaobaoRdcAligeniusWarehouseResendLogisticsMsgPostAPIResponse)
}

// ReleaseTaobaoRdcAligeniusWarehouseResendLogisticsMsgPostAPIResponse 将 TaobaoRdcAligeniusWarehouseResendLogisticsMsgPostAPIResponse 保存到 sync.Pool
func ReleaseTaobaoRdcAligeniusWarehouseResendLogisticsMsgPostAPIResponse(v *TaobaoRdcAligeniusWarehouseResendLogisticsMsgPostAPIResponse) {
	v.Reset()
	poolTaobaoRdcAligeniusWarehouseResendLogisticsMsgPostAPIResponse.Put(v)
}
