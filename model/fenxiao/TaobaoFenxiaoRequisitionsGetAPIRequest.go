package fenxiao

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFenxiaoRequisitionsGetAPIRequest 合作申请查询 API请求
// taobao.fenxiao.requisitions.get
//
// 合作申请查询
type TaobaoFenxiaoRequisitionsGetAPIRequest struct {
	model.Params
	// 申请开始时间yyyy-MM-dd
	_applyStart string
	// 申请结束时间yyyy-MM-dd
	_applyEnd string
	// 申请状态（1-申请中、2-成功、3-被退回、4-已撤消、5-过期）
	_status int64
	// 页码（大于0的整数，默认1）
	_pageNo int64
	// 每页记录数（默认20，最大50）
	_pageSize int64
}

// NewTaobaoFenxiaoRequisitionsGetRequest 初始化TaobaoFenxiaoRequisitionsGetAPIRequest对象
func NewTaobaoFenxiaoRequisitionsGetRequest() *TaobaoFenxiaoRequisitionsGetAPIRequest {
	return &TaobaoFenxiaoRequisitionsGetAPIRequest{
		Params: model.NewParams(5),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoFenxiaoRequisitionsGetAPIRequest) Reset() {
	r._applyStart = ""
	r._applyEnd = ""
	r._status = 0
	r._pageNo = 0
	r._pageSize = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFenxiaoRequisitionsGetAPIRequest) GetApiMethodName() string {
	return "taobao.fenxiao.requisitions.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFenxiaoRequisitionsGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoFenxiaoRequisitionsGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetApplyStart is ApplyStart Setter
// 申请开始时间yyyy-MM-dd
func (r *TaobaoFenxiaoRequisitionsGetAPIRequest) SetApplyStart(_applyStart string) error {
	r._applyStart = _applyStart
	r.Set("apply_start", _applyStart)
	return nil
}

// GetApplyStart ApplyStart Getter
func (r TaobaoFenxiaoRequisitionsGetAPIRequest) GetApplyStart() string {
	return r._applyStart
}

// SetApplyEnd is ApplyEnd Setter
// 申请结束时间yyyy-MM-dd
func (r *TaobaoFenxiaoRequisitionsGetAPIRequest) SetApplyEnd(_applyEnd string) error {
	r._applyEnd = _applyEnd
	r.Set("apply_end", _applyEnd)
	return nil
}

// GetApplyEnd ApplyEnd Getter
func (r TaobaoFenxiaoRequisitionsGetAPIRequest) GetApplyEnd() string {
	return r._applyEnd
}

// SetStatus is Status Setter
// 申请状态（1-申请中、2-成功、3-被退回、4-已撤消、5-过期）
func (r *TaobaoFenxiaoRequisitionsGetAPIRequest) SetStatus(_status int64) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r TaobaoFenxiaoRequisitionsGetAPIRequest) GetStatus() int64 {
	return r._status
}

// SetPageNo is PageNo Setter
// 页码（大于0的整数，默认1）
func (r *TaobaoFenxiaoRequisitionsGetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaoFenxiaoRequisitionsGetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// 每页记录数（默认20，最大50）
func (r *TaobaoFenxiaoRequisitionsGetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoFenxiaoRequisitionsGetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

var poolTaobaoFenxiaoRequisitionsGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoFenxiaoRequisitionsGetRequest()
	},
}

// GetTaobaoFenxiaoRequisitionsGetRequest 从 sync.Pool 获取 TaobaoFenxiaoRequisitionsGetAPIRequest
func GetTaobaoFenxiaoRequisitionsGetAPIRequest() *TaobaoFenxiaoRequisitionsGetAPIRequest {
	return poolTaobaoFenxiaoRequisitionsGetAPIRequest.Get().(*TaobaoFenxiaoRequisitionsGetAPIRequest)
}

// ReleaseTaobaoFenxiaoRequisitionsGetAPIRequest 将 TaobaoFenxiaoRequisitionsGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoFenxiaoRequisitionsGetAPIRequest(v *TaobaoFenxiaoRequisitionsGetAPIRequest) {
	v.Reset()
	poolTaobaoFenxiaoRequisitionsGetAPIRequest.Put(v)
}
