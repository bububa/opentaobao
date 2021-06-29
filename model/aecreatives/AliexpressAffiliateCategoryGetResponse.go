package aecreatives

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
AE流量推广类目信息获取API APIResponse
aliexpress.affiliate.category.get

获取AE流量推广类目的API
*/
type AliexpressAffiliateCategoryGetAPIResponse struct {
    model.CommonResponse
    AliexpressAffiliateCategoryGetResponse
}

type AliexpressAffiliateCategoryGetResponse struct {
    XMLName xml.Name `xml:"aliexpress_affiliate_category_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    RespResult   *ResponseResult `json:"resp_result,omitempty" xml:"resp_result,omitempty"`

    
}
