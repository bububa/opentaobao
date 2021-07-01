package xhotelitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoXhotelServicetimeGetAPIRequest
查询实体对应的服务时间数据 API请求
taobao.xhotel.servicetime.get

通过实体来获取对应的服务时间数据 */
type TaobaoXhotelServicetimeGetAPIRequest struct {
	model.Params
	// 酒店id
	_hid int64
}

// New
