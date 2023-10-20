package hotelalliance

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/hotelalliance"
)

// Alitriphotelalliancesettleordersyn 菲住联盟分账成功订单同步
// alitrip.hotel.alliance.settle.order.syn
//
// 用于菲住联盟分账成功订单同步
func Alitriphotelalliancesettleordersyn(clt *core.SDKClient, req *hotelalliance.AlitriphotelalliancesettleordersynAPIRequest, session string) (*hotelalliance.AlitriphotelalliancesettleordersynAPIResponse, error) {
	var resp hotelalliance.AlitriphotelalliancesettleordersynAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
