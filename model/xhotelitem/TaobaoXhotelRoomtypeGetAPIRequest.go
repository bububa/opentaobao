package xhotelitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoXhotelRoomtypeGetAPIRequest
房型查询接口 API请求
taobao.xhotel.roomtype.get

房型查询房型查询接口返回结果增加date_confirm字段 */
type TaobaoXhotelRoomtypeGetAPIRequest struct {
	model.Params
	// 废弃，使用商家房型ID
	_rid int64
	// 商家房型ID
	_outerId string
	// 系统商，一般不填写，使用须申请
	_vendor string
}

// New
