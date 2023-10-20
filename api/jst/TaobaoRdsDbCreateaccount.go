package jst

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

// Taobaordsdbcreateaccount rds创建数据库账户
// taobao.rds.db.createaccount
//
// rds创建数据库账户
func Taobaordsdbcreateaccount(clt *core.SDKClient, req *jst.TaobaordsdbcreateaccountAPIRequest, session string) (*jst.TaobaordsdbcreateaccountAPIResponse, error) {
	var resp jst.TaobaordsdbcreateaccountAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
