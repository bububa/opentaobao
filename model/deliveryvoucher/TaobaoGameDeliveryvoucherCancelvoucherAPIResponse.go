package deliveryvoucher

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoGameDeliveryvoucherCancelvoucherAPIResponse 作废券 API返回值
// taobao.game.deliveryvoucher.cancelvoucher
//
// 提货券发券接口：同步券和订单的关联信息
type TaobaoGameDeliveryvoucherCancelvoucherAPIResponse struct {
	model.CommonResponse
	TaobaoGameDeliveryvoucherCancelvoucherAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoGameDeliveryvoucherCancelvoucherAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoGameDeliveryvoucherCancelvoucherAPIResponseModel).Reset()
}

// TaobaoGameDeliveryvoucherCancelvoucherAPIResponseModel is 作废券 成功返回结果
type TaobaoGameDeliveryvoucherCancelvoucherAPIResponseModel struct {
	XMLName xml.Name `xml:"game_deliveryvoucher_cancelvoucher_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// code
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// success
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoGameDeliveryvoucherCancelvoucherAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Message = ""
	m.ResultCode = ""
	m.IsSuccess = false
}

var poolTaobaoGameDeliveryvoucherCancelvoucherAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoGameDeliveryvoucherCancelvoucherAPIResponse)
	},
}

// GetTaobaoGameDeliveryvoucherCancelvoucherAPIResponse 从 sync.Pool 获取 TaobaoGameDeliveryvoucherCancelvoucherAPIResponse
func GetTaobaoGameDeliveryvoucherCancelvoucherAPIResponse() *TaobaoGameDeliveryvoucherCancelvoucherAPIResponse {
	return poolTaobaoGameDeliveryvoucherCancelvoucherAPIResponse.Get().(*TaobaoGameDeliveryvoucherCancelvoucherAPIResponse)
}

// ReleaseTaobaoGameDeliveryvoucherCancelvoucherAPIResponse 将 TaobaoGameDeliveryvoucherCancelvoucherAPIResponse 保存到 sync.Pool
func ReleaseTaobaoGameDeliveryvoucherCancelvoucherAPIResponse(v *TaobaoGameDeliveryvoucherCancelvoucherAPIResponse) {
	v.Reset()
	poolTaobaoGameDeliveryvoucherCancelvoucherAPIResponse.Put(v)
}
