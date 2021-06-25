package category

import (
    "github.com/bububa/opentaobao/model"
)

/* 
展示类目获取接口 APIResponse
aliexpress.social.discategory.get

AE展示类目获取接口
*/
type AliexpressSocialDiscategoryGetAPIResponse struct {
    model.CommonResponse
    Response *AliexpressSocialDiscategoryGetResponse `json:"aliexpress_social_discategory_get_response,omitempty"`
}

type AliexpressSocialDiscategoryGetResponse struct {

    // result
    Result   *ItemPickPagingResult `json:"result,omitempty"`

}
