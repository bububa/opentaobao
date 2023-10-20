package ascpffo

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpffo"
)

// Aliexpressascppoquery AliExpress采购单查询API
// aliexpress.ascp.po.query
//
// AE仓发业务采购单查询
func Aliexpressascppoquery(clt *core.SDKClient, req *ascpffo.AliexpressascppoqueryAPIRequest, session string) (*ascpffo.AliexpressascppoqueryAPIResponse, error) {
	var resp ascpffo.AliexpressascppoqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
