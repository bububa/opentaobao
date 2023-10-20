package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// Taobaoxhoteldelete 删除酒店接口
// taobao.xhotel.delete
//
// 删除飞猪酒店数据接口
func Taobaoxhoteldelete(clt *core.SDKClient, req *xhotelitem.TaobaoxhoteldeleteAPIRequest, session string) (*xhotelitem.TaobaoxhoteldeleteAPIResponse, error) {
	var resp xhotelitem.TaobaoxhoteldeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
