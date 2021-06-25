package fenxiao

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/fenxiao"
)

/* 
创建IC商品与后端商品的映射关系 
taobao.scitem.map.add

创建IC商品或分销商品与后端商品的映射关系
*/
func TaobaoScitemMapAdd(clt *core.SDKClient, req *fenxiao.TaobaoScitemMapAddRequest, session string) (*fenxiao.TaobaoScitemMapAddResponse, error) {
    var resp fenxiao.TaobaoScitemMapAddAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
