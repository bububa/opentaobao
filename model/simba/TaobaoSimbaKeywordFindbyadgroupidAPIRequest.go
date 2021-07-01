package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaKeywordFindbyadgroupidAPIRequest
根据推广单元id获取关键词 API请求
taobao.simba.keyword.findbyadgroupid

根据一个关键词Id列表取得一组关键词 */
type TaobaoSimbaKeywordFindbyadgroupidAPIRequest struct {
	model.Params
	// 推广单元id
	_adgroupId int64
}

// New
