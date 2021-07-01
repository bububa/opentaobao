package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaSerchcrowdBatchDeleteAPIRequest
单品搜索人群批量取消溢价 API请求
taobao.simba.serchcrowd.batch.delete

删除单品搜索人群溢价功能 */
type TaobaoSimbaSerchcrowdBatchDeleteAPIRequest struct {
	model.Params
	// 被操作者的淘宝昵称
	_nick string
	// 子帐号nick
	_subNick string
	// 需要删除的人群id
	_adgroupCrowdIds []int64
}

// New
