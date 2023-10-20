package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Taobaowdkequipmentconveyorwcsbtoccontainerassignedtoconveyor 容器被预分拣器分配到悬挂链
// taobao.wdk.equipment.conveyor.wcsbtoc.containerassignedtoconveyor
//
// 容器被预分拣器分配到悬挂链
func Taobaowdkequipmentconveyorwcsbtoccontainerassignedtoconveyor(clt *core.SDKClient, req *wdk.TaobaowdkequipmentconveyorwcsbtoccontainerassignedtoconveyorAPIRequest, session string) (*wdk.TaobaowdkequipmentconveyorwcsbtoccontainerassignedtoconveyorAPIResponse, error) {
	var resp wdk.TaobaowdkequipmentconveyorwcsbtoccontainerassignedtoconveyorAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
