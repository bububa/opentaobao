package ascpffo

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpffo"
)

// Aliexpressascpitemquery AliExpress货品查询查询API
// aliexpress.ascp.item.query
//
// AE货品查询API
func Aliexpressascpitemquery(clt *core.SDKClient, req *ascpffo.AliexpressascpitemqueryAPIRequest, session string) (*ascpffo.AliexpressascpitemqueryAPIResponse, error) {
	var resp ascpffo.AliexpressascpitemqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
