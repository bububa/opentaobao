package paimai

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/paimai"
)

// TaobaoPaimaiItemCooperationSync 商品同步
// taobao.paimai.item.cooperation.sync
//
// 商品同步
func TaobaoPaimaiItemCooperationSync(clt *core.SDKClient, req *paimai.TaobaoPaimaiItemCooperationSyncAPIRequest, resp *paimai.TaobaoPaimaiItemCooperationSyncAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
