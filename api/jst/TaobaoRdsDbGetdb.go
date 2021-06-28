package jst

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/jst"
)

/* 
rds获取RDS的DB 
taobao.rds.db.getdb

rds获取RDS的DB
*/
func TaobaoRdsDbGetdb(clt *core.SDKClient, req *jst.TaobaoRdsDbGetdbRequest, session string) (*jst.TaobaoRdsDbGetdbAPIResponse, error) {
    var resp jst.TaobaoRdsDbGetdbAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
