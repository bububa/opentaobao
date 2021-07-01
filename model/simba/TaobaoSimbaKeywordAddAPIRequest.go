package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaKeywordAddAPIRequest
（新）关键词新增接口 API请求
taobao.simba.keyword.add

（新）关键词更新相关接口 */
type TaobaoSimbaKeywordAddAPIRequest struct {
	model.Params
	// 关键词相关信息
	_bidwords []SiriusBidwordDto
	// 推广单元id
	_adgroupId int64
}

// New
