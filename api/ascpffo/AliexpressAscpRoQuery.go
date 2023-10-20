package ascpffo

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpffo"
)

// Aliexpressascproquery AliExpress退供单查询API
// aliexpress.ascp.ro.query
//
// AE仓发商家单个退供单查询接口
func Aliexpressascproquery(clt *core.SDKClient, req *ascpffo.AliexpressascproqueryAPIRequest, session string) (*ascpffo.AliexpressascproqueryAPIResponse, error) {
	var resp ascpffo.AliexpressascproqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
