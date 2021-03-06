package product

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaGspSupplyImageUploadAPIResponse gsp图片上传 API返回值
// alibaba.gsp.supply.image.upload
//
// 上传图片至目标海外平台的素材空间
type AlibabaGspSupplyImageUploadAPIResponse struct {
	model.CommonResponse
	AlibabaGspSupplyImageUploadAPIResponseModel
}

// AlibabaGspSupplyImageUploadAPIResponseModel is gsp图片上传 成功返回结果
type AlibabaGspSupplyImageUploadAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_gsp_supply_image_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码
	ServiceErrorCode string `json:"service_error_code,omitempty" xml:"service_error_code,omitempty"`
	// 错误信息
	ServiceErrorMsg string `json:"service_error_msg,omitempty" xml:"service_error_msg,omitempty"`
	// 数据
	Model *UploadImageResp `json:"model,omitempty" xml:"model,omitempty"`
	// 执行结果
	ServiceSuccess bool `json:"service_success,omitempty" xml:"service_success,omitempty"`
	// 是否重试
	NeedRetry bool `json:"need_retry,omitempty" xml:"need_retry,omitempty"`
}
