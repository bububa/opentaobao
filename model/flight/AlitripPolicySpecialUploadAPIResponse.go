package flight

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitrippolicyspecialuploadAPIResponse 特殊政策上传 API返回值
// alitrip.policy.special.upload
//
// 上传特殊类型的单程/往返政策
type AlitrippolicyspecialuploadAPIResponse struct {
	model.CommonResponse
	AlitrippolicyspecialuploadAPIResponseModel
}

// AlitrippolicyspecialuploadAPIResponseModel is 特殊政策上传 成功返回结果
type AlitrippolicyspecialuploadAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_policy_special_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 异步获取历史数据接口返回结果
	Result *AlitrippolicyspecialuploadResult `json:"result,omitempty" xml:"result,omitempty"`
}
