package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaKeywordsbyadgroupidGetAPIRequest
取得一个推广组的所有关键词 API请求
taobao.simba.keywordsbyadgroupid.get

取得一个推广组的所有关键词 */
type TaobaoSimbaKeywordsbyadgroupidGetAPIRequest struct {
	model.Params
	// 主人昵称
	_nick string
	// 推广组Id
	_adgroupId int64
}

// New
