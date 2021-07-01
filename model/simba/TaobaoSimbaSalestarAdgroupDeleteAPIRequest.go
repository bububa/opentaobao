package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaSalestarAdgroupDeleteAPIRequest
(新)销量明星删除推广单元接口 API请求
taobao.simba.salestar.adgroup.delete

删除一个推广组 */
type TaobaoSimbaSalestarAdgroupDeleteAPIRequest struct {
	model.Params
	// 主人昵称
	_nick string
	// 推广组Id
	_adgroupId int64
}

// New
