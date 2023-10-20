package openmall

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoopenmalltradebatchgetAPIRequest 批量获取openmall订单 API请求
// taobao.openmall.trade.batch.get
//
// 批量获取openmall订单
// 注意：该接口数据存在延迟，实时数据请通过taobao.openmall.trade.get获取
type TaobaoopenmalltradebatchgetAPIRequest struct {
	model.Params
	// 查询范围结束时间，闭区间
	_endCreated string
	// 渠道商Nick
	_distributor string
	// 查询范围开始时间，闭区间
	_startCreated string
	// 查询页码，从1开始
	_pageIndex int64
	// 页面大小，不超过100
	_pageSize int64
}

// NewTaobaoopenmalltradebatchgetRequest 初始化TaobaoopenmalltradebatchgetAPIRequest对象
func NewTaobaoopenmalltradebatchgetRequest() *TaobaoopenmalltradebatchgetAPIRequest {
	return &TaobaoopenmalltradebatchgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoopenmalltradebatchgetAPIRequest) GetApiMethodName() string {
	return "taobao.openmall.trade.batch.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoopenmalltradebatchgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoopenmalltradebatchgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEndCreated is EndCreated Setter
// 查询范围结束时间，闭区间
func (r *TaobaoopenmalltradebatchgetAPIRequest) SetEndCreated(_endCreated string) error {
	r._endCreated = _endCreated
	r.Set("end_created", _endCreated)
	return nil
}

// GetEndCreated EndCreated Getter
func (r TaobaoopenmalltradebatchgetAPIRequest) GetEndCreated() string {
	return r._endCreated
}

// SetDistributor is Distributor Setter
// 渠道商Nick
func (r *TaobaoopenmalltradebatchgetAPIRequest) SetDistributor(_distributor string) error {
	r._distributor = _distributor
	r.Set("distributor", _distributor)
	return nil
}

// GetDistributor Distributor Getter
func (r TaobaoopenmalltradebatchgetAPIRequest) GetDistributor() string {
	return r._distributor
}

// SetStartCreated is StartCreated Setter
// 查询范围开始时间，闭区间
func (r *TaobaoopenmalltradebatchgetAPIRequest) SetStartCreated(_startCreated string) error {
	r._startCreated = _startCreated
	r.Set("start_created", _startCreated)
	return nil
}

// GetStartCreated StartCreated Getter
func (r TaobaoopenmalltradebatchgetAPIRequest) GetStartCreated() string {
	return r._startCreated
}

// SetPageIndex is PageIndex Setter
// 查询页码，从1开始
func (r *TaobaoopenmalltradebatchgetAPIRequest) SetPageIndex(_pageIndex int64) error {
	r._pageIndex = _pageIndex
	r.Set("page_index", _pageIndex)
	return nil
}

// GetPageIndex PageIndex Getter
func (r TaobaoopenmalltradebatchgetAPIRequest) GetPageIndex() int64 {
	return r._pageIndex
}

// SetPageSize is PageSize Setter
// 页面大小，不超过100
func (r *TaobaoopenmalltradebatchgetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoopenmalltradebatchgetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
