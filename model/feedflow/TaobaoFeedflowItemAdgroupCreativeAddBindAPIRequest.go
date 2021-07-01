package feedflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFeedflowItemAdgroupCreativeAddBindAPIRequest
信息流新增并且绑定创意 API请求
taobao.feedflow.item.adgroup.creative.add.bind

信息流新增并且绑定创意 */
type TaobaoFeedflowItemAdgroupCreativeAddBindAPIRequest struct {
	model.Params
	// 新增绑定的创意，一次最多2个
	_creativeBindList []CreativeBindDto
	// 单元id
	_adgroupId int64
}

// New
