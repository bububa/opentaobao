package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaKeywordsQscoreSplitGetAPIRequest
新质量分服务 API请求
taobao.simba.keywords.qscore.split.get

获取关键词新的质量分 */
type TaobaoSimbaKeywordsQscoreSplitGetAPIRequest struct {
	model.Params
	// 账号昵称
	_nick string
	// 推广组id
	_adGroupId int64
	// 词id数组（最多批量获取20个）
	_bidwordIds []int64
}

// New
