package flight

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitrippolicyruleuploadAPIResponse 规则政策上传 API返回值
// alitrip.policy.rule.upload
//
// 上传特殊类型的单程/往返政策
type AlitrippolicyruleuploadAPIResponse struct {
	model.CommonResponse
	AlitrippolicyruleuploadAPIResponseModel
}

// AlitrippolicyruleuploadAPIResponseModel is 规则政策上传 成功返回结果
type AlitrippolicyruleuploadAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_policy_rule_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 异步获取历史数据接口返回结果
	Result *AlitrippolicyruleuploadResult `json:"result,omitempty" xml:"result,omitempty"`
}
