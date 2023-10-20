package tmallservice

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterSpserviceorderEpocUploadAPIResponse 电子保单文件上传接口 API返回值
// tmall.servicecenter.spserviceorder.epoc.upload
//
// 电子保单文件上传接口
type TmallServicecenterSpserviceorderEpocUploadAPIResponse struct {
	model.CommonResponse
	TmallServicecenterSpserviceorderEpocUploadAPIResponseModel
}

// TmallServicecenterSpserviceorderEpocUploadAPIResponseModel is 电子保单文件上传接口 成功返回结果
type TmallServicecenterSpserviceorderEpocUploadAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_spserviceorder_epoc_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *FulfilplatformResult `json:"result,omitempty" xml:"result,omitempty"`
}
