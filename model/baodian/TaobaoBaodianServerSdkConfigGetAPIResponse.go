package baodian

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBaodianServerSdkConfigGetAPIResponse 获取宝点SDK的配置项（已迁移） API返回值
// taobao.baodian.server.sdk.config.get
//
// 获取SDK各种配置项（已迁移）
type TaobaoBaodianServerSdkConfigGetAPIResponse struct {
	model.CommonResponse
	TaobaoBaodianServerSdkConfigGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoBaodianServerSdkConfigGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoBaodianServerSdkConfigGetAPIResponseModel).Reset()
}

// TaobaoBaodianServerSdkConfigGetAPIResponseModel is 获取宝点SDK的配置项（已迁移） 成功返回结果
type TaobaoBaodianServerSdkConfigGetAPIResponseModel struct {
	XMLName xml.Name `xml:"baodian_server_sdk_config_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回sdk配置的字符串
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoBaodianServerSdkConfigGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = ""
}

var poolTaobaoBaodianServerSdkConfigGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoBaodianServerSdkConfigGetAPIResponse)
	},
}

// GetTaobaoBaodianServerSdkConfigGetAPIResponse 从 sync.Pool 获取 TaobaoBaodianServerSdkConfigGetAPIResponse
func GetTaobaoBaodianServerSdkConfigGetAPIResponse() *TaobaoBaodianServerSdkConfigGetAPIResponse {
	return poolTaobaoBaodianServerSdkConfigGetAPIResponse.Get().(*TaobaoBaodianServerSdkConfigGetAPIResponse)
}

// ReleaseTaobaoBaodianServerSdkConfigGetAPIResponse 将 TaobaoBaodianServerSdkConfigGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoBaodianServerSdkConfigGetAPIResponse(v *TaobaoBaodianServerSdkConfigGetAPIResponse) {
	v.Reset()
	poolTaobaoBaodianServerSdkConfigGetAPIResponse.Put(v)
}
