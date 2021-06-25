package fenxiao

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/fenxiao"
)

/* 
失效指定用户的商品与后端商品的映射关系 
taobao.scitem.map.delete

根据后端商品Id，失效指定用户的商品与后端商品的映射关系
*/
func TaobaoScitemMapDelete(clt *core.SDKClient, req *fenxiao.TaobaoScitemMapDeleteRequest, session string) (*fenxiao.TaobaoScitemMapDeleteResponse, error) {
    var resp fenxiao.TaobaoScitemMapDeleteAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
