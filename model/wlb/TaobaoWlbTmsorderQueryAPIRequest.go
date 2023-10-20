package wlb

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbTmsorderQueryAPIRequest 通过物流订单编号查询物流信息 API请求
// taobao.wlb.tmsorder.query
//
// 通过物流订单编号分页查询物流信息
type TaobaoWlbTmsorderQueryAPIRequest struct {
	model.Params
	// 物流订单编号
	_orderCode string
	// 当前页
	_pageNo int64
	// 分页记录个数，如果用户输入的记录数大于50，则一页显示50条记录
	_pageSize int64
}

// NewTaobaoWlbTmsorderQueryRequest 初始化TaobaoWlbTmsorderQueryAPIRequest对象
func NewTaobaoWlbTmsorderQueryRequest() *TaobaoWlbTmsorderQueryAPIRequest {
	return &TaobaoWlbTmsorderQueryAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoWlbTmsorderQueryAPIRequest) Reset() {
	r._orderCode = ""
	r._pageNo = 0
	r._pageSize = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWlbTmsorderQueryAPIRequest) GetApiMethodName() string {
	return "taobao.wlb.tmsorder.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWlbTmsorderQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoWlbTmsorderQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderCode is OrderCode Setter
// 物流订单编号
func (r *TaobaoWlbTmsorderQueryAPIRequest) SetOrderCode(_orderCode string) error {
	r._orderCode = _orderCode
	r.Set("order_code", _orderCode)
	return nil
}

// GetOrderCode OrderCode Getter
func (r TaobaoWlbTmsorderQueryAPIRequest) GetOrderCode() string {
	return r._orderCode
}

// SetPageNo is PageNo Setter
// 当前页
func (r *TaobaoWlbTmsorderQueryAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaoWlbTmsorderQueryAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// 分页记录个数，如果用户输入的记录数大于50，则一页显示50条记录
func (r *TaobaoWlbTmsorderQueryAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoWlbTmsorderQueryAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

var poolTaobaoWlbTmsorderQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoWlbTmsorderQueryRequest()
	},
}

// GetTaobaoWlbTmsorderQueryRequest 从 sync.Pool 获取 TaobaoWlbTmsorderQueryAPIRequest
func GetTaobaoWlbTmsorderQueryAPIRequest() *TaobaoWlbTmsorderQueryAPIRequest {
	return poolTaobaoWlbTmsorderQueryAPIRequest.Get().(*TaobaoWlbTmsorderQueryAPIRequest)
}

// ReleaseTaobaoWlbTmsorderQueryAPIRequest 将 TaobaoWlbTmsorderQueryAPIRequest 放入 sync.Pool
func ReleaseTaobaoWlbTmsorderQueryAPIRequest(v *TaobaoWlbTmsorderQueryAPIRequest) {
	v.Reset()
	poolTaobaoWlbTmsorderQueryAPIRequest.Put(v)
}
