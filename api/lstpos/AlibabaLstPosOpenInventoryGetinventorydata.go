package lstpos

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lstpos"
)

// Alibabalstposopeninventorygetinventorydata 商品库存只读接口(最多20条库存信息)
// alibaba.lst.pos.open.inventory.getinventorydata
//
// 商品库存只读接口(最多20条库存信息)
func Alibabalstposopeninventorygetinventorydata(clt *core.SDKClient, req *lstpos.AlibabalstposopeninventorygetinventorydataAPIRequest, session string) (*lstpos.AlibabalstposopeninventorygetinventorydataAPIResponse, error) {
	var resp lstpos.AlibabalstposopeninventorygetinventorydataAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
