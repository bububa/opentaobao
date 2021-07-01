package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

/* TaobaoRegionPriceManage
编辑区域价格
taobao.region.price.manage

编辑区域价格 */
func TaobaoRegionPriceManage(clt *core.SDKClient, req *fenxiao.TaobaoRegionPriceManageAPIRequest, session string) (*fenxiao.TaobaoRegionPriceManageAPIResponse, error) {
	var resp fenxiao.TaobaoRegionPriceManageAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
