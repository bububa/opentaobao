package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihouseNewhomeRcSyncAPIRequest
阿里房产图文草稿信息同步 API请求
alibaba.alihouse.newhome.rc.sync

接收图文草稿信息 */
type AlibabaAlihouseNewhomeRcSyncAPIRequest struct {
	model.Params
	// 图文内容
	_rc *RichContentDraftDto
}

// New
