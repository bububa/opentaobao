package happytrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/happytrip"
)

/* AlibabaHappytripTaxiIdGet
获取请求id
alibaba.happytrip.taxi.id.get

获取订单号 */
func AlibabaHappytripTaxiIdGet(clt *core.SDKClient, req *happytrip.AlibabaHappytripTaxiIdGetAPIRequest, session string) (*happytrip.AlibabaHappytripTaxiIdGetAPIResponse, error) {
	var resp happytrip.AlibabaHappytripTaxiIdGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
