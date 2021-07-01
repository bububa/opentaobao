package tttm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AliyunIndustryTttmStockSyncAPIRequest
天天特卖库存同步接口 API请求
aliyun.industry.tttm.stock.sync

天天特卖库存同步接口 */
type AliyunIndustryTttmStockSyncAPIRequest struct {
	model.Params
	// 库存
	_syncStock *StockInfoDto
}

// New
