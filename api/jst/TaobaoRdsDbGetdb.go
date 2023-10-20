package jst

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

// TaobaoRdsDbGetdb rds获取RDS的DB
// taobao.rds.db.getdb
//
// rds获取RDS的DB
func TaobaoRdsDbGetdb(clt *core.SDKClient, req *jst.TaobaoRdsDbGetdbAPIRequest, resp *jst.TaobaoRdsDbGetdbAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
