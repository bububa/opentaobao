package campus

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusGuardantGatewayCallbackAPIResponse 人卡关系回调 API返回值
// alibaba.campus.guardant.gateway.callback
//
// 门禁供应商回调平台通知同步结果
type AlibabaCampusGuardantGatewayCallbackAPIResponse struct {
	model.CommonResponse
	AlibabaCampusGuardantGatewayCallbackAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaCampusGuardantGatewayCallbackAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaCampusGuardantGatewayCallbackAPIResponseModel).Reset()
}

// AlibabaCampusGuardantGatewayCallbackAPIResponseModel is 人卡关系回调 成功返回结果
type AlibabaCampusGuardantGatewayCallbackAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_guardant_gateway_callback_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// PojoResult
	Result *PojoResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaCampusGuardantGatewayCallbackAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaCampusGuardantGatewayCallbackAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaCampusGuardantGatewayCallbackAPIResponse)
	},
}

// GetAlibabaCampusGuardantGatewayCallbackAPIResponse 从 sync.Pool 获取 AlibabaCampusGuardantGatewayCallbackAPIResponse
func GetAlibabaCampusGuardantGatewayCallbackAPIResponse() *AlibabaCampusGuardantGatewayCallbackAPIResponse {
	return poolAlibabaCampusGuardantGatewayCallbackAPIResponse.Get().(*AlibabaCampusGuardantGatewayCallbackAPIResponse)
}

// ReleaseAlibabaCampusGuardantGatewayCallbackAPIResponse 将 AlibabaCampusGuardantGatewayCallbackAPIResponse 保存到 sync.Pool
func ReleaseAlibabaCampusGuardantGatewayCallbackAPIResponse(v *AlibabaCampusGuardantGatewayCallbackAPIResponse) {
	v.Reset()
	poolAlibabaCampusGuardantGatewayCallbackAPIResponse.Put(v)
}
