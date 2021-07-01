package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCampusGuardDataSyncAPIRequest
卡巴数据同步 API请求
alibaba.campus.guard.data.sync

数据同步门禁系统 */
type AlibabaCampusGuardDataSyncAPIRequest struct {
	model.Params
	// 1-刷卡流水
	_dataType string
	// 供应商名称
	_supplierName string
	// json串
	_data string
}

// New
