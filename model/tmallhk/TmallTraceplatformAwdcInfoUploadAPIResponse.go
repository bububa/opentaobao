package tmallhk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmalltraceplatformawdcinfouploadAPIResponse AWDC提交溯源信息 API返回值
// tmall.traceplatform.awdc.info.upload
//
// 天猫溯源-AWDC-上传溯源信息
type TmalltraceplatformawdcinfouploadAPIResponse struct {
	model.CommonResponse
	TmalltraceplatformawdcinfouploadAPIResponseModel
}

// TmalltraceplatformawdcinfouploadAPIResponseModel is AWDC提交溯源信息 成功返回结果
type TmalltraceplatformawdcinfouploadAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_traceplatform_awdc_info_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *DataResult `json:"result,omitempty" xml:"result,omitempty"`
}
