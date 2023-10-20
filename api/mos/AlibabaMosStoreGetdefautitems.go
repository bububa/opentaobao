package mos

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mos"
)

// Alibabamosstoregetdefautitems 获取默认状态下商品列表
// alibaba.mos.store.getdefautitems
//
// 获取默认状态下商品列表
func Alibabamosstoregetdefautitems(clt *core.SDKClient, req *mos.AlibabamosstoregetdefautitemsAPIRequest, session string) (*mos.AlibabamosstoregetdefautitemsAPIResponse, error) {
	var resp mos.AlibabamosstoregetdefautitemsAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
