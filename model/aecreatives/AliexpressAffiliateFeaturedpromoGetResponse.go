package aecreatives

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
联盟主题推广活动信息获取 API返回值 
aliexpress.affiliate.featuredpromo.get

获取联盟主题推广活动信息
*/
type AliexpressAffiliateFeaturedpromoGetAPIResponse struct {
    model.CommonResponse
    AliexpressAffiliateFeaturedpromoGetResponse
}

// 联盟主题推广活动信息获取 成功返回结果
type AliexpressAffiliateFeaturedpromoGetResponse struct {
    XMLName xml.Name `xml:"aliexpress_affiliate_featuredpromo_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回结果
    RespResult   *ResponseDto `json:"resp_result,omitempty" xml:"resp_result,omitempty"`
}
