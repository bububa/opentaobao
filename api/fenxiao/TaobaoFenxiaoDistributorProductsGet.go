package fenxiao

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/fenxiao"
)

/* 
分销商查询产品信息 
taobao.fenxiao.distributor.products.get

分销商查询供应商产品信息
*/
func TaobaoFenxiaoDistributorProductsGet(clt *core.SDKClient, req *fenxiao.TaobaoFenxiaoDistributorProductsGetRequest, session string) (*fenxiao.TaobaoFenxiaoDistributorProductsGetAPIResponse, error) {
    var resp fenxiao.TaobaoFenxiaoDistributorProductsGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
