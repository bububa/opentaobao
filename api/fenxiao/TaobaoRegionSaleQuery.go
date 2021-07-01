package fenxiao

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/fenxiao"
)

/* 
查询商品销售区域 
taobao.region.sale.query

查询商品销售区域
*/
func TaobaoRegionSaleQuery(clt *core.SDKClient, req *fenxiao.TaobaoRegionSaleQueryAPIRequest, session string) (*fenxiao.TaobaoRegionSaleQueryAPIResponse, error) {
    var resp fenxiao.TaobaoRegionSaleQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
