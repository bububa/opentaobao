package jst

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

// Taobaordsdbcreate rds创建数据库
// taobao.rds.db.create
//
// 在rds实例里创建数据库
func Taobaordsdbcreate(clt *core.SDKClient, req *jst.TaobaordsdbcreateAPIRequest, session string) (*jst.TaobaordsdbcreateAPIResponse, error) {
	var resp jst.TaobaordsdbcreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
