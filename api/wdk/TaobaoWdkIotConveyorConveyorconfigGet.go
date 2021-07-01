package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

/* TaobaoWdkIotConveyorConveyorconfigGet
获取悬挂链基本配置信息
taobao.wdk.iot.conveyor.conveyorconfig.get

用于从云端WCS获取悬挂链基本配置信息 */
func TaobaoWdkIotConveyorConveyorconfigGet(clt *core.SDKClient, req *wdk.TaobaoWdkIotConveyorConveyorconfigGetAPIRequest, session string) (*wdk.TaobaoWdkIotConveyorConveyorconfigGetAPIResponse, error) {
	var resp wdk.TaobaoWdkIotConveyorConveyorconfigGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
