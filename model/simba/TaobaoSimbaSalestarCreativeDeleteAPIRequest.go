package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaSalestarCreativeDeleteAPIRequest
(新)销量明星删除创意相关接口 API请求
taobao.simba.salestar.creative.delete

删除一个创意 */
type TaobaoSimbaSalestarCreativeDeleteAPIRequest struct {
	model.Params
	// 创意Id
	_creativeId int64
}

// New
