package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaSalestarAdgroupUpdateAPIRequest
销量明星更新一个推广组的信息 API请求
taobao.simba.salestar.adgroup.update

更新一个推广组的信息，可以设置 是否上线 */
type TaobaoSimbaSalestarAdgroupUpdateAPIRequest struct {
	model.Params
	// 用户设置的上下线状态 offline-下线(暂停竞价)； online-上线；默认为online
	_onlineStatus string
	// 推广组Id
	_adgroupId int64
}

// New
