package axindata

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axindata"
)

// Taobaoalitriptravelaxinhotelpricebatchget 阿信酒店批量报价查询接口
// taobao.alitrip.travel.axin.hotel.price.batch.get
//
// 阿信酒店批量报价查询接口
func Taobaoalitriptravelaxinhotelpricebatchget(clt *core.SDKClient, req *axindata.TaobaoalitriptravelaxinhotelpricebatchgetAPIRequest, session string) (*axindata.TaobaoalitriptravelaxinhotelpricebatchgetAPIResponse, error) {
	var resp axindata.TaobaoalitriptravelaxinhotelpricebatchgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
