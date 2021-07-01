package filmtfavatar

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFilmTfavatarBillSalePrintQueryAPIRequest
获取影院卖品账单-核销账单 API请求
taobao.film.tfavatar.bill.sale.print.query

获取影院卖品账单-核销账单
返回值data属于加密字段, 并非大字段. */
type TaobaoFilmTfavatarBillSalePrintQueryAPIRequest struct {
	model.Params
	// 自运营开放平台APPKEY
	_openAppKey string
	// 影院ID
	_cinemaId int64
	// 开始时间
	_beginTime string
	// 结束时间
	_endTime string
	// 包含的订单状态, 默认不填
	_includedOrderStatus []string
	// offset 下标, 从0开始
	_offset int64
	// 页大小
	_pageSize int64
}

// New
