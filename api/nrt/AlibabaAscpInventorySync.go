package nrt

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nrt"
)

/* AlibabaAscpInventorySync
库存同步接口
alibaba.ascp.inventory.sync

新零售联盟商家库存同步接口，用于商家商品库存数量同步到阿里系统 */
func AlibabaAscpInventorySync(clt *core.SDKClient, req *nrt.AlibabaAscpInventorySyncAPIRequest, session string) (*nrt.AlibabaAscpInventorySyncAPIResponse, error) {
	var resp nrt.AlibabaAscpInventorySyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
