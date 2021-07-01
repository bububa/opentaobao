package ascpffo

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpffo"
)

/* AliexpressAscpRoQuery
AliExpress退供单查询API
aliexpress.ascp.ro.query

AE仓发商家单个退供单查询接口 */
func AliexpressAscpRoQuery(clt *core.SDKClient, req *ascpffo.AliexpressAscpRoQueryAPIRequest, session string) (*ascpffo.AliexpressAscpRoQueryAPIResponse, error) {
	var resp ascpffo.AliexpressAscpRoQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
