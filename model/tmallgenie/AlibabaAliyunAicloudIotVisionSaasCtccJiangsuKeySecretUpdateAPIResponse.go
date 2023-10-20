package tmallgenie

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAliyunAicloudIotVisionSaasCtccJiangsuKeySecretUpdateAPIResponse 天猫精灵 IoT 视频 SaaS 服务-江苏电信-appKeySecret 更新 API返回值
// alibaba.aliyun.aicloud.iot.vision.saas.ctcc.jiangsu.key.secret.update
//
// 天猫精灵 IoT 视频 SaaS 服务-江苏电信-appKeySecret 更新
type AlibabaAliyunAicloudIotVisionSaasCtccJiangsuKeySecretUpdateAPIResponse struct {
	model.CommonResponse
	AlibabaAliyunAicloudIotVisionSaasCtccJiangsuKeySecretUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAliyunAicloudIotVisionSaasCtccJiangsuKeySecretUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAliyunAicloudIotVisionSaasCtccJiangsuKeySecretUpdateAPIResponseModel).Reset()
}

// AlibabaAliyunAicloudIotVisionSaasCtccJiangsuKeySecretUpdateAPIResponseModel is 天猫精灵 IoT 视频 SaaS 服务-江苏电信-appKeySecret 更新 成功返回结果
type AlibabaAliyunAicloudIotVisionSaasCtccJiangsuKeySecretUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_aliyun_aicloud_iot_vision_saas_ctcc_jiangsu_key_secret_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息。success 表示成功
	RspMsg string `json:"rsp_msg,omitempty" xml:"rsp_msg,omitempty"`
	// 一次请求的唯一标识符，和请求中的 seq_id 对齐
	SeqId string `json:"seq_id,omitempty" xml:"seq_id,omitempty"`
	// 错误码。200 表示成功
	RspCode int64 `json:"rsp_code,omitempty" xml:"rsp_code,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAliyunAicloudIotVisionSaasCtccJiangsuKeySecretUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.RspMsg = ""
	m.SeqId = ""
	m.RspCode = 0
}

var poolAlibabaAliyunAicloudIotVisionSaasCtccJiangsuKeySecretUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAliyunAicloudIotVisionSaasCtccJiangsuKeySecretUpdateAPIResponse)
	},
}

// GetAlibabaAliyunAicloudIotVisionSaasCtccJiangsuKeySecretUpdateAPIResponse 从 sync.Pool 获取 AlibabaAliyunAicloudIotVisionSaasCtccJiangsuKeySecretUpdateAPIResponse
func GetAlibabaAliyunAicloudIotVisionSaasCtccJiangsuKeySecretUpdateAPIResponse() *AlibabaAliyunAicloudIotVisionSaasCtccJiangsuKeySecretUpdateAPIResponse {
	return poolAlibabaAliyunAicloudIotVisionSaasCtccJiangsuKeySecretUpdateAPIResponse.Get().(*AlibabaAliyunAicloudIotVisionSaasCtccJiangsuKeySecretUpdateAPIResponse)
}

// ReleaseAlibabaAliyunAicloudIotVisionSaasCtccJiangsuKeySecretUpdateAPIResponse 将 AlibabaAliyunAicloudIotVisionSaasCtccJiangsuKeySecretUpdateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAliyunAicloudIotVisionSaasCtccJiangsuKeySecretUpdateAPIResponse(v *AlibabaAliyunAicloudIotVisionSaasCtccJiangsuKeySecretUpdateAPIResponse) {
	v.Reset()
	poolAlibabaAliyunAicloudIotVisionSaasCtccJiangsuKeySecretUpdateAPIResponse.Put(v)
}
