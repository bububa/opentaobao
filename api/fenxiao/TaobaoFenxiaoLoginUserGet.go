package fenxiao

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/fenxiao"
)

/* 
获取分销用户登录信息 
taobao.fenxiao.login.user.get

获取用户登录信息
*/
func TaobaoFenxiaoLoginUserGet(clt *core.SDKClient, req *fenxiao.TaobaoFenxiaoLoginUserGetAPIRequest, session string) (*fenxiao.TaobaoFenxiaoLoginUserGetAPIResponse, error) {
    var resp fenxiao.TaobaoFenxiaoLoginUserGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
