package ju

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabajhscommunityidentitystoreAPIResponse 用户信息存储 API返回值
// alibaba.jhs.community.identity.store
//
// 用户信息存储
type AlibabajhscommunityidentitystoreAPIResponse struct {
	model.CommonResponse
	AlibabajhscommunityidentitystoreAPIResponseModel
}

// AlibabajhscommunityidentitystoreAPIResponseModel is 用户信息存储 成功返回结果
type AlibabajhscommunityidentitystoreAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_jhs_community_identity_store_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息
	Error string `json:"error,omitempty" xml:"error,omitempty"`
	// 操作是否成功
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
}
