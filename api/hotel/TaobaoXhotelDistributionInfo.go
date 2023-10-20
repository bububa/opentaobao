package hotel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/hotel"
)

// Taobaoxhoteldistributioninfo 飞猪分销通用酒店标准信息接口
// taobao.xhotel.distribution.info
//
// 飞猪分销通用酒店标准信息接口
func Taobaoxhoteldistributioninfo(clt *core.SDKClient, req *hotel.TaobaoxhoteldistributioninfoAPIRequest, session string) (*hotel.TaobaoxhoteldistributioninfoAPIResponse, error) {
	var resp hotel.TaobaoxhoteldistributioninfoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
