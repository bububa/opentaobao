package tbk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// Taobaotbkdgtpwdriskreport 淘宝客-推广者-淘口令预警及拦截查询
// taobao.tbk.dg.tpwd.risk.report
//
// 淘宝客-推广者-淘口令预警及拦截查询
func Taobaotbkdgtpwdriskreport(clt *core.SDKClient, req *tbk.TaobaotbkdgtpwdriskreportAPIRequest, session string) (*tbk.TaobaotbkdgtpwdriskreportAPIResponse, error) {
	var resp tbk.TaobaotbkdgtpwdriskreportAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
