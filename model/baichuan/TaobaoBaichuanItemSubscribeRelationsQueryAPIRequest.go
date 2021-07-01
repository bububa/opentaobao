package baichuan

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoBaichuanItemSubscribeRelationsQueryAPIRequest
按条件查询订阅关系 API请求
taobao.baichuan.item.subscribe.relations.query

按条件查询订阅关系 */
type TaobaoBaichuanItemSubscribeRelationsQueryAPIRequest struct {
	model.Params
	// 查询条件
	_condition *Condition
}

// New
