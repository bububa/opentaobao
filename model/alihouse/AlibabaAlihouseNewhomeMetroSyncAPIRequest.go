package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihouseNewhomeMetroSyncAPIRequest
地铁数据同步 API请求
alibaba.alihouse.newhome.metro.sync

地铁数据同步 */
type AlibabaAlihouseNewhomeMetroSyncAPIRequest struct {
	model.Params
	// 地铁入参数据
	_baseMetroLineDto *BaseMetroLineDto
}

// New
