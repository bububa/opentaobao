package tmallgenie

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAilabsAligenieOpencontentPushAPIResponse 天猫精灵内容接入标准接口 API返回值
// alibaba.ailabs.aligenie.opencontent.push
//
// 第三方内容接入天猫精灵内容库，供相关技能使用
type AlibabaAilabsAligenieOpencontentPushAPIResponse struct {
	model.CommonResponse
	AlibabaAilabsAligenieOpencontentPushAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAilabsAligenieOpencontentPushAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAilabsAligenieOpencontentPushAPIResponseModel).Reset()
}

// AlibabaAilabsAligenieOpencontentPushAPIResponseModel is 天猫精灵内容接入标准接口 成功返回结果
type AlibabaAilabsAligenieOpencontentPushAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ailabs_aligenie_opencontent_push_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaAilabsAligenieOpencontentPushResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAilabsAligenieOpencontentPushAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAilabsAligenieOpencontentPushAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAilabsAligenieOpencontentPushAPIResponse)
	},
}

// GetAlibabaAilabsAligenieOpencontentPushAPIResponse 从 sync.Pool 获取 AlibabaAilabsAligenieOpencontentPushAPIResponse
func GetAlibabaAilabsAligenieOpencontentPushAPIResponse() *AlibabaAilabsAligenieOpencontentPushAPIResponse {
	return poolAlibabaAilabsAligenieOpencontentPushAPIResponse.Get().(*AlibabaAilabsAligenieOpencontentPushAPIResponse)
}

// ReleaseAlibabaAilabsAligenieOpencontentPushAPIResponse 将 AlibabaAilabsAligenieOpencontentPushAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAilabsAligenieOpencontentPushAPIResponse(v *AlibabaAilabsAligenieOpencontentPushAPIResponse) {
	v.Reset()
	poolAlibabaAilabsAligenieOpencontentPushAPIResponse.Put(v)
}
