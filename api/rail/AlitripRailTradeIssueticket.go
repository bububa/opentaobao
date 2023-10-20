package rail

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/rail"
)

// Alitriprailtradeissueticket 德铁出票成功接口
// alitrip.rail.trade.issueticket
//
// 出票成功回调接口
func Alitriprailtradeissueticket(clt *core.SDKClient, req *rail.AlitriprailtradeissueticketAPIRequest, session string) (*rail.AlitriprailtradeissueticketAPIResponse, error) {
	var resp rail.AlitriprailtradeissueticketAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
