package fenxiao

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/fenxiao"
)

/* 
创建分销和后端商品映射关系 
taobao.fenxiao.product.map.add

创建分销和供应链商品映射关系。
*/
func TaobaoFenxiaoProductMapAdd(clt *core.SDKClient, req *fenxiao.TaobaoFenxiaoProductMapAddRequest, session string) (*fenxiao.TaobaoFenxiaoProductMapAddResponse, error) {
    var resp fenxiao.TaobaoFenxiaoProductMapAddAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
