package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Taobaowdkequipmentconveyorcontainerinfoget 获取批次或波次中容器的信息
// taobao.wdk.equipment.conveyor.containerinfo.get
//
// 获取批次或波次中容器的信息
func Taobaowdkequipmentconveyorcontainerinfoget(clt *core.SDKClient, req *wdk.TaobaowdkequipmentconveyorcontainerinfogetAPIRequest, session string) (*wdk.TaobaowdkequipmentconveyorcontainerinfogetAPIResponse, error) {
	var resp wdk.TaobaowdkequipmentconveyorcontainerinfogetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
