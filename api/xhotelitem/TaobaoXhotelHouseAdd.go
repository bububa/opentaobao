package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// TaobaoXhotelHouseAdd 非标准民宿房源添加
// taobao.xhotel.house.add
//
// 添加酒店或更新酒店
func TaobaoXhotelHouseAdd(clt *core.SDKClient, req *xhotelitem.TaobaoXhotelHouseAddAPIRequest, session string) (*xhotelitem.TaobaoXhotelHouseAddAPIResponse, error) {
	var resp xhotelitem.TaobaoXhotelHouseAddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
