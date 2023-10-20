package tbk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// Taobaotbkdgvegastljstop 淘宝客-推广者-淘礼金暂停发放
// taobao.tbk.dg.vegas.tlj.stop
//
// 淘宝客推广者可对已经创建的淘礼金暂停发放
func Taobaotbkdgvegastljstop(clt *core.SDKClient, req *tbk.TaobaotbkdgvegastljstopAPIRequest, session string) (*tbk.TaobaotbkdgvegastljstopAPIResponse, error) {
	var resp tbk.TaobaotbkdgvegastljstopAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
