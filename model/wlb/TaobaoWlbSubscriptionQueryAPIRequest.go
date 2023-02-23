package wlb

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbSubscriptionQueryAPIRequest 查询商家定购的所有服务 API请求
// taobao.wlb.subscription.query
//
// 查询商家定购的所有服务,可通过入参状态来筛选
type TaobaoWlbSubscriptionQueryAPIRequest struct {
	model.Params
	// 状态 <br/>AUDITING 1-待审核; <br/>CANCEL 2-撤销 ;<br/>CHECKED 3-审核通过 ;<br/>FAILED 4-审核未通过 ;<br/>SYNCHRONIZING 5-同步中;<br/>只允许输入上面指定的值，且可以为空，为空时查询所有状态。若输错了，则按AUDITING处理。
	_status string
	// 当前页
	_pageNo int64
	// 分页记录个数，如果用户输入的记录数大于50，则一页显示50条记录
	_pageSize int64
}

// NewTaobaoWlbSubscriptionQueryRequest 初始化TaobaoWlbSubscriptionQueryAPIRequest对象
func NewTaobaoWlbSubscriptionQueryRequest() *TaobaoWlbSubscriptionQueryAPIRequest {
	return &TaobaoWlbSubscriptionQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWlbSubscriptionQueryAPIRequest) GetApiMethodName() string {
	return "taobao.wlb.subscription.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWlbSubscriptionQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoWlbSubscriptionQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStatus is Status Setter
// 状态 &lt;br/&gt;AUDITING 1-待审核; &lt;br/&gt;CANCEL 2-撤销 ;&lt;br/&gt;CHECKED 3-审核通过 ;&lt;br/&gt;FAILED 4-审核未通过 ;&lt;br/&gt;SYNCHRONIZING 5-同步中;&lt;br/&gt;只允许输入上面指定的值，且可以为空，为空时查询所有状态。若输错了，则按AUDITING处理。
func (r *TaobaoWlbSubscriptionQueryAPIRequest) SetStatus(_status string) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r TaobaoWlbSubscriptionQueryAPIRequest) GetStatus() string {
	return r._status
}

// SetPageNo is PageNo Setter
// 当前页
func (r *TaobaoWlbSubscriptionQueryAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaoWlbSubscriptionQueryAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// 分页记录个数，如果用户输入的记录数大于50，则一页显示50条记录
func (r *TaobaoWlbSubscriptionQueryAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoWlbSubscriptionQueryAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
