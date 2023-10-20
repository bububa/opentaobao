package tbk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTbkDgVegasTljStopAPIResponse 淘宝客-推广者-淘礼金暂停发放 API返回值
// taobao.tbk.dg.vegas.tlj.stop
//
// 淘宝客推广者可对已经创建的淘礼金暂停发放
type TaobaoTbkDgVegasTljStopAPIResponse struct {
	model.CommonResponse
	TaobaoTbkDgVegasTljStopAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTbkDgVegasTljStopAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTbkDgVegasTljStopAPIResponseModel).Reset()
}

// TaobaoTbkDgVegasTljStopAPIResponseModel is 淘宝客-推广者-淘礼金暂停发放 成功返回结果
type TaobaoTbkDgVegasTljStopAPIResponseModel struct {
	XMLName xml.Name `xml:"tbk_dg_vegas_tlj_stop_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// model
	Model *UpdateStatusResult `json:"model,omitempty" xml:"model,omitempty"`
	// 调用接口是否成功
	ResultSuccess bool `json:"result_success,omitempty" xml:"result_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTbkDgVegasTljStopAPIResponseModel) Reset() {
	m.RequestId = ""
	m.MsgInfo = ""
	m.MsgCode = ""
	m.Model = nil
	m.ResultSuccess = false
}

var poolTaobaoTbkDgVegasTljStopAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTbkDgVegasTljStopAPIResponse)
	},
}

// GetTaobaoTbkDgVegasTljStopAPIResponse 从 sync.Pool 获取 TaobaoTbkDgVegasTljStopAPIResponse
func GetTaobaoTbkDgVegasTljStopAPIResponse() *TaobaoTbkDgVegasTljStopAPIResponse {
	return poolTaobaoTbkDgVegasTljStopAPIResponse.Get().(*TaobaoTbkDgVegasTljStopAPIResponse)
}

// ReleaseTaobaoTbkDgVegasTljStopAPIResponse 将 TaobaoTbkDgVegasTljStopAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTbkDgVegasTljStopAPIResponse(v *TaobaoTbkDgVegasTljStopAPIResponse) {
	v.Reset()
	poolTaobaoTbkDgVegasTljStopAPIResponse.Put(v)
}
