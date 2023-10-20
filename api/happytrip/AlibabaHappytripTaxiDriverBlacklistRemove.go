package happytrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/happytrip"
)

// AlibabaHappytripTaxiDriverBlacklistRemove 移除司机黑名单
// alibaba.happytrip.taxi.driver.blacklist.remove
//
// 移除司机黑名单
func AlibabaHappytripTaxiDriverBlacklistRemove(clt *core.SDKClient, req *happytrip.AlibabaHappytripTaxiDriverBlacklistRemoveAPIRequest, resp *happytrip.AlibabaHappytripTaxiDriverBlacklistRemoveAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
