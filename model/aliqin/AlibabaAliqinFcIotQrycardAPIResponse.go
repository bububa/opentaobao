package aliqin

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAliqinFcIotQrycardAPIResponse 查询终端信息 API返回值
// alibaba.aliqin.fc.iot.qrycard
//
// 查询终端信息
type AlibabaAliqinFcIotQrycardAPIResponse struct {
	model.CommonResponse
	AlibabaAliqinFcIotQrycardAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAliqinFcIotQrycardAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAliqinFcIotQrycardAPIResponseModel).Reset()
}

// AlibabaAliqinFcIotQrycardAPIResponseModel is 查询终端信息 成功返回结果
type AlibabaAliqinFcIotQrycardAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_aliqin_fc_iot_qrycard_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 系统自动生成
	Result *AlibabaAliqinFcIotQrycardResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAliqinFcIotQrycardAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAliqinFcIotQrycardAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAliqinFcIotQrycardAPIResponse)
	},
}

// GetAlibabaAliqinFcIotQrycardAPIResponse 从 sync.Pool 获取 AlibabaAliqinFcIotQrycardAPIResponse
func GetAlibabaAliqinFcIotQrycardAPIResponse() *AlibabaAliqinFcIotQrycardAPIResponse {
	return poolAlibabaAliqinFcIotQrycardAPIResponse.Get().(*AlibabaAliqinFcIotQrycardAPIResponse)
}

// ReleaseAlibabaAliqinFcIotQrycardAPIResponse 将 AlibabaAliqinFcIotQrycardAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAliqinFcIotQrycardAPIResponse(v *AlibabaAliqinFcIotQrycardAPIResponse) {
	v.Reset()
	poolAlibabaAliqinFcIotQrycardAPIResponse.Put(v)
}
