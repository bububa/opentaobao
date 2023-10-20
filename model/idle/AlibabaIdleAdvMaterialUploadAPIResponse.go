package idle

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleAdvMaterialUploadAPIResponse 闲鱼用户增长素材中心素材上传接口 API返回值
// alibaba.idle.adv.material.upload
//
// 闲鱼用户增长素材中心素材上传接口
type AlibabaIdleAdvMaterialUploadAPIResponse struct {
	model.CommonResponse
	AlibabaIdleAdvMaterialUploadAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIdleAdvMaterialUploadAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIdleAdvMaterialUploadAPIResponseModel).Reset()
}

// AlibabaIdleAdvMaterialUploadAPIResponseModel is 闲鱼用户增长素材中心素材上传接口 成功返回结果
type AlibabaIdleAdvMaterialUploadAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_adv_material_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 1
	Result *IdleAdvResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIdleAdvMaterialUploadAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaIdleAdvMaterialUploadAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIdleAdvMaterialUploadAPIResponse)
	},
}

// GetAlibabaIdleAdvMaterialUploadAPIResponse 从 sync.Pool 获取 AlibabaIdleAdvMaterialUploadAPIResponse
func GetAlibabaIdleAdvMaterialUploadAPIResponse() *AlibabaIdleAdvMaterialUploadAPIResponse {
	return poolAlibabaIdleAdvMaterialUploadAPIResponse.Get().(*AlibabaIdleAdvMaterialUploadAPIResponse)
}

// ReleaseAlibabaIdleAdvMaterialUploadAPIResponse 将 AlibabaIdleAdvMaterialUploadAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIdleAdvMaterialUploadAPIResponse(v *AlibabaIdleAdvMaterialUploadAPIResponse) {
	v.Reset()
	poolAlibabaIdleAdvMaterialUploadAPIResponse.Put(v)
}
