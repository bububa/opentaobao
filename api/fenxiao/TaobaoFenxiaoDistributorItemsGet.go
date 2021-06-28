package fenxiao

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/fenxiao"
)

/* 
查询商品下载记录 
taobao.fenxiao.distributor.items.get

供应商查询分销商商品下载记录。
*/
func TaobaoFenxiaoDistributorItemsGet(clt *core.SDKClient, req *fenxiao.TaobaoFenxiaoDistributorItemsGetRequest, session string) (*fenxiao.TaobaoFenxiaoDistributorItemsGetAPIResponse, error) {
    var resp fenxiao.TaobaoFenxiaoDistributorItemsGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
