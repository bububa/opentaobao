package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaKeywordsDeleteAPIRequest
删除一批关键词 API请求
taobao.simba.keywords.delete

删除一批关键词 */
type TaobaoSimbaKeywordsDeleteAPIRequest struct {
	model.Params
	// 主人昵称
	_nick string
	// 推广计划Id
	_campaignId int64
	// 关键词Id数组，最多100个
	_keywordIds []int64
}

// New
