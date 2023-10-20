package lstpos

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lstpos"
)

// AlibabaLstPosOpenInventorySyncinventorydata 商品库存修改同步接口(最多20条库存信息)
// alibaba.lst.pos.open.inventory.syncinventorydata
//
// 商品库存修改同步接口(最多20条库存信息)
func AlibabaLstPosOpenInventorySyncinventorydata(clt *core.SDKClient, req *lstpos.AlibabaLstPosOpenInventorySyncinventorydataAPIRequest, resp *lstpos.AlibabaLstPosOpenInventorySyncinventorydataAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
