package lstvending

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lstvending"
)

// AlibabaLstVendingEquipmentQuery 自动售卖机设备信息查询
// alibaba.lst.vending.equipment.query
//
// 为零售通品牌商提供已租赁的设备信息查询。
func AlibabaLstVendingEquipmentQuery(ctx context.Context, clt *core.SDKClient, req *lstvending.AlibabaLstVendingEquipmentQueryAPIRequest, resp *lstvending.AlibabaLstVendingEquipmentQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
