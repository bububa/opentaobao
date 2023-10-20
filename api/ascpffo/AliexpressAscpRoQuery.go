package ascpffo

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpffo"
)

// AliexpressAscpRoQuery AliExpress退供单查询API
// aliexpress.ascp.ro.query
//
// AE仓发商家单个退供单查询接口
func AliexpressAscpRoQuery(clt *core.SDKClient, req *ascpffo.AliexpressAscpRoQueryAPIRequest, resp *ascpffo.AliexpressAscpRoQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
