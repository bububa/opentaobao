package flight

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitrippolicydomfareflowdataAPIResponse 比价工具流量详情 API返回值
// alitrip.policy.domfare.flowdata
//
// 比价工具流量详情
type AlitrippolicydomfareflowdataAPIResponse struct {
	model.CommonResponse
	AlitrippolicydomfareflowdataAPIResponseModel
}

// AlitrippolicydomfareflowdataAPIResponseModel is 比价工具流量详情 成功返回结果
type AlitrippolicydomfareflowdataAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_policy_domfare_flowdata_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回信息
	Result *AlitrippolicydomfareflowdataResult `json:"result,omitempty" xml:"result,omitempty"`
}
