package icbulogistics

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取商品类型配置项 APIRequest
alibaba.onetouch.logistics.express.special.product.type.list

获取商品类型配置项
*/
type AlibabaOnetouchLogisticsExpressSpecialProductTypeListRequest struct {
    model.Params

}

func NewAlibabaOnetouchLogisticsExpressSpecialProductTypeListRequest() *AlibabaOnetouchLogisticsExpressSpecialProductTypeListRequest{
    return &AlibabaOnetouchLogisticsExpressSpecialProductTypeListRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaOnetouchLogisticsExpressSpecialProductTypeListRequest) GetApiMethodName() string {
    return "alibaba.onetouch.logistics.express.special.product.type.list"
}

func (r AlibabaOnetouchLogisticsExpressSpecialProductTypeListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


