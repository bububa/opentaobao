package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseExistinghomeHouseChangeCommunity 房屋、房源变更所属小区
// alibaba.alihouse.existinghome.house.change.community
//
// 房屋、房源变更所属小区
func AlibabaAlihouseExistinghomeHouseChangeCommunity(clt *core.SDKClient, req *alihouse.AlibabaAlihouseExistinghomeHouseChangeCommunityAPIRequest, session string) (*alihouse.AlibabaAlihouseExistinghomeHouseChangeCommunityAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseExistinghomeHouseChangeCommunityAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
