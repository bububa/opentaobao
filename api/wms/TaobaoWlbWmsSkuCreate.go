package wms

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wms"
)

/* TaobaoWlbWmsSkuCreate
商品同步
taobao.wlb.wms.sku.create

商品同步 */
func TaobaoWlbWmsSkuCreate(clt *core.SDKClient, req *wms.TaobaoWlbWmsSkuCreateAPIRequest, session string) (*wms.TaobaoWlbWmsSkuCreateAPIResponse, error) {
	var resp wms.TaobaoWlbWmsSkuCreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
