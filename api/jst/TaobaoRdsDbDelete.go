package jst

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

// TaobaoRdsDbDelete RDS数据库删除
// taobao.rds.db.delete
//
// 通过api删除用户RDS的数据库
func TaobaoRdsDbDelete(clt *core.SDKClient, req *jst.TaobaoRdsDbDeleteAPIRequest, resp *jst.TaobaoRdsDbDeleteAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
