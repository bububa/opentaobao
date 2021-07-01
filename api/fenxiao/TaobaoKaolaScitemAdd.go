package fenxiao

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/fenxiao"
)

/* 
考拉货品新增接口 
taobao.kaola.scitem.add

考拉货品新增接口
*/
func TaobaoKaolaScitemAdd(clt *core.SDKClient, req *fenxiao.TaobaoKaolaScitemAddAPIRequest, session string) (*fenxiao.TaobaoKaolaScitemAddAPIResponse, error) {
    var resp fenxiao.TaobaoKaolaScitemAddAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
