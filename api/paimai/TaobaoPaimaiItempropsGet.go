package paimai

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/paimai"
)

/* 
拍卖相关类目属性 
taobao.paimai.itemprops.get

读取拍卖相关类目属性
*/
func TaobaoPaimaiItempropsGet(clt *core.SDKClient, req *paimai.TaobaoPaimaiItempropsGetRequest, session string) (*paimai.TaobaoPaimaiItempropsGetResponse, error) {
    var resp paimai.TaobaoPaimaiItempropsGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
