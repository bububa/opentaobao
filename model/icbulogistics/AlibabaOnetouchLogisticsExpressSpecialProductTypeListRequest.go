package icbulogistics

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取商品类型配置项 API请求
alibaba.onetouch.logistics.express.special.product.type.list

获取商品类型配置项
*/
type AlibabaOnetouchLogisticsExpressSpecialProductTypeListRequest struct {
    model.Params
}

// 初始化AlibabaOnetouchLogisticsExpressSpecialProductTypeListRequest对象
func NewAlibabaOnetouchLogisticsExpressSpecialProductTypeListRequest() *AlibabaOnetouchLogisticsExpressSpecialProductTypeListRequest{
    return &AlibabaOnetouchLogisticsExpressSpecialProductTypeListRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaOnetouchLogisticsExpressSpecialProductTypeListRequest) GetApiMethodName() string {
    return "alibaba.onetouch.logistics.express.special.product.type.list"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaOnetouchLogisticsExpressSpecialProductTypeListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
