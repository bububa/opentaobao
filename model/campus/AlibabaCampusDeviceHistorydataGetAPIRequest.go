package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCampusDeviceHistorydataGetAPIRequest
设备历史数据批量获取 API请求
alibaba.campus.device.historydata.get

设备历史数据批量获取 */
type AlibabaCampusDeviceHistorydataGetAPIRequest struct {
	model.Params
	// workbench
	_workBenchContext *WorkBenchContext
	// 查询条件
	_query *DeviceHistoryBatchQuery
}

// New
