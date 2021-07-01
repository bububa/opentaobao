package xhotelitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoXhotelRoomtypeConflictDataAPIRequest
商家床型冲突数据接口 API请求
taobao.xhotel.roomtype.conflict.data

商家床型冲突数据接口 */
type TaobaoXhotelRoomtypeConflictDataAPIRequest struct {
	model.Params
	// 查询第几页数据
	_currentPage int64
}

// New
