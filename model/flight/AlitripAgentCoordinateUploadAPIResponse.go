package flight

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripAgentCoordinateUploadAPIResponse 协同单附件凭证上传 API返回值
// alitrip.agent.coordinate.upload
//
// 协同单附件凭证上传
type AlitripAgentCoordinateUploadAPIResponse struct {
	model.CommonResponse
	AlitripAgentCoordinateUploadAPIResponseModel
}

// AlitripAgentCoordinateUploadAPIResponseModel is 协同单附件凭证上传 成功返回结果
type AlitripAgentCoordinateUploadAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_agent_coordinate_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 协同单接口返回结果
	Result *AlitripAgentCoordinateUploadResult `json:"result,omitempty" xml:"result,omitempty"`
}
