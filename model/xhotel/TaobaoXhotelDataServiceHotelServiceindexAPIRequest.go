package xhotel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoXhotelDataServiceHotelServiceindexAPIRequest
酒店服务指数 API请求
taobao.xhotel.data.service.hotel.serviceindex

酒店服务指数 */
type TaobaoXhotelDataServiceHotelServiceindexAPIRequest struct {
	model.Params
	// 酒店id
	_hid int64
	// 渠道商名称
	_vendor string
	// 1
	_startRow int64
	// 10
	_pageSize int64
	// 查询时间段结束
	_reportEndDate string
	// 查询时间段开始
	_reportStartDate string
	// 供应商名称
	_supplier string
}

// New
