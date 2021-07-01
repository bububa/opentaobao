package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaAdgroupMobilediscountUpdateAPIRequest
对推广组进行单独移动溢价 API请求
taobao.simba.adgroup.mobilediscount.update

对推广组进行单独移动溢价 */
type TaobaoSimbaAdgroupMobilediscountUpdateAPIRequest struct {
	model.Params
	// 推广组id数组(推广组id集合元素个数在1-200个之间，推广组id需要在同一个推广计划中)
	_adgroupIds []int64
	// 折扣（折扣值在1-400之间）
	_mobileDiscount int64
	// 昵称
	_nick string
}

// New
