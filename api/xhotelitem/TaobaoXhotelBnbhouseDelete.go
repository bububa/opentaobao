package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// Taobaoxhotelbnbhousedelete 民宿门店删除接口
// taobao.xhotel.bnbhouse.delete
//
// 支持门店相关的门店删除，删除门店会级联删除门店下面的房源
func Taobaoxhotelbnbhousedelete(clt *core.SDKClient, req *xhotelitem.TaobaoxhotelbnbhousedeleteAPIRequest, session string) (*xhotelitem.TaobaoxhotelbnbhousedeleteAPIResponse, error) {
	var resp xhotelitem.TaobaoxhotelbnbhousedeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
