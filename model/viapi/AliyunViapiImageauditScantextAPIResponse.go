package viapi

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliyunViapiImageauditScantextAPIResponse 文字内容审核 API返回值
// aliyun.viapi.imageaudit.scantext
//
// 结合行为、内容，采用多维度、多模型、多检测手段，识别文本中的垃圾内容，规避色情、广告、灌水、渉政、辱骂等内容风险。
// 注意：如果返回结果里面的results为空，也代表指定类型检测通过。
type AliyunViapiImageauditScantextAPIResponse struct {
	model.CommonResponse
	AliyunViapiImageauditScantextAPIResponseModel
}

// Reset 清空结构体
func (m *AliyunViapiImageauditScantextAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AliyunViapiImageauditScantextAPIResponseModel).Reset()
}

// AliyunViapiImageauditScantextAPIResponseModel is 文字内容审核 成功返回结果
type AliyunViapiImageauditScantextAPIResponseModel struct {
	XMLName xml.Name `xml:"aliyun_viapi_imageaudit_scantext_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 请求ID
	TaobaoRequestId string `json:"taobao_request_id,omitempty" xml:"taobao_request_id,omitempty"`
	// 系统自动生成
	Data *AliyunViapiImageauditScantextData `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *AliyunViapiImageauditScantextAPIResponseModel) Reset() {
	m.RequestId = ""
	m.TaobaoRequestId = ""
	m.Data = nil
}

var poolAliyunViapiImageauditScantextAPIResponse = sync.Pool{
	New: func() any {
		return new(AliyunViapiImageauditScantextAPIResponse)
	},
}

// GetAliyunViapiImageauditScantextAPIResponse 从 sync.Pool 获取 AliyunViapiImageauditScantextAPIResponse
func GetAliyunViapiImageauditScantextAPIResponse() *AliyunViapiImageauditScantextAPIResponse {
	return poolAliyunViapiImageauditScantextAPIResponse.Get().(*AliyunViapiImageauditScantextAPIResponse)
}

// ReleaseAliyunViapiImageauditScantextAPIResponse 将 AliyunViapiImageauditScantextAPIResponse 保存到 sync.Pool
func ReleaseAliyunViapiImageauditScantextAPIResponse(v *AliyunViapiImageauditScantextAPIResponse) {
	v.Reset()
	poolAliyunViapiImageauditScantextAPIResponse.Put(v)
}
