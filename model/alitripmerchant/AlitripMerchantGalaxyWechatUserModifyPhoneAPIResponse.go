package alitripmerchant

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxywechatusermodifyphoneAPIResponse DFC-ID用户换绑手机号 API返回值
// alitrip.merchant.galaxy.wechat.user.modify.phone
//
// DFC-ID用户换绑手机号
type AlitripmerchantgalaxywechatusermodifyphoneAPIResponse struct {
	model.CommonResponse
	AlitripmerchantgalaxywechatusermodifyphoneAPIResponseModel
}

// AlitripmerchantgalaxywechatusermodifyphoneAPIResponseModel is DFC-ID用户换绑手机号 成功返回结果
type AlitripmerchantgalaxywechatusermodifyphoneAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_wechat_user_modify_phone_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlitripmerchantgalaxywechatusermodifyphoneResponse `json:"result,omitempty" xml:"result,omitempty"`
}
