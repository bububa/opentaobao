package shop

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取商品 APIResponse
alibaba.data.item.get

获取商品信息，作为客户端Weex鉴权的虚拟api
*/
type AlibabaDataItemGetAPIResponse struct {
    model.CommonResponse
    Response *AlibabaDataItemGetResponse `json:"alibaba_data_item_get_response,omitempty"`
}

type AlibabaDataItemGetResponse struct {

    // 获取商品信息，作为客户端Weex鉴权的虚拟api
    Unnamed   string `json:"unnamed,omitempty"`

}
