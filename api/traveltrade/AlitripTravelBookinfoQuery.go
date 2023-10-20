package traveltrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/traveltrade"
)

// Alitriptravelbookinfoquery 飞猪度假-订单二次预约查询接口
// alitrip.travel.bookinfo.query
//
// 飞猪度假订单二次预约详情查询接口
func Alitriptravelbookinfoquery(clt *core.SDKClient, req *traveltrade.AlitriptravelbookinfoqueryAPIRequest, session string) (*traveltrade.AlitriptravelbookinfoqueryAPIResponse, error) {
	var resp traveltrade.AlitriptravelbookinfoqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
