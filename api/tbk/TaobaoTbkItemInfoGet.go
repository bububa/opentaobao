package tbk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// Taobaotbkiteminfoget 淘宝客-公用-淘宝客商品详情查询(简版)
// taobao.tbk.item.info.get
//
// 淘宝客商品详情查询（简版）
func Taobaotbkiteminfoget(clt *core.SDKClient, req *tbk.TaobaotbkiteminfogetAPIRequest, session string) (*tbk.TaobaotbkiteminfogetAPIResponse, error) {
	var resp tbk.TaobaotbkiteminfogetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
