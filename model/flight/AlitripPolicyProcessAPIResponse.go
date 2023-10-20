package flight

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitrippolicyprocessAPIResponse 政策进度查询 API返回值
// alitrip.policy.process
//
// 上传特殊类型的单程/往返政策
type AlitrippolicyprocessAPIResponse struct {
	model.CommonResponse
	AlitrippolicyprocessAPIResponseModel
}

// AlitrippolicyprocessAPIResponseModel is 政策进度查询 成功返回结果
type AlitrippolicyprocessAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_policy_process_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 异步获取历史数据接口返回结果
	Result *AlitrippolicyprocessResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
