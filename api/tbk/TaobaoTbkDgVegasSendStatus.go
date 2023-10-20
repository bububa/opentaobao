package tbk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// Taobaotbkdgvegassendstatus 淘宝客-推广者-红包领取状态查询
// taobao.tbk.dg.vegas.send.status
//
// 淘宝客传入用户标识的信息，查询该用户是否有当前阶段待核销的红包（淘客接入前需签署协议 https://pub.alimama.com/fourth/protocol/common.htm?key=hangye_laxin）
func Taobaotbkdgvegassendstatus(clt *core.SDKClient, req *tbk.TaobaotbkdgvegassendstatusAPIRequest, session string) (*tbk.TaobaotbkdgvegassendstatusAPIResponse, error) {
	var resp tbk.TaobaotbkdgvegassendstatusAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
