package jst

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

// Taobaordsdbdelete RDS数据库删除
// taobao.rds.db.delete
//
// 通过api删除用户RDS的数据库
func Taobaordsdbdelete(clt *core.SDKClient, req *jst.TaobaordsdbdeleteAPIRequest, session string) (*jst.TaobaordsdbdeleteAPIResponse, error) {
	var resp jst.TaobaordsdbdeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
