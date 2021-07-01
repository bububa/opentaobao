package lstpos

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lstpos"
)

/* AlibabaLstPosOpenInventorySyncinventorydata
商品库存修改同步接口(最多20条库存信息)
alibaba.lst.pos.open.inventory.syncinventorydata

商品库存修改同步接口(最多20条库存信息) */
func AlibabaLstPosOpenInventorySyncinventorydata(clt *core.SDKClient, req *lstpos.AlibabaLstPosOpenInventorySyncinventorydataAPIRequest, session string) (*lstpos.AlibabaLstPosOpenInventorySyncinventorydataAPIResponse, error) {
	var resp lstpos.AlibabaLstPosOpenInventorySyncinventorydataAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
