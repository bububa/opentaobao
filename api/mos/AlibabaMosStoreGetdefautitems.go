package mos

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mos"
)

// AlibabaMosStoreGetdefautitems 获取默认状态下商品列表
// alibaba.mos.store.getdefautitems
//
// 获取默认状态下商品列表
func AlibabaMosStoreGetdefautitems(clt *core.SDKClient, req *mos.AlibabaMosStoreGetdefautitemsAPIRequest, resp *mos.AlibabaMosStoreGetdefautitemsAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
