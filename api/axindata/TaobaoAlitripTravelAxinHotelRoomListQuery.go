package axindata

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axindata"
)

// Taobaoalitriptravelaxinhotelroomlistquery 阿信酒店分销-标准酒店房型列表查询
// taobao.alitrip.travel.axin.hotel.room.list.query
//
// 标准酒店房型列表查询
func Taobaoalitriptravelaxinhotelroomlistquery(clt *core.SDKClient, req *axindata.TaobaoalitriptravelaxinhotelroomlistqueryAPIRequest, session string) (*axindata.TaobaoalitriptravelaxinhotelroomlistqueryAPIResponse, error) {
	var resp axindata.TaobaoalitriptravelaxinhotelroomlistqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
