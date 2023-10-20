package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Taobaowdkequipmentconveyorsystemeventget 获取悬挂链系统事件
// taobao.wdk.equipment.conveyor.systemevent.get
//
// 五道口悬挂链系统事件查询
func Taobaowdkequipmentconveyorsystemeventget(clt *core.SDKClient, req *wdk.TaobaowdkequipmentconveyorsystemeventgetAPIRequest, session string) (*wdk.TaobaowdkequipmentconveyorsystemeventgetAPIResponse, error) {
	var resp wdk.TaobaowdkequipmentconveyorsystemeventgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
