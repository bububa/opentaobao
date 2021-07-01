package hotel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/hotel"
)

/* AlitripHotelSearchListGet
酒店搜索List接口
alitrip.hotel.search.list.get

酒店搜索List接口 */
func AlitripHotelSearchListGet(clt *core.SDKClient, req *hotel.AlitripHotelSearchListGetAPIRequest, session string) (*hotel.AlitripHotelSearchListGetAPIResponse, error) {
	var resp hotel.AlitripHotelSearchListGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
