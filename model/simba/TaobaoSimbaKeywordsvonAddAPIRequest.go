package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaKeywordsvonAddAPIRequest
创建一批关键词 API请求
taobao.simba.keywordsvon.add

创建一批关键词 */
type TaobaoSimbaKeywordsvonAddAPIRequest struct {
	model.Params
	// 主人昵称
	_nick string
	// 推广组id
	_adgroupId int64
	// 关键词、计算机出价、无线出价和匹配方式json字符串，word:词，不能有一些特殊字符。maxPrice：计算机出价，是整数，以“分”为单位，不能小于5，不能大于日限额, maxMobilePrice：代表无线出价，规则同maxPice 当matchscope只能是1,4（1代表精确匹配，4代表广泛匹配）。
	_keywordPrices string
}

// New
