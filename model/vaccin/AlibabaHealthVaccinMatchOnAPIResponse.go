package vaccin

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHealthVaccinMatchOnAPIResponse isv自主上下线疫苗，可以选择上线还是下线 API返回值
// alibaba.health.vaccin.match.on
//
// isv自主上下线疫苗，可以选择上线还是下线
type AlibabaHealthVaccinMatchOnAPIResponse struct {
	model.CommonResponse
	AlibabaHealthVaccinMatchOnAPIResponseModel
}

// AlibabaHealthVaccinMatchOnAPIResponseModel is isv自主上下线疫苗，可以选择上线还是下线 成功返回结果
type AlibabaHealthVaccinMatchOnAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_health_vaccin_match_on_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// info
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// code
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
	// 对码是否成功
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
}
