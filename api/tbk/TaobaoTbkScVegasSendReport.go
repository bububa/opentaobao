package tbk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// Taobaotbkscvegassendreport 淘宝客-服务商-查询红包发放个数
// taobao.tbk.sc.vegas.send.report
//
// 服务商使用。可通过此接口查询对应推广者账号下的红包发放个数。
func Taobaotbkscvegassendreport(clt *core.SDKClient, req *tbk.TaobaotbkscvegassendreportAPIRequest, session string) (*tbk.TaobaotbkscvegassendreportAPIResponse, error) {
	var resp tbk.TaobaotbkscvegassendreportAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
