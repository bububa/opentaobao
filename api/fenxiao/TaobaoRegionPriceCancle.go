package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// TaobaoRegionPriceCancle 取消区域价格
// taobao.region.price.cancle
//
// 取消区域价格
func TaobaoRegionPriceCancle(clt *core.SDKClient, req *fenxiao.TaobaoRegionPriceCancleAPIRequest, resp *fenxiao.TaobaoRegionPriceCancleAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
