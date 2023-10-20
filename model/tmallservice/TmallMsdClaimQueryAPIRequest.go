package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallmsdclaimqueryAPIRequest 查询待理赔工单数据接口 API请求
// tmall.msd.claim.query
//
// 查询待理赔工单数据接口
type TmallmsdclaimqueryAPIRequest struct {
	model.Params
	// 工单id
	_workcardId int64
	// 淘宝交易订单号，实物子订单
	_bizOrderId int64
	// 当前页数
	_pageIndex int64
	// 每页大小
	_pageSize int64
	// 交易服务订单号
	_serviceOrderId int64
}

// NewTmallmsdclaimqueryRequest 初始化TmallmsdclaimqueryAPIRequest对象
func NewTmallmsdclaimqueryRequest() *TmallmsdclaimqueryAPIRequest {
	return &TmallmsdclaimqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallmsdclaimqueryAPIRequest) GetApiMethodName() string {
	return "tmall.msd.claim.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallmsdclaimqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallmsdclaimqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWorkcardId is WorkcardId Setter
// 工单id
func (r *TmallmsdclaimqueryAPIRequest) SetWorkcardId(_workcardId int64) error {
	r._workcardId = _workcardId
	r.Set("workcard_id", _workcardId)
	return nil
}

// GetWorkcardId WorkcardId Getter
func (r TmallmsdclaimqueryAPIRequest) GetWorkcardId() int64 {
	return r._workcardId
}

// SetBizOrderId is BizOrderId Setter
// 淘宝交易订单号，实物子订单
func (r *TmallmsdclaimqueryAPIRequest) SetBizOrderId(_bizOrderId int64) error {
	r._bizOrderId = _bizOrderId
	r.Set("biz_order_id", _bizOrderId)
	return nil
}

// GetBizOrderId BizOrderId Getter
func (r TmallmsdclaimqueryAPIRequest) GetBizOrderId() int64 {
	return r._bizOrderId
}

// SetPageIndex is PageIndex Setter
// 当前页数
func (r *TmallmsdclaimqueryAPIRequest) SetPageIndex(_pageIndex int64) error {
	r._pageIndex = _pageIndex
	r.Set("page_index", _pageIndex)
	return nil
}

// GetPageIndex PageIndex Getter
func (r TmallmsdclaimqueryAPIRequest) GetPageIndex() int64 {
	return r._pageIndex
}

// SetPageSize is PageSize Setter
// 每页大小
func (r *TmallmsdclaimqueryAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TmallmsdclaimqueryAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetServiceOrderId is ServiceOrderId Setter
// 交易服务订单号
func (r *TmallmsdclaimqueryAPIRequest) SetServiceOrderId(_serviceOrderId int64) error {
	r._serviceOrderId = _serviceOrderId
	r.Set("service_order_id", _serviceOrderId)
	return nil
}

// GetServiceOrderId ServiceOrderId Getter
func (r TmallmsdclaimqueryAPIRequest) GetServiceOrderId() int64 {
	return r._serviceOrderId
}
