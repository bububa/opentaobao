package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaSalestarKeywordsDeleteAPIRequest
销量明星关键词删除 API请求
taobao.simba.salestar.keywords.delete

（新）关键词删除相关接口 */
type TaobaoSimbaSalestarKeywordsDeleteAPIRequest struct {
	model.Params
	// 关键词ids
	_bidwordIds []int64
}

// New
