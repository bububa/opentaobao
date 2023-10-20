package ascpffo

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpffo"
)

// Aliexpressascpwarehouseinventoryquery AliExpress在仓库存查询API
// aliexpress.ascp.warehouse.inventory.query
//
// AliExpress在仓库存查询查询API
func Aliexpressascpwarehouseinventoryquery(clt *core.SDKClient, req *ascpffo.AliexpressascpwarehouseinventoryqueryAPIRequest, session string) (*ascpffo.AliexpressascpwarehouseinventoryqueryAPIResponse, error) {
	var resp ascpffo.AliexpressascpwarehouseinventoryqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
