package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// Taobaoxhotelbnbroomtypedelete 民宿房源删除接口
// taobao.xhotel.bnbroomtype.delete
//
// 增加民宿房源删除接口
func Taobaoxhotelbnbroomtypedelete(clt *core.SDKClient, req *xhotelitem.TaobaoxhotelbnbroomtypedeleteAPIRequest, session string) (*xhotelitem.TaobaoxhotelbnbroomtypedeleteAPIResponse, error) {
	var resp xhotelitem.TaobaoxhotelbnbroomtypedeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
