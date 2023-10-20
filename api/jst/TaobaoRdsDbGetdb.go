package jst

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

// Taobaordsdbgetdb rds获取RDS的DB
// taobao.rds.db.getdb
//
// rds获取RDS的DB
func Taobaordsdbgetdb(clt *core.SDKClient, req *jst.TaobaordsdbgetdbAPIRequest, session string) (*jst.TaobaordsdbgetdbAPIResponse, error) {
	var resp jst.TaobaordsdbgetdbAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
