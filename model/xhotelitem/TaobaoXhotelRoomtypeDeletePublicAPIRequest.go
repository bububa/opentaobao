package xhotelitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoXhotelRoomtypeDeletePublicAPIRequest
商家删除房型数据接口 API请求
taobao.xhotel.roomtype.delete.public

房型删除TOP接口 */
type TaobaoXhotelRoomtypeDeletePublicAPIRequest struct {
	model.Params
	// 房型rid ，传参方式：rid    或者   outer_id+vendor
	_rid int64
	// vendor
	_vendor string
	// 外部房型ID
	_outerRid string
	// 具体操作人，比如酒店帐号、小二名称等
	_operator string
}

// New
