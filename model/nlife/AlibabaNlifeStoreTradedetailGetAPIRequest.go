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

// NewAlibabaNlifeStoreTradedetailGetRequest 初始化AlibabaNlifeStoreTradedetailGetAPIRequest对象
func NewAlibabaNlifeStoreTradedetailGetRequest() *AlibabaNlifeStoreTradedetailGetAPIRequest {
	return &AlibabaNlifeStoreTradedetailGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaNlifeStoreTradedetailGetAPIRequest) GetApiMethodName() string {
	return "alibaba.nlife.store.tradedetail.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaNlifeStoreTradedetailGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ProcurementNo Setter
// 集团采购单号
func (r *AlibabaNlifeStoreTradedetailGetAPIRequest) SetProcurementNo(_procurementNo string) error {
	r._procurementNo = _procurementNo
	r.Set("procurement_no", _procurementNo)
	return nil
}

// Get ProcurementNo Getter
func (r AlibabaNlifeStoreTradedetailGetAPIRequest) GetProcurementNo() string {
	return r._procurementNo
}
