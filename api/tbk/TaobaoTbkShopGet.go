package tbk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// Taobaotbkshopget 淘宝客-推广者-店铺搜索
// taobao.tbk.shop.get
//
// 淘宝客店铺查询
func Taobaotbkshopget(clt *core.SDKClient, req *tbk.TaobaotbkshopgetAPIRequest, session string) (*tbk.TaobaotbkshopgetAPIResponse, error) {
	var resp tbk.TaobaotbkshopgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
