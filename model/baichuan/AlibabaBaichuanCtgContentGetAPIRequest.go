package baichuan

import (
	"net/url"

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
	// 分页大小
	_pageSize int64
	// 当前页
	_currentPage int64
	// 资源位
	_resId string
	// 日期
	_date string
}

// NewAlibabaBaichuanCtgContentGetRequest 初始化AlibabaBaichuanCtgContentGetAPIRequest对象
func NewAlibabaBaichuanCtgContentGetRequest() *AlibabaBaichuanCtgContentGetAPIRequest {
	return &AlibabaBaichuanCtgContentGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaBaichuanCtgContentGetAPIRequest) GetApiMethodName() string {
	return "alibaba.baichuan.ctg.content.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaBaichuanCtgContentGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is DeliveryId Setter
// 投放位置
func (r *AlibabaBaichuanCtgContentGetAPIRequest) SetDeliveryId(_deliveryId string) error {
	r._deliveryId = _deliveryId
	r.Set("delivery_id", _deliveryId)
	return nil
}

// Get DeliveryId Getter
func (r AlibabaBaichuanCtgContentGetAPIRequest) GetDeliveryId() string {
	return r._deliveryId
}

// Set is PageSize Setter
// 分页大小
func (r *AlibabaBaichuanCtgContentGetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// Get PageSize Getter
func (r AlibabaBaichuanCtgContentGetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// Set is CurrentPage Setter
// 当前页
func (r *AlibabaBaichuanCtgContentGetAPIRequest) SetCurrentPage(_currentPage int64) error {
	r._currentPage = _currentPage
	r.Set("current_page", _currentPage)
	return nil
}

// Get CurrentPage Getter
func (r AlibabaBaichuanCtgContentGetAPIRequest) GetCurrentPage() int64 {
	return r._currentPage
}

// Set is ResId Setter
// 资源位
func (r *AlibabaBaichuanCtgContentGetAPIRequest) SetResId(_resId string) error {
	r._resId = _resId
	r.Set("res_id", _resId)
	return nil
}

// Get ResId Getter
func (r AlibabaBaichuanCtgContentGetAPIRequest) GetResId() string {
	return r._resId
}

// Set is Date Setter
// 日期
func (r *AlibabaBaichuanCtgContentGetAPIRequest) SetDate(_date string) error {
	r._date = _date
	r.Set("date", _date)
	return nil
}

// Get Date Getter
func (r AlibabaBaichuanCtgContentGetAPIRequest) GetDate() string {
	return r._date
}
