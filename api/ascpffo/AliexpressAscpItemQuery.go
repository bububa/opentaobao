package ascpffo

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpffo"
)

// AliexpressAscpItemQuery AliExpress货品查询查询API
// aliexpress.ascp.item.query
//
// AE货品查询API
func AliexpressAscpItemQuery(clt *core.SDKClient, req *ascpffo.AliexpressAscpItemQueryAPIRequest, session string) (*ascpffo.AliexpressAscpItemQueryAPIResponse, error) {
	var resp ascpffo.AliexpressAscpItemQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
