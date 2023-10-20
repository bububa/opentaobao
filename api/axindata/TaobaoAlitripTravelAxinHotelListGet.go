package axindata

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axindata"
)

// Taobaoalitriptravelaxinhotellistget 标准酒店信息查询-阿信
// taobao.alitrip.travel.axin.hotel.list.get
//
// 阿信酒店分销基础数据查询
func Taobaoalitriptravelaxinhotellistget(clt *core.SDKClient, req *axindata.TaobaoalitriptravelaxinhotellistgetAPIRequest, session string) (*axindata.TaobaoalitriptravelaxinhotellistgetAPIResponse, error) {
	var resp axindata.TaobaoalitriptravelaxinhotellistgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
