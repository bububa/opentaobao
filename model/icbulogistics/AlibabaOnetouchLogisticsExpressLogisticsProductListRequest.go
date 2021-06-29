package icbulogistics

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询物流运力列表 APIRequest
alibaba.onetouch.logistics.express.logistics.product.list

查询物流产品&揽收仓库列表
*/
type AlibabaOnetouchLogisticsExpressLogisticsProductListRequest struct {
    model.Params

}

func NewAlibabaOnetouchLogisticsExpressLogisticsProductListRequest() *AlibabaOnetouchLogisticsExpressLogisticsProductListRequest{
    return &AlibabaOnetouchLogisticsExpressLogisticsProductListRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaOnetouchLogisticsExpressLogisticsProductListRequest) GetApiMethodName() string {
    return "alibaba.onetouch.logistics.express.logistics.product.list"
}

func (r AlibabaOnetouchLogisticsExpressLogisticsProductListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


