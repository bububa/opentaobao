package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihouseNewhomeRegionSyncAPIRequest
城区数据同步 API请求
alibaba.alihouse.newhome.region.sync

城区数据同步 */
type AlibabaAlihouseNewhomeRegionSyncAPIRequest struct {
	model.Params
	// 城区数据
	_baseRegionDto *BaseRegionDto
}

// New
