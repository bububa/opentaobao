package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

/* AlitripBtripCityCarApplyAdd
三方市内用车申请单同步
alitrip.btrip.city.car.apply.add

三方市内用车申请单同步 */
func AlitripBtripCityCarApplyAdd(clt *core.SDKClient, req *btrip.AlitripBtripCityCarApplyAddAPIRequest, session string) (*btrip.AlitripBtripCityCarApplyAddAPIResponse, error) {
	var resp btrip.AlitripBtripCityCarApplyAddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
