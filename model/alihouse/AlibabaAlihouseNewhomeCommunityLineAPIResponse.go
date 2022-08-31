package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeCommunityLineAPIResponse 小区上下架 API返回值
// alibaba.alihouse.newhome.community.line
//
// 小区上下架
type AlibabaAlihouseNewhomeCommunityLineAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseNewhomeCommunityLineAPIResponseModel
}

// AlibabaAlihouseNewhomeCommunityLineAPIResponseModel is 小区上下架 成功返回结果
type AlibabaAlihouseNewhomeCommunityLineAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_community_line_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseNewhomeCommunityLineResult `json:"result,omitempty" xml:"result,omitempty"`
}
