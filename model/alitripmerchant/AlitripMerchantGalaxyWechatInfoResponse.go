package alitripmerchant

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
星河-获取微信用户的信息 APIResponse
alitrip.merchant.galaxy.wechat.info

获取微信用户的openId和unionId
*/
type AlitripMerchantGalaxyWechatInfoAPIResponse struct {
    model.CommonResponse
    AlitripMerchantGalaxyWechatInfoResponse
}

type AlitripMerchantGalaxyWechatInfoResponse struct {
    XMLName xml.Name `xml:"alitrip_merchant_galaxy_wechat_info_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 默认描述
    
    Result   *Response `json:"result,omitempty" xml:"result,omitempty"`

    
}
