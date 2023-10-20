package lstvending

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lstvending"
)

// Alibabalstvendinggoodssave 自动售卖机商品回传
// alibaba.lst.vending.goods.save
//
// 零售通自动售卖机商品数据回流。
func Alibabalstvendinggoodssave(clt *core.SDKClient, req *lstvending.AlibabalstvendinggoodssaveAPIRequest, session string) (*lstvending.AlibabalstvendinggoodssaveAPIResponse, error) {
	var resp lstvending.AlibabalstvendinggoodssaveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
