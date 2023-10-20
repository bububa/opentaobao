package hotel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/hotel"
)

// TaobaoXhotelDistributionInfo 飞猪分销通用酒店标准信息接口
// taobao.xhotel.distribution.info
//
// 飞猪分销通用酒店标准信息接口
func TaobaoXhotelDistributionInfo(clt *core.SDKClient, req *hotel.TaobaoXhotelDistributionInfoAPIRequest, session string) (*hotel.TaobaoXhotelDistributionInfoAPIResponse, error) {
	var resp hotel.TaobaoXhotelDistributionInfoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
