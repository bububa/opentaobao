package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseExistinghomeHouseChangeCommunity 房屋、房源变更所属小区
// alibaba.alihouse.existinghome.house.change.community
//
// 房屋、房源变更所属小区
func AlibabaAlihouseExistinghomeHouseChangeCommunity(clt *core.SDKClient, req *alihouse.AlibabaAlihouseExistinghomeHouseChangeCommunityAPIRequest, resp *alihouse.AlibabaAlihouseExistinghomeHouseChangeCommunityAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
