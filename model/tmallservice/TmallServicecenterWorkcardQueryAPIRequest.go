package tmallservice

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterWorkcardQueryAPIRequest 工单查询接口 API请求
// tmall.servicecenter.workcard.query
//
// 工单查询接口
type TmallServicecenterWorkcardQueryAPIRequest struct {
	model.Params
	// 核销码
	_identifyCode string
	// 工单创建开始时间
	_gmtCreateStart string
	// 工单创建结束时间，必须与工单创建开始时间一起传入，且间隔不超过15分钟
	_gmtCreateEnd string
	// 门店/网点id
	_serviceStoreId int64
	// 工单id
	_id int64
	// 淘宝交易订单号。主订单或子订单均可
	_bizOrderId int64
	// 当前页数
	_currentPage int64
	// 每页大小
	_pageSize int64
	// 是否查询操作记录信息
	_needWorkcardOperation bool
}

// NewTmallServicecenterWorkcardQueryRequest 初始化TmallServicecenterWorkcardQueryAPIRequest对象
func NewTmallServicecenterWorkcardQueryRequest() *TmallServicecenterWorkcardQueryAPIRequest {
	return &TmallServicecenterWorkcardQueryAPIRequest{
		Params: model.NewParams(9),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallServicecenterWorkcardQueryAPIRequest) Reset() {
	r._identifyCode = ""
	r._gmtCreateStart = ""
	r._gmtCreateEnd = ""
	r._serviceStoreId = 0
	r._id = 0
	r._bizOrderId = 0
	r._currentPage = 0
	r._pageSize = 0
	r._needWorkcardOperation = false
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServicecenterWorkcardQueryAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.workcard.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServicecenterWorkcardQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallServicecenterWorkcardQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetIdentifyCode is IdentifyCode Setter
// 核销码
func (r *TmallServicecenterWorkcardQueryAPIRequest) SetIdentifyCode(_identifyCode string) error {
	r._identifyCode = _identifyCode
	r.Set("identify_code", _identifyCode)
	return nil
}

// GetIdentifyCode IdentifyCode Getter
func (r TmallServicecenterWorkcardQueryAPIRequest) GetIdentifyCode() string {
	return r._identifyCode
}

// SetGmtCreateStart is GmtCreateStart Setter
// 工单创建开始时间
func (r *TmallServicecenterWorkcardQueryAPIRequest) SetGmtCreateStart(_gmtCreateStart string) error {
	r._gmtCreateStart = _gmtCreateStart
	r.Set("gmt_create_start", _gmtCreateStart)
	return nil
}

// GetGmtCreateStart GmtCreateStart Getter
func (r TmallServicecenterWorkcardQueryAPIRequest) GetGmtCreateStart() string {
	return r._gmtCreateStart
}

// SetGmtCreateEnd is GmtCreateEnd Setter
// 工单创建结束时间，必须与工单创建开始时间一起传入，且间隔不超过15分钟
func (r *TmallServicecenterWorkcardQueryAPIRequest) SetGmtCreateEnd(_gmtCreateEnd string) error {
	r._gmtCreateEnd = _gmtCreateEnd
	r.Set("gmt_create_end", _gmtCreateEnd)
	return nil
}

// GetGmtCreateEnd GmtCreateEnd Getter
func (r TmallServicecenterWorkcardQueryAPIRequest) GetGmtCreateEnd() string {
	return r._gmtCreateEnd
}

// SetServiceStoreId is ServiceStoreId Setter
// 门店/网点id
func (r *TmallServicecenterWorkcardQueryAPIRequest) SetServiceStoreId(_serviceStoreId int64) error {
	r._serviceStoreId = _serviceStoreId
	r.Set("service_store_id", _serviceStoreId)
	return nil
}

// GetServiceStoreId ServiceStoreId Getter
func (r TmallServicecenterWorkcardQueryAPIRequest) GetServiceStoreId() int64 {
	return r._serviceStoreId
}

// SetId is Id Setter
// 工单id
func (r *TmallServicecenterWorkcardQueryAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r TmallServicecenterWorkcardQueryAPIRequest) GetId() int64 {
	return r._id
}

// SetBizOrderId is BizOrderId Setter
// 淘宝交易订单号。主订单或子订单均可
func (r *TmallServicecenterWorkcardQueryAPIRequest) SetBizOrderId(_bizOrderId int64) error {
	r._bizOrderId = _bizOrderId
	r.Set("biz_order_id", _bizOrderId)
	return nil
}

// GetBizOrderId BizOrderId Getter
func (r TmallServicecenterWorkcardQueryAPIRequest) GetBizOrderId() int64 {
	return r._bizOrderId
}

// SetCurrentPage is CurrentPage Setter
// 当前页数
func (r *TmallServicecenterWorkcardQueryAPIRequest) SetCurrentPage(_currentPage int64) error {
	r._currentPage = _currentPage
	r.Set("current_page", _currentPage)
	return nil
}

// GetCurrentPage CurrentPage Getter
func (r TmallServicecenterWorkcardQueryAPIRequest) GetCurrentPage() int64 {
	return r._currentPage
}

// SetPageSize is PageSize Setter
// 每页大小
func (r *TmallServicecenterWorkcardQueryAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TmallServicecenterWorkcardQueryAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetNeedWorkcardOperation is NeedWorkcardOperation Setter
// 是否查询操作记录信息
func (r *TmallServicecenterWorkcardQueryAPIRequest) SetNeedWorkcardOperation(_needWorkcardOperation bool) error {
	r._needWorkcardOperation = _needWorkcardOperation
	r.Set("need_workcard_operation", _needWorkcardOperation)
	return nil
}

// GetNeedWorkcardOperation NeedWorkcardOperation Getter
func (r TmallServicecenterWorkcardQueryAPIRequest) GetNeedWorkcardOperation() bool {
	return r._needWorkcardOperation
}

var poolTmallServicecenterWorkcardQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallServicecenterWorkcardQueryRequest()
	},
}

// GetTmallServicecenterWorkcardQueryRequest 从 sync.Pool 获取 TmallServicecenterWorkcardQueryAPIRequest
func GetTmallServicecenterWorkcardQueryAPIRequest() *TmallServicecenterWorkcardQueryAPIRequest {
	return poolTmallServicecenterWorkcardQueryAPIRequest.Get().(*TmallServicecenterWorkcardQueryAPIRequest)
}

// ReleaseTmallServicecenterWorkcardQueryAPIRequest 将 TmallServicecenterWorkcardQueryAPIRequest 放入 sync.Pool
func ReleaseTmallServicecenterWorkcardQueryAPIRequest(v *TmallServicecenterWorkcardQueryAPIRequest) {
	v.Reset()
	poolTmallServicecenterWorkcardQueryAPIRequest.Put(v)
}
