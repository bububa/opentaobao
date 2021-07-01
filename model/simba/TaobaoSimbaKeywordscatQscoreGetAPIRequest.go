package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaKeywordscatQscoreGetAPIRequest
取得一个推广组的所有关键词和类目出价的质量得分 API请求
taobao.simba.keywordscat.qscore.get

取得一个推广组的所有关键词和类目出价的质量得分列表 */
type TaobaoSimbaKeywordscatQscoreGetAPIRequest struct {
	model.Params
	// 主人昵称
	_nick string
	// 推广组Id
	_adgroupId int64
}

// New
