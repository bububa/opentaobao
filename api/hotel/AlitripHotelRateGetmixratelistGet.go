package hotel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/hotel"
)

// Alitriphotelrategetmixratelistget 酒店评论接口
// alitrip.hotel.rate.getmixratelist.get
//
// 酒店评论接口
func Alitriphotelrategetmixratelistget(clt *core.SDKClient, req *hotel.AlitriphotelrategetmixratelistgetAPIRequest, session string) (*hotel.AlitriphotelrategetmixratelistgetAPIResponse, error) {
	var resp hotel.AlitriphotelrategetmixratelistgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
