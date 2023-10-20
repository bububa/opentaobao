package alitripmerchant

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxywechatinfoAPIResponse 星河-获取微信用户的信息 API返回值
// alitrip.merchant.galaxy.wechat.info
//
// 获取微信用户的openId和unionId
type AlitripmerchantgalaxywechatinfoAPIResponse struct {
	model.CommonResponse
	AlitripmerchantgalaxywechatinfoAPIResponseModel
}

// AlitripmerchantgalaxywechatinfoAPIResponseModel is 星河-获取微信用户的信息 成功返回结果
type AlitripmerchantgalaxywechatinfoAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_wechat_info_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 默认描述
	Result *AlitripmerchantgalaxywechatinfoResponse `json:"result,omitempty" xml:"result,omitempty"`
}
