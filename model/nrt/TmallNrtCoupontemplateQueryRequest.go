package nrt

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
券模板查询 APIRequest
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

func NewTmallNrtCoupontemplateQueryRequest() *TmallNrtCoupontemplateQueryRequest{
    return &TmallNrtCoupontemplateQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TmallNrtCoupontemplateQueryRequest) GetApiMethodName() string {
    return "tmall.nrt.coupontemplate.query"
}

func (r TmallNrtCoupontemplateQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallNrtCoupontemplateQueryRequest) SetCouponTypeList(couponTypeList []int64) error {
    r.couponTypeList = couponTypeList
    r.Set("coupon_type_list", couponTypeList)
    return nil
}

func (r TmallNrtCoupontemplateQueryRequest) GetCouponTypeList() []int64 {
    return r.couponTypeList
}

func (r *TmallNrtCoupontemplateQueryRequest) SetCurrentPage(currentPage int64) error {
    r.currentPage = currentPage
    r.Set("current_page", currentPage)
    return nil
}

func (r TmallNrtCoupontemplateQueryRequest) GetCurrentPage() int64 {
    return r.currentPage
}

func (r *TmallNrtCoupontemplateQueryRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TmallNrtCoupontemplateQueryRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *TmallNrtCoupontemplateQueryRequest) SetBizCode(bizCode string) error {
    r.bizCode = bizCode
    r.Set("biz_code", bizCode)
    return nil
}

func (r TmallNrtCoupontemplateQueryRequest) GetBizCode() string {
    return r.bizCode
}

