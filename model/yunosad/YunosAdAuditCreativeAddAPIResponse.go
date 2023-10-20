package yunosad

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosAdAuditCreativeAddAPIResponse 单个创意预审接口 API返回值
// yunos.ad.audit.creative.add
//
// YunOS广告业务第三方DSP单个创意预审接口
type YunosAdAuditCreativeAddAPIResponse struct {
	model.CommonResponse
	YunosAdAuditCreativeAddAPIResponseModel
}

// Reset 清空结构体
func (m *YunosAdAuditCreativeAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YunosAdAuditCreativeAddAPIResponseModel).Reset()
}

// YunosAdAuditCreativeAddAPIResponseModel is 单个创意预审接口 成功返回结果
type YunosAdAuditCreativeAddAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_ad_audit_creative_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// statusCode
	StatusCode int64 `json:"status_code,omitempty" xml:"status_code,omitempty"`
	// isOk
	IsOk bool `json:"is_ok,omitempty" xml:"is_ok,omitempty"`
}

// Reset 清空结构体
func (m *YunosAdAuditCreativeAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Message = ""
	m.StatusCode = 0
	m.IsOk = false
}

var poolYunosAdAuditCreativeAddAPIResponse = sync.Pool{
	New: func() any {
		return new(YunosAdAuditCreativeAddAPIResponse)
	},
}

// GetYunosAdAuditCreativeAddAPIResponse 从 sync.Pool 获取 YunosAdAuditCreativeAddAPIResponse
func GetYunosAdAuditCreativeAddAPIResponse() *YunosAdAuditCreativeAddAPIResponse {
	return poolYunosAdAuditCreativeAddAPIResponse.Get().(*YunosAdAuditCreativeAddAPIResponse)
}

// ReleaseYunosAdAuditCreativeAddAPIResponse 将 YunosAdAuditCreativeAddAPIResponse 保存到 sync.Pool
func ReleaseYunosAdAuditCreativeAddAPIResponse(v *YunosAdAuditCreativeAddAPIResponse) {
	v.Reset()
	poolYunosAdAuditCreativeAddAPIResponse.Put(v)
}
