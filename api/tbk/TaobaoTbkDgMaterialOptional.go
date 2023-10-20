package tbk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// Taobaotbkdgmaterialoptional 淘宝客-推广者-物料搜索
// taobao.tbk.dg.material.optional
//
// 通用物料搜索API（导购）
func Taobaotbkdgmaterialoptional(clt *core.SDKClient, req *tbk.TaobaotbkdgmaterialoptionalAPIRequest, session string) (*tbk.TaobaotbkdgmaterialoptionalAPIResponse, error) {
	var resp tbk.TaobaotbkdgmaterialoptionalAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
