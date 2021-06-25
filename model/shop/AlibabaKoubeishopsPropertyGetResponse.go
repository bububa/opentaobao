package shop

import (
    "github.com/bububa/opentaobao/model"
)

/* 
口碑店铺列表推荐 APIResponse
alibaba.koubeishops.property.get

推荐用户附近的美食门店
*/
type AlibabaKoubeishopsPropertyGetAPIResponse struct {
    model.CommonResponse
    Response *AlibabaKoubeishopsPropertyGetResponse `json:"alibaba_koubeishops_property_get_response,omitempty"`
}

type AlibabaKoubeishopsPropertyGetResponse struct {

    // 返回结果
    Result   *OpenApiSearchResult `json:"result,omitempty"`

}
