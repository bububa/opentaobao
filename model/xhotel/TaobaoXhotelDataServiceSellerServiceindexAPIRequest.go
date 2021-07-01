package xhotel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoXhotelDataServiceSellerServiceindexAPIRequest
卖家服务指数查询 API请求
taobao.xhotel.data.service.seller.serviceindex

卖家服务指数查询 */
type TaobaoXhotelDataServiceSellerServiceindexAPIRequest struct {
	model.Params
	// 渠道商名称
	_vendor string
	// 分页参数
	_startRow int64
	// 分页参数
	_pageSize int64
	// 查询截止日期
	_reportEndDate string
	// 查询开始日期
	_reportStartDate string
}

// New
