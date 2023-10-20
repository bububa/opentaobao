package moscm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/moscm"
)

// AlibabaMosIsvInventoryScrollquery 滚动查询库存数据
// alibaba.mos.isv.inventory.scrollquery
//
// 按专柜滚动查询有库存商品
func AlibabaMosIsvInventoryScrollquery(clt *core.SDKClient, req *moscm.AlibabaMosIsvInventoryScrollqueryAPIRequest, session string) (*moscm.AlibabaMosIsvInventoryScrollqueryAPIResponse, error) {
	var resp moscm.AlibabaMosIsvInventoryScrollqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
