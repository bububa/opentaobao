package ju

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabajhscommunitysubmittingtextAPIResponse 聚划算社群动态文案下发接口 API返回值
// alibaba.jhs.community.submitting.text
//
// 聚划算社群动态文案下发接口
type AlibabajhscommunitysubmittingtextAPIResponse struct {
	model.CommonResponse
	AlibabajhscommunitysubmittingtextAPIResponseModel
}

// AlibabajhscommunitysubmittingtextAPIResponseModel is 聚划算社群动态文案下发接口 成功返回结果
type AlibabajhscommunitysubmittingtextAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_jhs_community_submitting_text_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// error
	Error string `json:"error,omitempty" xml:"error,omitempty"`
	// 数据结果
	Data *AlibabajhscommunitysubmittingtextData `json:"data,omitempty" xml:"data,omitempty"`
}
