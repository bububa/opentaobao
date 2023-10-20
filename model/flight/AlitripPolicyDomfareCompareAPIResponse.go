package flight

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitrippolicydomfarecompareAPIResponse 比价工具 API返回值
// alitrip.policy.domfare.compare
//
// 比价工具
type AlitrippolicydomfarecompareAPIResponse struct {
	model.CommonResponse
	AlitrippolicydomfarecompareAPIResponseModel
}

// AlitrippolicydomfarecompareAPIResponseModel is 比价工具 成功返回结果
type AlitrippolicydomfarecompareAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_policy_domfare_compare_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回信息
	Result *AlitrippolicydomfarecompareResult `json:"result,omitempty" xml:"result,omitempty"`
}
