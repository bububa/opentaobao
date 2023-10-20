package viapi

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AliyunviapiimageauditscantextAPIResponse 文字内容审核 API返回值
// aliyun.viapi.imageaudit.scantext
//
// 结合行为、内容，采用多维度、多模型、多检测手段，识别文本中的垃圾内容，规避色情、广告、灌水、渉政、辱骂等内容风险。
// 注意：如果返回结果里面的results为空，也代表指定类型检测通过。
type AliyunviapiimageauditscantextAPIResponse struct {
	model.CommonResponse
	AliyunviapiimageauditscantextAPIResponseModel
}

// AliyunviapiimageauditscantextAPIResponseModel is 文字内容审核 成功返回结果
type AliyunviapiimageauditscantextAPIResponseModel struct {
	XMLName xml.Name `xml:"aliyun_viapi_imageaudit_scantext_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 请求ID
	TaobaoRequestId string `json:"taobao_request_id,omitempty" xml:"taobao_request_id,omitempty"`
	// 系统自动生成
	Data *AliyunviapiimageauditscantextData `json:"data,omitempty" xml:"data,omitempty"`
}
