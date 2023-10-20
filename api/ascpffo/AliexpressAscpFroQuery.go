package ascpffo

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpffo"
)

// Aliexpressascpfroquery AliExpress销退单查询API
// aliexpress.ascp.fro.query
//
// AE履约销退单查询接口
func Aliexpressascpfroquery(clt *core.SDKClient, req *ascpffo.AliexpressascpfroqueryAPIRequest, session string) (*ascpffo.AliexpressascpfroqueryAPIResponse, error) {
	var resp ascpffo.AliexpressascpfroqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
