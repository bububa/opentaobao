package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// TaobaoRegionPriceManage 编辑区域价格
// taobao.region.price.manage
//
// 编辑区域价格
func TaobaoRegionPriceManage(clt *core.SDKClient, req *fenxiao.TaobaoRegionPriceManageAPIRequest, resp *fenxiao.TaobaoRegionPriceManageAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
