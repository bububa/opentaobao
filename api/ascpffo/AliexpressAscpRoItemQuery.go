package ascpffo

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpffo"
)

// Aliexpressascproitemquery AliExpress退供单明细查询API
// aliexpress.ascp.ro.item.query
//
// AE仓发 单个退供单明细查询
func Aliexpressascproitemquery(clt *core.SDKClient, req *ascpffo.AliexpressascproitemqueryAPIRequest, session string) (*ascpffo.AliexpressascproitemqueryAPIResponse, error) {
	var resp ascpffo.AliexpressascproitemqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
