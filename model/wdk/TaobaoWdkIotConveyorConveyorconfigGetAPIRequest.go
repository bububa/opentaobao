package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWdkIotConveyorConveyorconfigGetAPIRequest
获取悬挂链基本配置信息 API请求
taobao.wdk.iot.conveyor.conveyorconfig.get

用于从云端WCS获取悬挂链基本配置信息 */
type TaobaoWdkIotConveyorConveyorconfigGetAPIRequest struct {
	model.Params
	// 仓编码
	_warehouseCode string
	// 悬挂链id，默认为1
	_conveyorId int64
}

// New
