package jst

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

// Taobaordsdbget 查询rds下的数据库
// taobao.rds.db.get
//
// 查询rds实例下的数据库
func Taobaordsdbget(clt *core.SDKClient, req *jst.TaobaordsdbgetAPIRequest, session string) (*jst.TaobaordsdbgetAPIResponse, error) {
	var resp jst.TaobaordsdbgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
