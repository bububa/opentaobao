package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Taobaowdkequipmentconveyorconveyorinfoget 获取五道口悬挂链信息
// taobao.wdk.equipment.conveyor.conveyorinfo.get
//
// 获取五道口悬挂链信息
func Taobaowdkequipmentconveyorconveyorinfoget(clt *core.SDKClient, req *wdk.TaobaowdkequipmentconveyorconveyorinfogetAPIRequest, session string) (*wdk.TaobaowdkequipmentconveyorconveyorinfogetAPIResponse, error) {
	var resp wdk.TaobaowdkequipmentconveyorconveyorinfogetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
