package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaKeywordUpdateAPIRequest
（新）关键词更新相关接口 API请求
taobao.simba.keyword.update

（新）关键词更新相关接口 */
type TaobaoSimbaKeywordUpdateAPIRequest struct {
	model.Params
	// 关键词相关信息
	_bidwords []SiriusBidwordDto
}

// New
