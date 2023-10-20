package axindata

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axindata"
)

// Taobaoalitriptravelaxinhotelroommatch 阿信酒店房型匹配
// taobao.alitrip.travel.axin.hotel.room.match
//
// 阿信酒店房型匹配
func Taobaoalitriptravelaxinhotelroommatch(clt *core.SDKClient, req *axindata.TaobaoalitriptravelaxinhotelroommatchAPIRequest, session string) (*axindata.TaobaoalitriptravelaxinhotelroommatchAPIResponse, error) {
	var resp axindata.TaobaoalitriptravelaxinhotelroommatchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
