package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

/* TaobaoWdkEquipmentConveyorStatusconfirm
悬挂链状态回传确认
taobao.wdk.equipment.conveyor.statusconfirm

悬挂链状态回传确认 */
func TaobaoWdkEquipmentConveyorStatusconfirm(clt *core.SDKClient, req *wdk.TaobaoWdkEquipmentConveyorStatusconfirmAPIRequest, session string) (*wdk.TaobaoWdkEquipmentConveyorStatusconfirmAPIResponse, error) {
	var resp wdk.TaobaoWdkEquipmentConveyorStatusconfirmAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
