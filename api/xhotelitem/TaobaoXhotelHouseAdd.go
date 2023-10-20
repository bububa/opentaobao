package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// Taobaoxhotelhouseadd 非标准民宿房源添加
// taobao.xhotel.house.add
//
// 添加酒店或更新酒店
func Taobaoxhotelhouseadd(clt *core.SDKClient, req *xhotelitem.TaobaoxhotelhouseaddAPIRequest, session string) (*xhotelitem.TaobaoxhotelhouseaddAPIResponse, error) {
	var resp xhotelitem.TaobaoxhotelhouseaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
