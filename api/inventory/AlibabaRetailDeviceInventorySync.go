package inventory

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/inventory"
)

// Alibabaretaildeviceinventorysync 库存同步接口
// alibaba.retail.device.inventory.sync
//
// 商库存同步接口
func Alibabaretaildeviceinventorysync(clt *core.SDKClient, req *inventory.AlibabaretaildeviceinventorysyncAPIRequest, session string) (*inventory.AlibabaretaildeviceinventorysyncAPIResponse, error) {
	var resp inventory.AlibabaretaildeviceinventorysyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
