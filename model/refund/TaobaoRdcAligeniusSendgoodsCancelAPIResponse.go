package refund

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoRdcAligeniusSendgoodsCancelAPIResponse 取消发货 API返回值
// taobao.rdc.aligenius.sendgoods.cancel
//
// 提供商家在仅退款中发送取消发货状态
type TaobaoRdcAligeniusSendgoodsCancelAPIResponse struct {
	model.CommonResponse
	TaobaoRdcAligeniusSendgoodsCancelAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoRdcAligeniusSendgoodsCancelAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoRdcAligeniusSendgoodsCancelAPIResponseModel).Reset()
}

// TaobaoRdcAligeniusSendgoodsCancelAPIResponseModel is 取消发货 成功返回结果
type TaobaoRdcAligeniusSendgoodsCancelAPIResponseModel struct {
	XMLName xml.Name `xml:"rdc_aligenius_sendgoods_cancel_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaoRdcAligeniusSendgoodsCancelResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoRdcAligeniusSendgoodsCancelAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoRdcAligeniusSendgoodsCancelAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoRdcAligeniusSendgoodsCancelAPIResponse)
	},
}

// GetTaobaoRdcAligeniusSendgoodsCancelAPIResponse 从 sync.Pool 获取 TaobaoRdcAligeniusSendgoodsCancelAPIResponse
func GetTaobaoRdcAligeniusSendgoodsCancelAPIResponse() *TaobaoRdcAligeniusSendgoodsCancelAPIResponse {
	return poolTaobaoRdcAligeniusSendgoodsCancelAPIResponse.Get().(*TaobaoRdcAligeniusSendgoodsCancelAPIResponse)
}

// ReleaseTaobaoRdcAligeniusSendgoodsCancelAPIResponse 将 TaobaoRdcAligeniusSendgoodsCancelAPIResponse 保存到 sync.Pool
func ReleaseTaobaoRdcAligeniusSendgoodsCancelAPIResponse(v *TaobaoRdcAligeniusSendgoodsCancelAPIResponse) {
	v.Reset()
	poolTaobaoRdcAligeniusSendgoodsCancelAPIResponse.Put(v)
}
