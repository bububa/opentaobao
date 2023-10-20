package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// TaobaoXhotelIncrementInfoGet 酒店状态增量查询接口
// taobao.xhotel.increment.info.get
//
// 酒店状态增量查询接口
func TaobaoXhotelIncrementInfoGet(clt *core.SDKClient, req *xhotelitem.TaobaoXhotelIncrementInfoGetAPIRequest, session string) (*xhotelitem.TaobaoXhotelIncrementInfoGetAPIResponse, error) {
	var resp xhotelitem.TaobaoXhotelIncrementInfoGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
