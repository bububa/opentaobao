package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaSalestarAdgroupFindbycampidAPIRequest
(销量明星)批量获取推广计划下的推广组信息 API请求
taobao.simba.salestar.adgroup.findbycampid

批量得到推广计划下的推广组 */
type TaobaoSimbaSalestarAdgroupFindbycampidAPIRequest struct {
	model.Params
	// 推广计划Id
	_campaignId int64
	// 页尺寸，最大200，如果入参adgroup_ids有传入值，则page_size和page_no值不起作用。如果adgrpup_ids为空而campaign_id有值，此时page_size和page_no值才是返回的页数据大小和页码
	_pageSize int64
	// 页码，从1开始
	_pageNo int64
}

// New
