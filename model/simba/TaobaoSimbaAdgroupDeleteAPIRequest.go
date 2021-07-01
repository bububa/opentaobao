package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaAdgroupDeleteAPIRequest
删除一个推广组 API请求
taobao.simba.adgroup.delete

删除一个推广组 */
type TaobaoSimbaAdgroupDeleteAPIRequest struct {
	model.Params
	// 主人昵称
	_nick string
	// 推广组Id
	_adgroupId int64
}

// New
