package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoUmpMbbsListGetAPIRequest
通过ids列表获取营销积木块列表 API请求
taobao.ump.mbbs.list.get

通过营销积木id列表来获取营销积木块列表。 */
type TaobaoUmpMbbsListGetAPIRequest struct {
	model.Params
	// 营销积木块id组成的字符串。
	_ids []int64
}

// New
