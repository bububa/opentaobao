package promotion

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabalafitesellerbenefitlistAPIResponse 商家自运营权益列表 API返回值
// alibaba.lafite.seller.benefit.list
//
// 小程序isv可使用该接口获取权益列表
type AlibabalafitesellerbenefitlistAPIResponse struct {
	model.CommonResponse
	AlibabalafitesellerbenefitlistAPIResponseModel
}

// AlibabalafitesellerbenefitlistAPIResponseModel is 商家自运营权益列表 成功返回结果
type AlibabalafitesellerbenefitlistAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_lafite_seller_benefit_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabalafitesellerbenefitlistResult `json:"result,omitempty" xml:"result,omitempty"`
}
