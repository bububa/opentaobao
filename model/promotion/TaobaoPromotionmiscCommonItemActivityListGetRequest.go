package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询通用单品优惠活动列表 API请求
taobao.promotionmisc.common.item.activity.list.get

查询通用单品优惠活动列表。
*/
type TaobaoPromotionmiscCommonItemActivityListGetRequest struct {
    model.Params
    // 分页页码，页码从1开始
    _pageNo   int64
    // 分页大小，不能超过50
    _pageSize   int64
}

// 初始化TaobaoPromotionmiscCommonItemActivityListGetRequest对象
func NewTaobaoPromotionmiscCommonItemActivityListGetRequest() *TaobaoPromotionmiscCommonItemActivityListGetRequest{
    return &TaobaoPromotionmiscCommonItemActivityListGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoPromotionmiscCommonItemActivityListGetRequest) GetApiMethodName() string {
    return "taobao.promotionmisc.common.item.activity.list.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoPromotionmiscCommonItemActivityListGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PageNo Setter
// 分页页码，页码从1开始
func (r *TaobaoPromotionmiscCommonItemActivityListGetRequest) SetPageNo(_pageNo int64) error {
    r._pageNo = _pageNo
    r.Set("page_no", _pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoPromotionmiscCommonItemActivityListGetRequest) GetPageNo() int64 {
    return r._pageNo
}
// PageSize Setter
// 分页大小，不能超过50
func (r *TaobaoPromotionmiscCommonItemActivityListGetRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoPromotionmiscCommonItemActivityListGetRequest) GetPageSize() int64 {
    return r._pageSize
}
