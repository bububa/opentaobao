package jst

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

// TaobaoRdsDbCreateaccount rds创建数据库账户
// taobao.rds.db.createaccount
//
// rds创建数据库账户
func TaobaoRdsDbCreateaccount(clt *core.SDKClient, req *jst.TaobaoRdsDbCreateaccountAPIRequest, resp *jst.TaobaoRdsDbCreateaccountAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
