package nrt

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallnrtcoupontemplatequeryAPIRequest 券模板查询 API请求
// tmall.nrt.coupontemplate.query
//
// 新零售场景，商家拉取在新零售工作台设置的券数据
type TmallnrtcoupontemplatequeryAPIRequest struct {
	model.Params
	// 券列表
	_couponTypeList []int64
	// 业务code阿里指定
	_bizCode string
	// 当前页
	_currentPage int64
	// 每页数据数量
	_pageSize int64
}

// NewTmallnrtcoupontemplatequeryRequest 初始化TmallnrtcoupontemplatequeryAPIRequest对象
func NewTmallnrtcoupontemplatequeryRequest() *TmallnrtcoupontemplatequeryAPIRequest {
	return &TmallnrtcoupontemplatequeryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallnrtcoupontemplatequeryAPIRequest) GetApiMethodName() string {
	return "tmall.nrt.coupontemplate.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallnrtcoupontemplatequeryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallnrtcoupontemplatequeryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCouponTypeList is CouponTypeList Setter
// 券列表
func (r *TmallnrtcoupontemplatequeryAPIRequest) SetCouponTypeList(_couponTypeList []int64) error {
	r._couponTypeList = _couponTypeList
	r.Set("coupon_type_list", _couponTypeList)
	return nil
}

// GetCouponTypeList CouponTypeList Getter
func (r TmallnrtcoupontemplatequeryAPIRequest) GetCouponTypeList() []int64 {
	return r._couponTypeList
}

// SetBizCode is BizCode Setter
// 业务code阿里指定
func (r *TmallnrtcoupontemplatequeryAPIRequest) SetBizCode(_bizCode string) error {
	r._bizCode = _bizCode
	r.Set("biz_code", _bizCode)
	return nil
}

// GetBizCode BizCode Getter
func (r TmallnrtcoupontemplatequeryAPIRequest) GetBizCode() string {
	return r._bizCode
}

// SetCurrentPage is CurrentPage Setter
// 当前页
func (r *TmallnrtcoupontemplatequeryAPIRequest) SetCurrentPage(_currentPage int64) error {
	r._currentPage = _currentPage
	r.Set("current_page", _currentPage)
	return nil
}

// GetCurrentPage CurrentPage Getter
func (r TmallnrtcoupontemplatequeryAPIRequest) GetCurrentPage() int64 {
	return r._currentPage
}

// SetPageSize is PageSize Setter
// 每页数据数量
func (r *TmallnrtcoupontemplatequeryAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TmallnrtcoupontemplatequeryAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
