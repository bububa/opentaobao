package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaAdgroupsbyadgroupidsGetAPIRequest
批量得到推广组 API请求
taobao.simba.adgroupsbyadgroupids.get

批量得到推广组 */
type TaobaoSimbaAdgroupsbyadgroupidsGetAPIRequest struct {
	model.Params
	// 主人昵称
	_nick string
	// 推广组Id列表
	_adgroupIds []int64
	// 页尺寸，最大200，如果入参adgroup_ids有传入值，则page_size和page_no值不起作用。如果adgrpup_ids为空而campaign_id有值，此时page_size和page_no值才是返回的页数据大小和页码
	_pageSize int64
	// 页码，从1开始
	_pageNo int64
}

// New
