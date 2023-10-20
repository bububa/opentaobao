package openmall

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenmallTradeBatchGetAPIRequest 批量获取openmall订单 API请求
// taobao.openmall.trade.batch.get
//
// 批量获取openmall订单
// 注意：该接口数据存在延迟，实时数据请通过taobao.openmall.trade.get获取
type TaobaoOpenmallTradeBatchGetAPIRequest struct {
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

// NewTaobaoOpenmallTradeBatchGetRequest 初始化TaobaoOpenmallTradeBatchGetAPIRequest对象
func NewTaobaoOpenmallTradeBatchGetRequest() *TaobaoOpenmallTradeBatchGetAPIRequest {
	return &TaobaoOpenmallTradeBatchGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOpenmallTradeBatchGetAPIRequest) GetApiMethodName() string {
	return "taobao.openmall.trade.batch.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOpenmallTradeBatchGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoOpenmallTradeBatchGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEndCreated is EndCreated Setter
// 查询范围结束时间，闭区间
func (r *TaobaoOpenmallTradeBatchGetAPIRequest) SetEndCreated(_endCreated string) error {
	r._endCreated = _endCreated
	r.Set("end_created", _endCreated)
	return nil
}

// GetEndCreated EndCreated Getter
func (r TaobaoOpenmallTradeBatchGetAPIRequest) GetEndCreated() string {
	return r._endCreated
}

// SetDistributor is Distributor Setter
// 渠道商Nick
func (r *TaobaoOpenmallTradeBatchGetAPIRequest) SetDistributor(_distributor string) error {
	r._distributor = _distributor
	r.Set("distributor", _distributor)
	return nil
}

// GetDistributor Distributor Getter
func (r TaobaoOpenmallTradeBatchGetAPIRequest) GetDistributor() string {
	return r._distributor
}

// SetStartCreated is StartCreated Setter
// 查询范围开始时间，闭区间
func (r *TaobaoOpenmallTradeBatchGetAPIRequest) SetStartCreated(_startCreated string) error {
	r._startCreated = _startCreated
	r.Set("start_created", _startCreated)
	return nil
}

// GetStartCreated StartCreated Getter
func (r TaobaoOpenmallTradeBatchGetAPIRequest) GetStartCreated() string {
	return r._startCreated
}

// SetPageIndex is PageIndex Setter
// 查询页码，从1开始
func (r *TaobaoOpenmallTradeBatchGetAPIRequest) SetPageIndex(_pageIndex int64) error {
	r._pageIndex = _pageIndex
	r.Set("page_index", _pageIndex)
	return nil
}

// GetPageIndex PageIndex Getter
func (r TaobaoOpenmallTradeBatchGetAPIRequest) GetPageIndex() int64 {
	return r._pageIndex
}

// SetPageSize is PageSize Setter
// 页面大小，不超过100
func (r *TaobaoOpenmallTradeBatchGetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoOpenmallTradeBatchGetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
