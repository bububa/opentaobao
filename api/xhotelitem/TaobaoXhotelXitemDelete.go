package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// TaobaoXhotelXitemDelete 删除 x 元素
// taobao.xhotel.xitem.delete
//
// 删除 x 元素
func TaobaoXhotelXitemDelete(clt *core.SDKClient, req *xhotelitem.TaobaoXhotelXitemDeleteAPIRequest, session string) (*xhotelitem.TaobaoXhotelXitemDeleteAPIResponse, error) {
	var resp xhotelitem.TaobaoXhotelXitemDeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
