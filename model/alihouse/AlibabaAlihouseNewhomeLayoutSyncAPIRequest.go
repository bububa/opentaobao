package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihouseNewhomeLayoutSyncAPIRequest
房通户型数据同步 API请求
alibaba.alihouse.newhome.layout.sync

房通户型数据同步 */
type AlibabaAlihouseNewhomeLayoutSyncAPIRequest struct {
	model.Params
	// 数据
	_syncProjectLayoutData *SyncProjectLayoutDto
}

// New
