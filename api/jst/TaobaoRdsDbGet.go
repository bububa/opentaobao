package jst

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/jst"
)

/* 
查询rds下的数据库 
taobao.rds.db.get

查询rds实例下的数据库
*/
func TaobaoRdsDbGet(clt *core.SDKClient, req *jst.TaobaoRdsDbGetAPIRequest, session string) (*jst.TaobaoRdsDbGetAPIResponse, error) {
    var resp jst.TaobaoRdsDbGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
