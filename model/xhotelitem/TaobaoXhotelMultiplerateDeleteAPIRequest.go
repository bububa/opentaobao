package xhotelitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoXhotelMultiplerateDeleteAPIRequest
复杂价格删除 API请求
taobao.xhotel.multiplerate.delete

酒店产品库rate删除 */
type TaobaoXhotelMultiplerateDeleteAPIRequest struct {
	model.Params
	// 渠道，和推送房价所使用的渠道保持一致
	_vendor string
	// 商家价格政策编码
	_rateplanCode string
	// 商家房型编码
	_outRid string
	// 连住天数
	_occupancy int64
	// 入住人数
	_lengthofstay int64
}

// New
