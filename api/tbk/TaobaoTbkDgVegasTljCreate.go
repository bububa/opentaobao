package tbk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// Taobaotbkdgvegastljcreate 淘宝客-推广者-淘礼金创建
// taobao.tbk.dg.vegas.tlj.create
//
// 创建淘礼金
func Taobaotbkdgvegastljcreate(clt *core.SDKClient, req *tbk.TaobaotbkdgvegastljcreateAPIRequest, session string) (*tbk.TaobaotbkdgvegastljcreateAPIResponse, error) {
	var resp tbk.TaobaotbkdgvegastljcreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
