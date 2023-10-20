package idle

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidlereportmediauploadAPIResponse 验货报告上传文件 API返回值
// alibaba.idle.report.media.upload
//
// 服务商上传文件、图片
type AlibabaidlereportmediauploadAPIResponse struct {
	model.CommonResponse
	AlibabaidlereportmediauploadAPIResponseModel
}

// AlibabaidlereportmediauploadAPIResponseModel is 验货报告上传文件 成功返回结果
type AlibabaidlereportmediauploadAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_report_media_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}
