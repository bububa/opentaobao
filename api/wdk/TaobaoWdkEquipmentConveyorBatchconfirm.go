package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Taobaowdkequipmentconveyorbatchconfirm 五道口悬挂链信息批量确认
// taobao.wdk.equipment.conveyor.batchconfirm
//
// 批量消息确认
func Taobaowdkequipmentconveyorbatchconfirm(clt *core.SDKClient, req *wdk.TaobaowdkequipmentconveyorbatchconfirmAPIRequest, session string) (*wdk.TaobaowdkequipmentconveyorbatchconfirmAPIResponse, error) {
	var resp wdk.TaobaowdkequipmentconveyorbatchconfirmAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
