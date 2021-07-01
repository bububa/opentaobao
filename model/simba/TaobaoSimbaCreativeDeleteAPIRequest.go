package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaCreativeDeleteAPIRequest
删除创意 API请求
taobao.simba.creative.delete

删除一个创意 */
type TaobaoSimbaCreativeDeleteAPIRequest struct {
	model.Params
	// 主人昵称
	_nick string
	// 创意Id
	_creativeId int64
}

// New
