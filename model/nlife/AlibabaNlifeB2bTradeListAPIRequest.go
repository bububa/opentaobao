package nlife

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaNlifeB2bTradeListAPIRequest
获取企业下的采购单列表 API请求
alibaba.nlife.b2b.trade.list

获取指定门店下的采购单列表 */
type AlibabaNlifeB2bTradeListAPIRequest struct {
	model.Params
	// 采购单生效时间开始范围
	_startEffectiveDate string
	// 采购单生效时间结束范围
	_endEffectiveDate string
	// 查询的页码
	_pageNo int64
	// 每页的数量
	_pageSize int64
	// 企业ID
	_entId int64
}

// NewAlibabaNlifeB2bTradeListRequest 初始化AlibabaNlifeB2bTradeListAPIRequest对象
func NewAlibabaNlifeB2bTradeListRequest() *AlibabaNlifeB2bTradeListAPIRequest {
	return &AlibabaNlifeB2bTradeListAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaNlifeB2bTradeListAPIRequest) GetApiMethodName() string {
	return "alibaba.nlife.b2b.trade.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaNlifeB2bTradeListAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is StartEffectiveDate Setter
// 采购单生效时间开始范围
func (r *AlibabaNlifeB2bTradeListAPIRequest) SetStartEffectiveDate(_startEffectiveDate string) error {
	r._startEffectiveDate = _startEffectiveDate
	r.Set("start_effective_date", _startEffectiveDate)
	return nil
}

// Get StartEffectiveDate Getter
func (r AlibabaNlifeB2bTradeListAPIRequest) GetStartEffectiveDate() string {
	return r._startEffectiveDate
}

// Set is EndEffectiveDate Setter
// 采购单生效时间结束范围
func (r *AlibabaNlifeB2bTradeListAPIRequest) SetEndEffectiveDate(_endEffectiveDate string) error {
	r._endEffectiveDate = _endEffectiveDate
	r.Set("end_effective_date", _endEffectiveDate)
	return nil
}

// Get EndEffectiveDate Getter
func (r AlibabaNlifeB2bTradeListAPIRequest) GetEndEffectiveDate() string {
	return r._endEffectiveDate
}

// Set is PageNo Setter
// 查询的页码
func (r *AlibabaNlifeB2bTradeListAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// Get PageNo Getter
func (r AlibabaNlifeB2bTradeListAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// Set is PageSize Setter
// 每页的数量
func (r *AlibabaNlifeB2bTradeListAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// Get PageSize Getter
func (r AlibabaNlifeB2bTradeListAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// Set is EntId Setter
// 企业ID
func (r *AlibabaNlifeB2bTradeListAPIRequest) SetEntId(_entId int64) error {
	r._entId = _entId
	r.Set("ent_id", _entId)
	return nil
}

// Get EntId Getter
func (r AlibabaNlifeB2bTradeListAPIRequest) GetEntId() int64 {
	return r._entId
}
