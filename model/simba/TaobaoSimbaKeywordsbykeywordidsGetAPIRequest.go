package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaKeywordsbykeywordidsGetAPIRequest
根据一个关键词Id列表取得一组关键词 API请求
taobao.simba.keywordsbykeywordids.get

根据一个关键词Id列表取得一组关键词 */
type TaobaoSimbaKeywordsbykeywordidsGetAPIRequest struct {
	model.Params
	// 主人昵称
	_nick string
	// 关键词Id数组，最多200个；
	_keywordIds []int64
}

// New
