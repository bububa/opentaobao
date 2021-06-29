package category

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
展示类目获取接口 APIResponse
aliexpress.social.discategory.get

AE展示类目获取接口
*/
type AliexpressSocialDiscategoryGetAPIResponse struct {
    model.CommonResponse
    AliexpressSocialDiscategoryGetResponse
}

type AliexpressSocialDiscategoryGetResponse struct {
    XMLName xml.Name `xml:"aliexpress_social_discategory_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *ItemPickPagingResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
