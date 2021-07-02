package nrt

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallNrtCoupontemplateQueryAPIRequest 券模板查询 API请求
// tmall.nrt.coupontemplate.query
//
// 新零售场景，商家拉取在新零售工作台设置的券数据
type TmallNrtCoupontemplateQueryAPIRequest struct {
	model.Params
	// 券列表
	_couponTypeList []int64
	// 当前页
	_currentPage int64
	// 每页数据数量
	_pageSize int64
	// 业务code阿里指定
	_bizCode string
}

// NewTmallNrtCoupontemplateQueryRequest 初始化TmallNrtCoupontemplateQueryAPIRequest对象
func NewTmallNrtCoupontemplateQueryRequest() *TmallNrtCoupontemplateQueryAPIRequest {
	return &TmallNrtCoupontemplateQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallNrtCoupontemplateQueryAPIRequest) GetApiMethodName() string {
	return "tmall.nrt.coupontemplate.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallNrtCoupontemplateQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetCouponTypeList is CouponTypeList Setter
// 券列表
func (r *TmallNrtCoupontemplateQueryAPIRequest) SetCouponTypeList(_couponTypeList []int64) error {
	r._couponTypeList = _couponTypeList
	r.Set("coupon_type_list", _couponTypeList)
	return nil
}

// GetCouponTypeList CouponTypeList Getter
func (r TmallNrtCoupontemplateQueryAPIRequest) GetCouponTypeList() []int64 {
	return r._couponTypeList
}

// SetCurrentPage is CurrentPage Setter
// 当前页
func (r *TmallNrtCoupontemplateQueryAPIRequest) SetCurrentPage(_currentPage int64) error {
	r._currentPage = _currentPage
	r.Set("current_page", _currentPage)
	return nil
}

// GetCurrentPage CurrentPage Getter
func (r TmallNrtCoupontemplateQueryAPIRequest) GetCurrentPage() int64 {
	return r._currentPage
}

// SetPageSize is PageSize Setter
// 每页数据数量
func (r *TmallNrtCoupontemplateQueryAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TmallNrtCoupontemplateQueryAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetBizCode is BizCode Setter
// 业务code阿里指定
func (r *TmallNrtCoupontemplateQueryAPIRequest) SetBizCode(_bizCode string) error {
	r._bizCode = _bizCode
	r.Set("biz_code", _bizCode)
	return nil
}

// GetBizCode BizCode Getter
func (r TmallNrtCoupontemplateQueryAPIRequest) GetBizCode() string {
	return r._bizCode
}
