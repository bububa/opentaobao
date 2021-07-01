package xhotelitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoXhotelBaseinfoRoomGetAPIRequest
酒店房型与房价查询 API请求
taobao.xhotel.baseinfo.room.get

根据outHid/hid获取酒店的房型和价格信息 */
type TaobaoXhotelBaseinfoRoomGetAPIRequest struct {
	model.Params
	// 卖家系统中的酒店ID。
	_outHid string
	// 用于标示该酒店发布的渠道信息
	_vendor string
	// 是否需要房价基本信息（false为不需要），默认为需要
	_isNeedRatePlan bool
}

// New
