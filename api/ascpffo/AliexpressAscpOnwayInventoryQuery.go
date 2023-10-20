package ascpffo

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpffo"
)

// Aliexpressascponwayinventoryquery AliExpress在途库存查询API
// aliexpress.ascp.onway.inventory.query
//
// AliExpress在途库存查询API
func Aliexpressascponwayinventoryquery(clt *core.SDKClient, req *ascpffo.AliexpressascponwayinventoryqueryAPIRequest, session string) (*ascpffo.AliexpressascponwayinventoryqueryAPIResponse, error) {
	var resp ascpffo.AliexpressascponwayinventoryqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
