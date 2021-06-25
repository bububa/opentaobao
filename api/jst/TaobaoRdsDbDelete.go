package jst

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/jst"
)

/* 
RDS数据库删除 
taobao.rds.db.delete

通过api删除用户RDS的数据库
*/
func TaobaoRdsDbDelete(clt *core.SDKClient, req *jst.TaobaoRdsDbDeleteRequest, session string) (*jst.TaobaoRdsDbDeleteResponse, error) {
    var resp jst.TaobaoRdsDbDeleteAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
