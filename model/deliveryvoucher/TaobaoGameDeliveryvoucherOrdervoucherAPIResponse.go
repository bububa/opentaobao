package deliveryvoucher

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoGameDeliveryvoucherOrdervoucherAPIResponse 预约接口 API返回值
// taobao.game.deliveryvoucher.ordervoucher
//
// 提货券发券接口：同步券和订单的关联信息
type TaobaoGameDeliveryvoucherOrdervoucherAPIResponse struct {
	model.CommonResponse
	TaobaoGameDeliveryvoucherOrdervoucherAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoGameDeliveryvoucherOrdervoucherAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoGameDeliveryvoucherOrdervoucherAPIResponseModel).Reset()
}

// TaobaoGameDeliveryvoucherOrdervoucherAPIResponseModel is 预约接口 成功返回结果
type TaobaoGameDeliveryvoucherOrdervoucherAPIResponseModel struct {
	XMLName xml.Name `xml:"game_deliveryvoucher_ordervoucher_response"`
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
func (m *TaobaoGameDeliveryvoucherOrdervoucherAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Message = ""
	m.ResultCode = ""
	m.IsSuccess = false
}

var poolTaobaoGameDeliveryvoucherOrdervoucherAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoGameDeliveryvoucherOrdervoucherAPIResponse)
	},
}

// GetTaobaoGameDeliveryvoucherOrdervoucherAPIResponse 从 sync.Pool 获取 TaobaoGameDeliveryvoucherOrdervoucherAPIResponse
func GetTaobaoGameDeliveryvoucherOrdervoucherAPIResponse() *TaobaoGameDeliveryvoucherOrdervoucherAPIResponse {
	return poolTaobaoGameDeliveryvoucherOrdervoucherAPIResponse.Get().(*TaobaoGameDeliveryvoucherOrdervoucherAPIResponse)
}

// ReleaseTaobaoGameDeliveryvoucherOrdervoucherAPIResponse 将 TaobaoGameDeliveryvoucherOrdervoucherAPIResponse 保存到 sync.Pool
func ReleaseTaobaoGameDeliveryvoucherOrdervoucherAPIResponse(v *TaobaoGameDeliveryvoucherOrdervoucherAPIResponse) {
	v.Reset()
	poolTaobaoGameDeliveryvoucherOrdervoucherAPIResponse.Put(v)
}
