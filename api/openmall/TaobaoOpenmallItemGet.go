package openmall

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openmall"
)

// TaobaoOpenmallItemGet 获取商品详情物料
// taobao.openmall.item.get
//
// 获取联盟开放的openmall商品
func TaobaoOpenmallItemGet(clt *core.SDKClient, req *openmall.TaobaoOpenmallItemGetAPIRequest, resp *openmall.TaobaoOpenmallItemGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
