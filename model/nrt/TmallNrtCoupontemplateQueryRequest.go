package nrt

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
券模板查询 API请求
tmall.nrt.coupontemplate.query

新零售场景，商家拉取在新零售工作台设置的券数据
*/
type TmallNrtCoupontemplateQueryRequest struct {
    model.Params
    // 券列表
    _couponTypeList   []int64
    // 当前页
    _currentPage   int64
    // 每页数据数量
    _pageSize   int64
    // 业务code阿里指定
    _bizCode   string
}

// 初始化TmallNrtCoupontemplateQueryRequest对象
func NewTmallNrtCoupontemplateQueryRequest() *TmallNrtCoupontemplateQueryRequest{
    return &TmallNrtCoupontemplateQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallNrtCoupontemplateQueryRequest) GetApiMethodName() string {
    return "tmall.nrt.coupontemplate.query"
}

// IRequest interface 方法, 获取API参数
func (r TmallNrtCoupontemplateQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CouponTypeList Setter
// 券列表
func (r *TmallNrtCoupontemplateQueryRequest) SetCouponTypeList(_couponTypeList []int64) error {
    r._couponTypeList = _couponTypeList
    r.Set("coupon_type_list", _couponTypeList)
    return nil
}

// CouponTypeList Getter
func (r TmallNrtCoupontemplateQueryRequest) GetCouponTypeList() []int64 {
    return r._couponTypeList
}
// CurrentPage Setter
// 当前页
func (r *TmallNrtCoupontemplateQueryRequest) SetCurrentPage(_currentPage int64) error {
    r._currentPage = _currentPage
    r.Set("current_page", _currentPage)
    return nil
}

// CurrentPage Getter
func (r TmallNrtCoupontemplateQueryRequest) GetCurrentPage() int64 {
    return r._currentPage
}
// PageSize Setter
// 每页数据数量
func (r *TmallNrtCoupontemplateQueryRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TmallNrtCoupontemplateQueryRequest) GetPageSize() int64 {
    return r._pageSize
}
// BizCode Setter
// 业务code阿里指定
func (r *TmallNrtCoupontemplateQueryRequest) SetBizCode(_bizCode string) error {
    r._bizCode = _bizCode
    r.Set("biz_code", _bizCode)
    return nil
}

// BizCode Getter
func (r TmallNrtCoupontemplateQueryRequest) GetBizCode() string {
    return r._bizCode
}
