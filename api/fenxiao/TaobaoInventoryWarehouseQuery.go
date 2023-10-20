package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// Taobaoinventorywarehousequery 分页查询商家仓信息
// taobao.inventory.warehouse.query
//
// 分页查询商家仓信息
func Taobaoinventorywarehousequery(clt *core.SDKClient, req *fenxiao.TaobaoinventorywarehousequeryAPIRequest, session string) (*fenxiao.TaobaoinventorywarehousequeryAPIResponse, error) {
	var resp fenxiao.TaobaoinventorywarehousequeryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
