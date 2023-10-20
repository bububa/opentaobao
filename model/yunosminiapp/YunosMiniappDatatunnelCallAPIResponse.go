package yunosminiapp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosMiniappDatatunnelCallAPIResponse 车载小程序外部服务调用 API返回值
// yunos.miniapp.datatunnel.call
//
// 对客户提供的api进行统一封装调用。
type YunosMiniappDatatunnelCallAPIResponse struct {
	model.CommonResponse
	YunosMiniappDatatunnelCallAPIResponseModel
}

// Reset 清空结构体
func (m *YunosMiniappDatatunnelCallAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YunosMiniappDatatunnelCallAPIResponseModel).Reset()
}

// YunosMiniappDatatunnelCallAPIResponseModel is 车载小程序外部服务调用 成功返回结果
type YunosMiniappDatatunnelCallAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_miniapp_datatunnel_call_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *YunosMiniappDatatunnelCallMapResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *YunosMiniappDatatunnelCallAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolYunosMiniappDatatunnelCallAPIResponse = sync.Pool{
	New: func() any {
		return new(YunosMiniappDatatunnelCallAPIResponse)
	},
}

// GetYunosMiniappDatatunnelCallAPIResponse 从 sync.Pool 获取 YunosMiniappDatatunnelCallAPIResponse
func GetYunosMiniappDatatunnelCallAPIResponse() *YunosMiniappDatatunnelCallAPIResponse {
	return poolYunosMiniappDatatunnelCallAPIResponse.Get().(*YunosMiniappDatatunnelCallAPIResponse)
}

// ReleaseYunosMiniappDatatunnelCallAPIResponse 将 YunosMiniappDatatunnelCallAPIResponse 保存到 sync.Pool
func ReleaseYunosMiniappDatatunnelCallAPIResponse(v *YunosMiniappDatatunnelCallAPIResponse) {
	v.Reset()
	poolYunosMiniappDatatunnelCallAPIResponse.Put(v)
}
