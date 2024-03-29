package baichuan

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaBaichuanCtgContentGetAPIRequest 百川内容平台内容获取 API请求
// alibaba.baichuan.ctg.content.get
//
// 百川内容平台内容获取
type AlibabaBaichuanCtgContentGetAPIRequest struct {
	model.Params
	// 投放位置
	_deliveryId string
	// 资源位
	_resId string
	// 日期
	_date string
	// 分页大小
	_pageSize int64
	// 当前页
	_currentPage int64
}

// NewAlibabaBaichuanCtgContentGetRequest 初始化AlibabaBaichuanCtgContentGetAPIRequest对象
func NewAlibabaBaichuanCtgContentGetRequest() *AlibabaBaichuanCtgContentGetAPIRequest {
	return &AlibabaBaichuanCtgContentGetAPIRequest{
		Params: model.NewParams(5),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaBaichuanCtgContentGetAPIRequest) Reset() {
	r._deliveryId = ""
	r._resId = ""
	r._date = ""
	r._pageSize = 0
	r._currentPage = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaBaichuanCtgContentGetAPIRequest) GetApiMethodName() string {
	return "alibaba.baichuan.ctg.content.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaBaichuanCtgContentGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaBaichuanCtgContentGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDeliveryId is DeliveryId Setter
// 投放位置
func (r *AlibabaBaichuanCtgContentGetAPIRequest) SetDeliveryId(_deliveryId string) error {
	r._deliveryId = _deliveryId
	r.Set("delivery_id", _deliveryId)
	return nil
}

// GetDeliveryId DeliveryId Getter
func (r AlibabaBaichuanCtgContentGetAPIRequest) GetDeliveryId() string {
	return r._deliveryId
}

// SetResId is ResId Setter
// 资源位
func (r *AlibabaBaichuanCtgContentGetAPIRequest) SetResId(_resId string) error {
	r._resId = _resId
	r.Set("res_id", _resId)
	return nil
}

// GetResId ResId Getter
func (r AlibabaBaichuanCtgContentGetAPIRequest) GetResId() string {
	return r._resId
}

// SetDate is Date Setter
// 日期
func (r *AlibabaBaichuanCtgContentGetAPIRequest) SetDate(_date string) error {
	r._date = _date
	r.Set("date", _date)
	return nil
}

// GetDate Date Getter
func (r AlibabaBaichuanCtgContentGetAPIRequest) GetDate() string {
	return r._date
}

// SetPageSize is PageSize Setter
// 分页大小
func (r *AlibabaBaichuanCtgContentGetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AlibabaBaichuanCtgContentGetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetCurrentPage is CurrentPage Setter
// 当前页
func (r *AlibabaBaichuanCtgContentGetAPIRequest) SetCurrentPage(_currentPage int64) error {
	r._currentPage = _currentPage
	r.Set("current_page", _currentPage)
	return nil
}

// GetCurrentPage CurrentPage Getter
func (r AlibabaBaichuanCtgContentGetAPIRequest) GetCurrentPage() int64 {
	return r._currentPage
}

var poolAlibabaBaichuanCtgContentGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaBaichuanCtgContentGetRequest()
	},
}

// GetAlibabaBaichuanCtgContentGetRequest 从 sync.Pool 获取 AlibabaBaichuanCtgContentGetAPIRequest
func GetAlibabaBaichuanCtgContentGetAPIRequest() *AlibabaBaichuanCtgContentGetAPIRequest {
	return poolAlibabaBaichuanCtgContentGetAPIRequest.Get().(*AlibabaBaichuanCtgContentGetAPIRequest)
}

// ReleaseAlibabaBaichuanCtgContentGetAPIRequest 将 AlibabaBaichuanCtgContentGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaBaichuanCtgContentGetAPIRequest(v *AlibabaBaichuanCtgContentGetAPIRequest) {
	v.Reset()
	poolAlibabaBaichuanCtgContentGetAPIRequest.Put(v)
}
