package flight

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitrippolicynormalcompressionuploadAPIResponse 大批量上传普通类型的单程/往返政策 API返回值
// alitrip.policy.normal.compression.upload
//
// 大批量上传普通类型的单程/往返政策
type AlitrippolicynormalcompressionuploadAPIResponse struct {
	model.CommonResponse
	AlitrippolicynormalcompressionuploadAPIResponseModel
}

// AlitrippolicynormalcompressionuploadAPIResponseModel is 大批量上传普通类型的单程/往返政策 成功返回结果
type AlitrippolicynormalcompressionuploadAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_policy_normal_compression_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 异步获取历史数据接口返回结果
	Result *ResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
