package icbulogistics

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询物流运力列表 API请求
alibaba.onetouch.logistics.express.logistics.product.list

查询物流产品&揽收仓库列表
*/
type AlibabaOnetouchLogisticsExpressLogisticsProductListRequest struct {
    model.Params
}

// 初始化AlibabaOnetouchLogisticsExpressLogisticsProductListRequest对象
func NewAlibabaOnetouchLogisticsExpressLogisticsProductListRequest() *AlibabaOnetouchLogisticsExpressLogisticsProductListRequest{
    return &AlibabaOnetouchLogisticsExpressLogisticsProductListRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaOnetouchLogisticsExpressLogisticsProductListRequest) GetApiMethodName() string {
    return "alibaba.onetouch.logistics.express.logistics.product.list"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaOnetouchLogisticsExpressLogisticsProductListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
