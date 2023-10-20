package tmallsc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallSerivcecenterServicerorderInsuranceCallbackAPIResponse 服务商回传保单信息 API返回值
// tmall.serivcecenter.servicerorder.insurance.callback
//
// 服务商回传保单信息
type TmallSerivcecenterServicerorderInsuranceCallbackAPIResponse struct {
	model.CommonResponse
	TmallSerivcecenterServicerorderInsuranceCallbackAPIResponseModel
}

// Reset 清空结构体
func (m *TmallSerivcecenterServicerorderInsuranceCallbackAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallSerivcecenterServicerorderInsuranceCallbackAPIResponseModel).Reset()
}

// TmallSerivcecenterServicerorderInsuranceCallbackAPIResponseModel is 服务商回传保单信息 成功返回结果
type TmallSerivcecenterServicerorderInsuranceCallbackAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_serivcecenter_servicerorder_insurance_callback_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 错误信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TmallSerivcecenterServicerorderInsuranceCallbackAPIResponseModel) Reset() {
	m.RequestId = ""
	m.MsgCode = ""
	m.MsgInfo = ""
	m.IsSuccess = false
}

var poolTmallSerivcecenterServicerorderInsuranceCallbackAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallSerivcecenterServicerorderInsuranceCallbackAPIResponse)
	},
}

// GetTmallSerivcecenterServicerorderInsuranceCallbackAPIResponse 从 sync.Pool 获取 TmallSerivcecenterServicerorderInsuranceCallbackAPIResponse
func GetTmallSerivcecenterServicerorderInsuranceCallbackAPIResponse() *TmallSerivcecenterServicerorderInsuranceCallbackAPIResponse {
	return poolTmallSerivcecenterServicerorderInsuranceCallbackAPIResponse.Get().(*TmallSerivcecenterServicerorderInsuranceCallbackAPIResponse)
}

// ReleaseTmallSerivcecenterServicerorderInsuranceCallbackAPIResponse 将 TmallSerivcecenterServicerorderInsuranceCallbackAPIResponse 保存到 sync.Pool
func ReleaseTmallSerivcecenterServicerorderInsuranceCallbackAPIResponse(v *TmallSerivcecenterServicerorderInsuranceCallbackAPIResponse) {
	v.Reset()
	poolTmallSerivcecenterServicerorderInsuranceCallbackAPIResponse.Put(v)
}
