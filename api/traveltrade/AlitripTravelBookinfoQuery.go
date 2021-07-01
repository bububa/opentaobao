package traveltrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/traveltrade"
)

/* AlitripTravelBookinfoQuery
飞猪度假-订单二次预约查询接口
alitrip.travel.bookinfo.query

飞猪度假订单二次预约详情查询接口 */
func AlitripTravelBookinfoQuery(clt *core.SDKClient, req *traveltrade.AlitripTravelBookinfoQueryAPIRequest, session string) (*traveltrade.AlitripTravelBookinfoQueryAPIResponse, error) {
	var resp traveltrade.AlitripTravelBookinfoQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
