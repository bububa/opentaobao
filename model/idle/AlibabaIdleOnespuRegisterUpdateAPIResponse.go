package idle

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleOnespuRegisterUpdateAPIResponse 闲鱼 ONESPU 挂载接口 API返回值
// alibaba.idle.onespu.register.update
//
// 闲鱼 ONESPU 挂载接口
type AlibabaIdleOnespuRegisterUpdateAPIResponse struct {
	model.CommonResponse
	AlibabaIdleOnespuRegisterUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIdleOnespuRegisterUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIdleOnespuRegisterUpdateAPIResponseModel).Reset()
}

// AlibabaIdleOnespuRegisterUpdateAPIResponseModel is 闲鱼 ONESPU 挂载接口 成功返回结果
type AlibabaIdleOnespuRegisterUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_onespu_register_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *AlibabaIdleOnespuRegisterUpdateResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIdleOnespuRegisterUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaIdleOnespuRegisterUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIdleOnespuRegisterUpdateAPIResponse)
	},
}

// GetAlibabaIdleOnespuRegisterUpdateAPIResponse 从 sync.Pool 获取 AlibabaIdleOnespuRegisterUpdateAPIResponse
func GetAlibabaIdleOnespuRegisterUpdateAPIResponse() *AlibabaIdleOnespuRegisterUpdateAPIResponse {
	return poolAlibabaIdleOnespuRegisterUpdateAPIResponse.Get().(*AlibabaIdleOnespuRegisterUpdateAPIResponse)
}

// ReleaseAlibabaIdleOnespuRegisterUpdateAPIResponse 将 AlibabaIdleOnespuRegisterUpdateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIdleOnespuRegisterUpdateAPIResponse(v *AlibabaIdleOnespuRegisterUpdateAPIResponse) {
	v.Reset()
	poolAlibabaIdleOnespuRegisterUpdateAPIResponse.Put(v)
}
