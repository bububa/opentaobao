package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// Taobaoxhotelincrementinfoget 酒店状态增量查询接口
// taobao.xhotel.increment.info.get
//
// 酒店状态增量查询接口
func Taobaoxhotelincrementinfoget(clt *core.SDKClient, req *xhotelitem.TaobaoxhotelincrementinfogetAPIRequest, session string) (*xhotelitem.TaobaoxhotelincrementinfogetAPIResponse, error) {
	var resp xhotelitem.TaobaoxhotelincrementinfogetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
