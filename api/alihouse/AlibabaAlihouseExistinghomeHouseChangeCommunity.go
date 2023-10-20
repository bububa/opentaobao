package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihouseexistinghomehousechangecommunity 房屋、房源变更所属小区
// alibaba.alihouse.existinghome.house.change.community
//
// 房屋、房源变更所属小区
func Alibabaalihouseexistinghomehousechangecommunity(clt *core.SDKClient, req *alihouse.AlibabaalihouseexistinghomehousechangecommunityAPIRequest, session string) (*alihouse.AlibabaalihouseexistinghomehousechangecommunityAPIResponse, error) {
	var resp alihouse.AlibabaalihouseexistinghomehousechangecommunityAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
