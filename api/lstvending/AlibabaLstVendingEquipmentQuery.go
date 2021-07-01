package lstvending

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lstvending"
)

/* AlibabaLstVendingEquipmentQuery
自动售卖机设备信息查询
alibaba.lst.vending.equipment.query

为零售通品牌商提供已租赁的设备信息查询。 */
func AlibabaLstVendingEquipmentQuery(clt *core.SDKClient, req *lstvending.AlibabaLstVendingEquipmentQueryAPIRequest, session string) (*lstvending.AlibabaLstVendingEquipmentQueryAPIResponse, error) {
	var resp lstvending.AlibabaLstVendingEquipmentQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
