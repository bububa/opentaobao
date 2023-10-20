package rail

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/rail"
)

// Alitriprailtradecloseticket 出票失败关单接口
// alitrip.rail.trade.closeticket
//
// 出票成功回调接口
func Alitriprailtradecloseticket(clt *core.SDKClient, req *rail.AlitriprailtradecloseticketAPIRequest, session string) (*rail.AlitriprailtradecloseticketAPIResponse, error) {
	var resp rail.AlitriprailtradecloseticketAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
