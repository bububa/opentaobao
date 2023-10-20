package happytrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/happytrip"
)

// AlibabaHappytripTaxiDriverLocationGet 司机位置
// alibaba.happytrip.taxi.driver.location.get
//
// 获取司机实时位置
func AlibabaHappytripTaxiDriverLocationGet(clt *core.SDKClient, req *happytrip.AlibabaHappytripTaxiDriverLocationGetAPIRequest, resp *happytrip.AlibabaHappytripTaxiDriverLocationGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
