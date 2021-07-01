package tmallcar

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallAliautoMetaReceiveAPIRequest
汽车说明书元数据上传 API请求
tmall.aliauto.meta.receive

天猫汽车对外提供的汽车资源元数据上传接口 */
type TmallAliautoMetaReceiveAPIRequest struct {
	model.Params
	// 元数据参数集
	_command *ResourceMetaCommand
}

// New
