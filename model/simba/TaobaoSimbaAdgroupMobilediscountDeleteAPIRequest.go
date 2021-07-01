package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaAdgroupMobilediscountDeleteAPIRequest
批量删除adgroup的移动溢价 API请求
taobao.simba.adgroup.mobilediscount.delete

批量删除adgroup的移动溢价 */
type TaobaoSimbaAdgroupMobilediscountDeleteAPIRequest struct {
	model.Params
	// 昵称
	_nick string
	// adgroup主键数组（批量最多支持200个）
	_adgroupIds []int64
}

// New
