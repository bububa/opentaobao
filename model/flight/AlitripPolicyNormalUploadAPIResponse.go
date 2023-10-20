package flight

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripPolicyNormalUploadAPIResponse 普通政策上传 API返回值
// alitrip.policy.normal.upload
//
// 上传普通类型的单程/往返政策
type AlitripPolicyNormalUploadAPIResponse struct {
	model.CommonResponse
	AlitripPolicyNormalUploadAPIResponseModel
}

// AlitripPolicyNormalUploadAPIResponseModel is 普通政策上传 成功返回结果
type AlitripPolicyNormalUploadAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_policy_normal_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 异步获取历史数据接口返回结果
	Result *AlitripPolicyNormalUploadResult `json:"result,omitempty" xml:"result,omitempty"`
}
