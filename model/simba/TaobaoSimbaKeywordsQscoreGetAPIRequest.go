package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaKeywordsQscoreGetAPIRequest
取得一个推广组的所有关键词的质量得分或者根据关键词Id列表取得一组关键词的质量得分 API请求
taobao.simba.keywords.qscore.get

取得一个推广组的所有关键词的质量得分列表 */
type TaobaoSimbaKeywordsQscoreGetAPIRequest struct {
	model.Params
	// 主人昵称
	_nick string
	// 推广组Id
	_adgroupId int64
}

// New
