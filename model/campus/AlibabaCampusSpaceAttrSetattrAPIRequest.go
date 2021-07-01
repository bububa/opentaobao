package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCampusSpaceAttrSetattrAPIRequest
新增业务属性实例接口 API请求
alibaba.campus.space.attr.setattr

新增业务属性实例接口 */
type AlibabaCampusSpaceAttrSetattrAPIRequest struct {
	model.Params
	// 操作用户上下文
	_context *WorkBenchContext
	// 业务属性实例集合
	_list []TypeAttrInstanceRequest
}

// New
