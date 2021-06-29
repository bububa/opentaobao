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
    _type   string
    // 拉取数据开始时间
    _startTime   string
    // 页码，从1开始
    _pageNo   int64
    // 每页数量，不能大于1000
    _pageSize   int64
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
func (r *TaobaoOmniorderGuideDataGetRequest) SetType(_type string) error {
    r._type = _type
    r.Set("type", _type)
    return nil
}

// Type Getter
func (r TaobaoOmniorderGuideDataGetRequest) GetType() string {
    return r._type
}
// StartTime Setter
// 拉取数据开始时间
func (r *TaobaoOmniorderGuideDataGetRequest) SetStartTime(_startTime string) error {
    r._startTime = _startTime
    r.Set("start_time", _startTime)
    return nil
}

// StartTime Getter
func (r TaobaoOmniorderGuideDataGetRequest) GetStartTime() string {
    return r._startTime
}
// PageNo Setter
// 页码，从1开始
func (r *TaobaoOmniorderGuideDataGetRequest) SetPageNo(_pageNo int64) error {
    r._pageNo = _pageNo
    r.Set("page_no", _pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoOmniorderGuideDataGetRequest) GetPageNo() int64 {
    return r._pageNo
}
// PageSize Setter
// 每页数量，不能大于1000
func (r *TaobaoOmniorderGuideDataGetRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoOmniorderGuideDataGetRequest) GetPageSize() int64 {
    return r._pageSize
}
