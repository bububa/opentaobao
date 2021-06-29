package westcrm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
关注品牌号 API请求
alibaba.ibizapi.brand.subscribe

关注品牌号服务
*/
type AlibabaIbizapiBrandSubscribeRequest struct {
    model.Params
}

// 初始化AlibabaIbizapiBrandSubscribeRequest对象
func NewAlibabaIbizapiBrandSubscribeRequest() *AlibabaIbizapiBrandSubscribeRequest{
    return &AlibabaIbizapiBrandSubscribeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIbizapiBrandSubscribeRequest) GetApiMethodName() string {
    return "alibaba.ibizapi.brand.subscribe"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIbizapiBrandSubscribeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
