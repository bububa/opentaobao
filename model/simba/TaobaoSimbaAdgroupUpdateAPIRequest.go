package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaAdgroupUpdateAPIRequest
更新一个推广组的信息 API请求
taobao.simba.adgroup.update

更新一个推广组的信息，可以设置默认出价、是否上线、非搜索出价、非搜索是否使用默认出价 */
type TaobaoSimbaAdgroupUpdateAPIRequest struct {
	model.Params
	// 非搜索是否使用默认出价，false-不用；true-使用；默认为true;
	_useNonsearchDefaultPrice string
	// 用户设置的上下线状态 offline-下线(暂停竞价)； online-上线；默认为online
	_onlineStatus string
	// 推广组Id
	_adgroupId int64
	// 默认出价，单位是分，不能小于5
	_defaultPrice int64
	// 非搜索出价，单位是分，不能小于5，如果use_nonseatch_default_price为使用默认出价，则此nonsearch_max_price字段传入的数据不起作用，商品将使用默认非搜索出价
	_nonsearchMaxPrice int64
	// 主人昵称
	_nick string
}

// New
