package logistic

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoRdcAligeniusWarehouseResendUpdateAPIResponse 补发单状态回传 API返回值
// taobao.rdc.aligenius.warehouse.resend.update
//
// 补发单状态回传接口
type TaobaoRdcAligeniusWarehouseResendUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoRdcAligeniusWarehouseResendUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoRdcAligeniusWarehouseResendUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoRdcAligeniusWarehouseResendUpdateAPIResponseModel).Reset()
}

// TaobaoRdcAligeniusWarehouseResendUpdateAPIResponseModel is 补发单状态回传 成功返回结果
type TaobaoRdcAligeniusWarehouseResendUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"rdc_aligenius_warehouse_resend_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaoRdcAligeniusWarehouseResendUpdateResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoRdcAligeniusWarehouseResendUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoRdcAligeniusWarehouseResendUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoRdcAligeniusWarehouseResendUpdateAPIResponse)
	},
}

// GetTaobaoRdcAligeniusWarehouseResendUpdateAPIResponse 从 sync.Pool 获取 TaobaoRdcAligeniusWarehouseResendUpdateAPIResponse
func GetTaobaoRdcAligeniusWarehouseResendUpdateAPIResponse() *TaobaoRdcAligeniusWarehouseResendUpdateAPIResponse {
	return poolTaobaoRdcAligeniusWarehouseResendUpdateAPIResponse.Get().(*TaobaoRdcAligeniusWarehouseResendUpdateAPIResponse)
}

// ReleaseTaobaoRdcAligeniusWarehouseResendUpdateAPIResponse 将 TaobaoRdcAligeniusWarehouseResendUpdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoRdcAligeniusWarehouseResendUpdateAPIResponse(v *TaobaoRdcAligeniusWarehouseResendUpdateAPIResponse) {
	v.Reset()
	poolTaobaoRdcAligeniusWarehouseResendUpdateAPIResponse.Put(v)
}
