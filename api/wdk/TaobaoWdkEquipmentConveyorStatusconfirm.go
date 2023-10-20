package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Taobaowdkequipmentconveyorstatusconfirm 悬挂链状态回传确认
// taobao.wdk.equipment.conveyor.statusconfirm
//
// 悬挂链状态回传确认
func Taobaowdkequipmentconveyorstatusconfirm(clt *core.SDKClient, req *wdk.TaobaowdkequipmentconveyorstatusconfirmAPIRequest, session string) (*wdk.TaobaowdkequipmentconveyorstatusconfirmAPIResponse, error) {
	var resp wdk.TaobaowdkequipmentconveyorstatusconfirmAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
