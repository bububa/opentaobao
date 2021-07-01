package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoUmpMbbsGetAPIRequest
获取营销积木块列表 API请求
taobao.ump.mbbs.get

获取营销积木块列表，可以根据类型获取，也可以将该字段设为空，获取所有的 */
type TaobaoUmpMbbsGetAPIRequest struct {
	model.Params
	// 积木块类型。如果该字段为空表示查出所有类型的<br/>现在有且仅有如下几种：resource,condition,action,target
	_type string
}

// New
