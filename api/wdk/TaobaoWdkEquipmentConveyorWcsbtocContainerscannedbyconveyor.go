package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

/* TaobaoWdkEquipmentConveyorWcsbtocContainerscannedbyconveyor
容器被悬挂链扫描
taobao.wdk.equipment.conveyor.wcsbtoc.containerscannedbyconveyor

容器被悬挂链扫描 */
func TaobaoWdkEquipmentConveyorWcsbtocContainerscannedbyconveyor(clt *core.SDKClient, req *wdk.TaobaoWdkEquipmentConveyorWcsbtocContainerscannedbyconveyorAPIRequest, session string) (*wdk.TaobaoWdkEquipmentConveyorWcsbtocContainerscannedbyconveyorAPIResponse, error) {
	var resp wdk.TaobaoWdkEquipmentConveyorWcsbtocContainerscannedbyconveyorAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
