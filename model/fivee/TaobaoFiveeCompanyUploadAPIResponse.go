package fivee

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaofiveecompanyuploadAPIResponse 上传商信息接口 API返回值
// taobao.fivee.company.upload
//
// 资质共享平台上传资质证照
type TaobaofiveecompanyuploadAPIResponse struct {
	model.CommonResponse
	TaobaofiveecompanyuploadAPIResponseModel
}

// TaobaofiveecompanyuploadAPIResponseModel is 上传商信息接口 成功返回结果
type TaobaofiveecompanyuploadAPIResponseModel struct {
	XMLName xml.Name `xml:"fivee_company_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回素材id
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// code
	CodeT int64 `json:"code_t,omitempty" xml:"code_t,omitempty"`
	// 是否成功
	SuccessT bool `json:"success_t,omitempty" xml:"success_t,omitempty"`
}
