package lstpos

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lstpos"
)

// AlibabaLstPosOpenInventoryGetinventorydata 商品库存只读接口(最多20条库存信息)
// alibaba.lst.pos.open.inventory.getinventorydata
//
// 商品库存只读接口(最多20条库存信息)
func AlibabaLstPosOpenInventoryGetinventorydata(clt *core.SDKClient, req *lstpos.AlibabaLstPosOpenInventoryGetinventorydataAPIRequest, resp *lstpos.AlibabaLstPosOpenInventoryGetinventorydataAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
