package nlife

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaNlifeStoreTradedetailGetAPIRequest
查询采购单详情信息 API请求
alibaba.nlife.store.tradedetail.get

根据集团id和采购单号，查询采购单的详情信息 */
type AlibabaNlifeStoreTradedetailGetAPIRequest struct {
	model.Params
	// 集团采购单号
	_procurementNo string
}

// New
