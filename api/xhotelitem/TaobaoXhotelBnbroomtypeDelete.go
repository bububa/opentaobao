package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// TaobaoXhotelBnbroomtypeDelete 民宿房源删除接口
// taobao.xhotel.bnbroomtype.delete
//
// 增加民宿房源删除接口
func TaobaoXhotelBnbroomtypeDelete(clt *core.SDKClient, req *xhotelitem.TaobaoXhotelBnbroomtypeDeleteAPIRequest, session string) (*xhotelitem.TaobaoXhotelBnbroomtypeDeleteAPIResponse, error) {
	var resp xhotelitem.TaobaoXhotelBnbroomtypeDeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
