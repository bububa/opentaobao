package axindata

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axindata"
)

// Taobaoalitriptravelaxinhotelshidlistquery 阿信酒店分销-标准酒店id列表查询
// taobao.alitrip.travel.axin.hotel.shid.list.query
//
// 标准酒店id列表查询
func Taobaoalitriptravelaxinhotelshidlistquery(clt *core.SDKClient, req *axindata.TaobaoalitriptravelaxinhotelshidlistqueryAPIRequest, session string) (*axindata.TaobaoalitriptravelaxinhotelshidlistqueryAPIResponse, error) {
	var resp axindata.TaobaoalitriptravelaxinhotelshidlistqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
