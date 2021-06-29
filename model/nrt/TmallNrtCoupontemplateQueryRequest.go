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
    couponTypeList   []int64
    // 当前页
    currentPage   int64
    // 每页数据数量
    pageSize   int64
    // 业务code阿里指定
    bizCode   string
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
func (r *TmallNrtCoupontemplateQueryRequest) SetCouponTypeList(couponTypeList []int64) error {
    r.couponTypeList = couponTypeList
    r.Set("coupon_type_list", couponTypeList)
    return nil
}

// CouponTypeList Getter
func (r TmallNrtCoupontemplateQueryRequest) GetCouponTypeList() []int64 {
    return r.couponTypeList
}
// CurrentPage Setter
// 当前页
func (r *TmallNrtCoupontemplateQueryRequest) SetCurrentPage(currentPage int64) error {
    r.currentPage = currentPage
    r.Set("current_page", currentPage)
    return nil
}

// CurrentPage Getter
func (r TmallNrtCoupontemplateQueryRequest) GetCurrentPage() int64 {
    return r.currentPage
}
// PageSize Setter
// 每页数据数量
func (r *TmallNrtCoupontemplateQueryRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r TmallNrtCoupontemplateQueryRequest) GetPageSize() int64 {
    return r.pageSize
}
// BizCode Setter
// 业务code阿里指定
func (r *TmallNrtCoupontemplateQueryRequest) SetBizCode(bizCode string) error {
    r.bizCode = bizCode
    r.Set("biz_code", bizCode)
    return nil
}

// BizCode Getter
func (r TmallNrtCoupontemplateQueryRequest) GetBizCode() string {
    return r.bizCode
}
