package charity

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabacharityuserbindgeturiAPIResponse 获取用户绑定uri API返回值
// alibaba.charity.user.bind.geturi
//
// 获取用户绑定uri
type AlibabacharityuserbindgeturiAPIResponse struct {
	model.CommonResponse
	AlibabacharityuserbindgeturiAPIResponseModel
}

// AlibabacharityuserbindgeturiAPIResponseModel is 获取用户绑定uri 成功返回结果
type AlibabacharityuserbindgeturiAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_charity_user_bind_geturi_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *ThreehoursResult `json:"result,omitempty" xml:"result,omitempty"`
}
