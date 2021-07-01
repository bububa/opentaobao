package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaKeywordFindbyidsAPIRequest
（新）根据一堆关键词ids获取关键词 API请求
taobao.simba.keyword.findbyids

根据一个关键词Id列表取得一组关键词 */
type TaobaoSimbaKeywordFindbyidsAPIRequest struct {
	model.Params
	// 关键词ids
	_bidwordIds []int64
}

// New
