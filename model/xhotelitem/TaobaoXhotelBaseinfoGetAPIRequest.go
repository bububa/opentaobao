package xhotelitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoXhotelBaseinfoGetAPIRequest
酒店基础信息查询接口 API请求
taobao.xhotel.baseinfo.get

酒店基础信息(酒店/房型/房价定义)查询接口， 包括 酒店房型可售, 以及 hid 下 的标准房型列表 */
type TaobaoXhotelBaseinfoGetAPIRequest struct {
	model.Params
	// 淘宝酒店ID
	_hid int64
	// 推荐使用卖家系统中的酒店ID。
	_outHid string
	// 用于标示该酒店发布的渠道信息
	_vendor string
	// 是否需要房价基本信息（false为不需要），默认为需要
	_isNeedRatePlan bool
	// 是否需要房型基本信息（false为不需要），默认为需要
	_isNeedRoomType bool
	// 是否需要 根据 hid 查询 标准房型列表
	_needSRoomTypeList bool
	// 是否需要酒店房型可售详情
	_needHotelDynamicInfo bool
	// 在查询酒店房型可售详情 时的入参JSON , {@link com.taobao.trip.hpc.client.query.HotelSellerInvQuery}
	_jsonHotelSellerInvQuery string
}

// New
