package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeEcodeUpdateAPIResponse 新房货变更E码 API返回值
// alibaba.alihouse.newhome.ecode.update
//
// 新房货变更E码
type AlibabaAlihouseNewhomeEcodeUpdateAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseNewhomeEcodeUpdateAPIResponseModel
}

// AlibabaAlihouseNewhomeEcodeUpdateAPIResponseModel is 新房货变更E码 成功返回结果
type AlibabaAlihouseNewhomeEcodeUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_ecode_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseNewhomeEcodeUpdateResult `json:"result,omitempty" xml:"result,omitempty"`
}
