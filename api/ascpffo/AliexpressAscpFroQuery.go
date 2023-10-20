package ascpffo

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpffo"
)

// AliexpressAscpFroQuery AliExpress销退单查询API
// aliexpress.ascp.fro.query
//
// AE履约销退单查询接口
func AliexpressAscpFroQuery(clt *core.SDKClient, req *ascpffo.AliexpressAscpFroQueryAPIRequest, session string) (*ascpffo.AliexpressAscpFroQueryAPIResponse, error) {
	var resp ascpffo.AliexpressAscpFroQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
