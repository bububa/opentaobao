package nlife

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaNlifeB2cTradeDownloadAPIRequest
b2c下载订单 API请求
alibaba.nlife.b2c.trade.download

下载零售商在零售+平台创建的订单 */
type AlibabaNlifeB2cTradeDownloadAPIRequest struct {
	model.Params
	// 页码
	_pageNo int64
	// 分页大小
	_pageSize int64
	// 零售门店在零售+平台对应的ID
	_storeId string
	// 开始时间
	_startDate string
	// 结束时间
	_endDate string
}

// New
