package omniorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取全渠道导购产品数据 API请求
taobao.omniorder.guide.data.get

获取全渠道导购产品，目前包括随心购、随身购扫码、加购和交易数据。
*/
type TaobaoOmniorderGuideDataGetRequest struct {
    model.Params
    // detail_smg_scan: 扫码购扫码明细;detail_smg_cart: 扫码购加购明细;detail_smg_order: 扫码购订单明细;detail_sxg_search: 随心购搜索明细;detail_sxg_view_item: 随心购商品浏览明细;detail_sxg_cart: 随心购加购明细;detail_sxg_order: 随心购订单明细
    type   string
    // 拉取数据开始时间
    startTime   string
    // 页码，从1开始
    pageNo   int64
    // 每页数量，不能大于1000
    pageSize   int64
}

// 初始化TaobaoOmniorderGuideDataGetRequest对象
func NewTaobaoOmniorderGuideDataGetRequest() *TaobaoOmniorderGuideDataGetRequest{
    return &TaobaoOmniorderGuideDataGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOmniorderGuideDataGetRequest) GetApiMethodName() string {
    return "taobao.omniorder.guide.data.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOmniorderGuideDataGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Type Setter
// detail_smg_scan: 扫码购扫码明细;detail_smg_cart: 扫码购加购明细;detail_smg_order: 扫码购订单明细;detail_sxg_search: 随心购搜索明细;detail_sxg_view_item: 随心购商品浏览明细;detail_sxg_cart: 随心购加购明细;detail_sxg_order: 随心购订单明细
func (r *TaobaoOmniorderGuideDataGetRequest) SetType(type string) error {
    r.type = type
    r.Set("type", type)
    return nil
}

// Type Getter
func (r TaobaoOmniorderGuideDataGetRequest) GetType() string {
    return r.type
}
// StartTime Setter
// 拉取数据开始时间
func (r *TaobaoOmniorderGuideDataGetRequest) SetStartTime(startTime string) error {
    r.startTime = startTime
    r.Set("start_time", startTime)
    return nil
}

// StartTime Getter
func (r TaobaoOmniorderGuideDataGetRequest) GetStartTime() string {
    return r.startTime
}
// PageNo Setter
// 页码，从1开始
func (r *TaobaoOmniorderGuideDataGetRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoOmniorderGuideDataGetRequest) GetPageNo() int64 {
    return r.pageNo
}
// PageSize Setter
// 每页数量，不能大于1000
func (r *TaobaoOmniorderGuideDataGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoOmniorderGuideDataGetRequest) GetPageSize() int64 {
    return r.pageSize
}
