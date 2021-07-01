package icbulogistics

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/icbulogistics"
)

/* 
获取商品类型配置项 
alibaba.onetouch.logistics.express.special.product.type.list

获取商品类型配置项
*/
func AlibabaOnetouchLogisticsExpressSpecialProductTypeList(clt *core.SDKClient, req *icbulogistics.AlibabaOnetouchLogisticsExpressSpecialProductTypeListAPIRequest, session string) (*icbulogistics.AlibabaOnetouchLogisticsExpressSpecialProductTypeListAPIResponse, error) {
    var resp icbulogistics.AlibabaOnetouchLogisticsExpressSpecialProductTypeListAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
