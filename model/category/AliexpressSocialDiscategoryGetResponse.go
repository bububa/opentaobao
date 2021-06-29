package category

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
展示类目获取接口 API返回值 
aliexpress.social.discategory.get

AE展示类目获取接口
*/
type AliexpressSocialDiscategoryGetAPIResponse struct {
    model.CommonResponse
    AliexpressSocialDiscategoryGetResponse
}

// 展示类目获取接口 成功返回结果
type AliexpressSocialDiscategoryGetResponse struct {
    XMLName xml.Name `xml:"aliexpress_social_discategory_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *ItemPickPagingResult `json:"result,omitempty" xml:"result,omitempty"`
}
