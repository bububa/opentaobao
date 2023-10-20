package jst

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

// TaobaoRdsDbCreateaccount rds创建数据库账户
// taobao.rds.db.createaccount
//
// rds创建数据库账户
func TaobaoRdsDbCreateaccount(clt *core.SDKClient, req *jst.TaobaoRdsDbCreateaccountAPIRequest, session string) (*jst.TaobaoRdsDbCreateaccountAPIResponse, error) {
	var resp jst.TaobaoRdsDbCreateaccountAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
