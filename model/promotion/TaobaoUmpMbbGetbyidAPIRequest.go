package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoUmpMbbGetbyidAPIRequest
获取营销积木块 API请求
taobao.ump.mbb.getbyid

根据积木块id获取营销积木块。 */
type TaobaoUmpMbbGetbyidAPIRequest struct {
	model.Params
	// 积木块的id
	_id int64
}

// New
