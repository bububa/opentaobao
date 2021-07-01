package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihouseNewhomeBusinessSyncAPIRequest
商圈数据同步 API请求
alibaba.alihouse.newhome.business.sync

商圈数据同步 */
type AlibabaAlihouseNewhomeBusinessSyncAPIRequest struct {
	model.Params
	// 入参数据
	_baseBusinessDto *BaseBusinessDto
}

// New
