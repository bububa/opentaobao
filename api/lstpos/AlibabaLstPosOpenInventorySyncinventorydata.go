package lstpos

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lstpos"
)

// Alibabalstposopeninventorysyncinventorydata 商品库存修改同步接口(最多20条库存信息)
// alibaba.lst.pos.open.inventory.syncinventorydata
//
// 商品库存修改同步接口(最多20条库存信息)
func Alibabalstposopeninventorysyncinventorydata(clt *core.SDKClient, req *lstpos.AlibabalstposopeninventorysyncinventorydataAPIRequest, session string) (*lstpos.AlibabalstposopeninventorysyncinventorydataAPIResponse, error) {
	var resp lstpos.AlibabalstposopeninventorysyncinventorydataAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
