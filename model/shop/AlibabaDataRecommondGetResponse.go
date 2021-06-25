package shop

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取推荐信息 APIResponse
alibaba.data.recommond.get

获取优惠券信息，仅作客户端鉴权虚拟api使用
*/
type AlibabaDataRecommondGetAPIResponse struct {
    model.CommonResponse
    Response *AlibabaDataRecommondGetResponse `json:"alibaba_data_recommond_get_response,omitempty"`
}

type AlibabaDataRecommondGetResponse struct {

    // 客户端鉴权虚拟api使用
    Unnamed   string `json:"unnamed,omitempty"`

}
