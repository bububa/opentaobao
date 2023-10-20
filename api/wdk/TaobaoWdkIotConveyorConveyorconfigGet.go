package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Taobaowdkiotconveyorconveyorconfigget 获取悬挂链基本配置信息
// taobao.wdk.iot.conveyor.conveyorconfig.get
//
// 用于从云端WCS获取悬挂链基本配置信息
func Taobaowdkiotconveyorconveyorconfigget(clt *core.SDKClient, req *wdk.TaobaowdkiotconveyorconveyorconfiggetAPIRequest, session string) (*wdk.TaobaowdkiotconveyorconveyorconfiggetAPIResponse, error) {
	var resp wdk.TaobaowdkiotconveyorconveyorconfiggetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
