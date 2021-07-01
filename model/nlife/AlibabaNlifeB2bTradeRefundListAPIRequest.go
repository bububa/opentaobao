package nlife

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaNlifeB2bTradeRefundListAPIRequest
获取采购退货单列表 API请求
alibaba.nlife.b2b.trade.refund.list

获取采购退货单列表 */
type AlibabaNlifeB2bTradeRefundListAPIRequest struct {
	model.Params
	// 采购退货单创建时间开始范围
	_startEffectiveDate string
	// 采购退货单创建时间结束范围
	_endEffectiveDate string
	// 查询的页数
	_pageNo int64
	// 每页的数量
	_pageSize int64
	// 企业Id
	_entId int64
}

// NewAlibabaNlifeB2bTradeRefundListRequest 初始化AlibabaNlifeB2bTradeRefundListAPIRequest对象
func NewAlibabaNlifeB2bTradeRefundListRequest() *AlibabaNlifeB2bTradeRefundListAPIRequest {
	return &AlibabaNlifeB2bTradeRefundListAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaNlifeB2bTradeRefundListAPIRequest) GetApiMethodName() string {
	return "alibaba.nlife.b2b.trade.refund.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaNlifeB2bTradeRefundListAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is StartEffectiveDate Setter
// 采购退货单创建时间开始范围
func (r *AlibabaNlifeB2bTradeRefundListAPIRequest) SetStartEffectiveDate(_startEffectiveDate string) error {
	r._startEffectiveDate = _startEffectiveDate
	r.Set("start_effective_date", _startEffectiveDate)
	return nil
}

// Get StartEffectiveDate Getter
func (r AlibabaNlifeB2bTradeRefundListAPIRequest) GetStartEffectiveDate() string {
	return r._startEffectiveDate
}

// Set is EndEffectiveDate Setter
// 采购退货单创建时间结束范围
func (r *AlibabaNlifeB2bTradeRefundListAPIRequest) SetEndEffectiveDate(_endEffectiveDate string) error {
	r._endEffectiveDate = _endEffectiveDate
	r.Set("end_effective_date", _endEffectiveDate)
	return nil
}

// Get EndEffectiveDate Getter
func (r AlibabaNlifeB2bTradeRefundListAPIRequest) GetEndEffectiveDate() string {
	return r._endEffectiveDate
}

// Set is PageNo Setter
// 查询的页数
func (r *AlibabaNlifeB2bTradeRefundListAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// Get PageNo Getter
func (r AlibabaNlifeB2bTradeRefundListAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// Set is PageSize Setter
// 每页的数量
func (r *AlibabaNlifeB2bTradeRefundListAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// Get PageSize Getter
func (r AlibabaNlifeB2bTradeRefundListAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// Set is EntId Setter
// 企业Id
func (r *AlibabaNlifeB2bTradeRefundListAPIRequest) SetEntId(_entId int64) error {
	r._entId = _entId
	r.Set("ent_id", _entId)
	return nil
}

// Get EntId Getter
func (r AlibabaNlifeB2bTradeRefundListAPIRequest) GetEntId() int64 {
	return r._entId
}
