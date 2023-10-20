package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// Taobaoxhotelrateplandelete 价格计划rateplan删除
// taobao.xhotel.rateplan.delete
//
// 酒店产品库rateplan删除，同时删除级联的rate，请谨慎使用
func Taobaoxhotelrateplandelete(clt *core.SDKClient, req *xhotelitem.TaobaoxhotelrateplandeleteAPIRequest, session string) (*xhotelitem.TaobaoxhotelrateplandeleteAPIResponse, error) {
	var resp xhotelitem.TaobaoxhotelrateplandeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
