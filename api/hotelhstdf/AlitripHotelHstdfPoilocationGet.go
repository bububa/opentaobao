package hotelhstdf

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/hotelhstdf"
)

/* AlitripHotelHstdfPoilocationGet
根据平台城市id分页查询poi location
alitrip.hotel.hstdf.poilocation.get

根据平台城市id分页查询poi location */
func AlitripHotelHstdfPoilocationGet(clt *core.SDKClient, req *hotelhstdf.AlitripHotelHstdfPoilocationGetAPIRequest, session string) (*hotelhstdf.AlitripHotelHstdfPoilocationGetAPIResponse, error) {
	var resp hotelhstdf.AlitripHotelHstdfPoilocationGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
