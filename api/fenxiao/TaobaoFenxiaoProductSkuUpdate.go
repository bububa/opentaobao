package fenxiao

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/fenxiao"
)

/* 
产品sku编辑接口 
taobao.fenxiao.product.sku.update

产品SKU信息更新
*/
func TaobaoFenxiaoProductSkuUpdate(clt *core.SDKClient, req *fenxiao.TaobaoFenxiaoProductSkuUpdateRequest, session string) (*fenxiao.TaobaoFenxiaoProductSkuUpdateResponse, error) {
    var resp fenxiao.TaobaoFenxiaoProductSkuUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
