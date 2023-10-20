package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Taobaowdkequipmentconveyorwcsbtoccontainerscannedbyconveyor 容器被悬挂链扫描
// taobao.wdk.equipment.conveyor.wcsbtoc.containerscannedbyconveyor
//
// 容器被悬挂链扫描
func Taobaowdkequipmentconveyorwcsbtoccontainerscannedbyconveyor(clt *core.SDKClient, req *wdk.TaobaowdkequipmentconveyorwcsbtoccontainerscannedbyconveyorAPIRequest, session string) (*wdk.TaobaowdkequipmentconveyorwcsbtoccontainerscannedbyconveyorAPIResponse, error) {
	var resp wdk.TaobaowdkequipmentconveyorwcsbtoccontainerscannedbyconveyorAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
