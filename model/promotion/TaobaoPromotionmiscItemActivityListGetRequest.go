package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询无条件单品优惠活动列表 API请求
taobao.promotionmisc.item.activity.list.get

查询无条件单品优惠活动列表
*/
type TaobaoPromotionmiscItemActivityListGetRequest struct {
    model.Params
    // 页码。
    _pageNo   int64
    // 每页记录数，最大支持50 。
    _pageSize   int64
}

// 初始化TaobaoPromotionmiscItemActivityListGetRequest对象
func NewTaobaoPromotionmiscItemActivityListGetRequest() *TaobaoPromotionmiscItemActivityListGetRequest{
    return &TaobaoPromotionmiscItemActivityListGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoPromotionmiscItemActivityListGetRequest) GetApiMethodName() string {
    return "taobao.promotionmisc.item.activity.list.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoPromotionmiscItemActivityListGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PageNo Setter
// 页码。
func (r *TaobaoPromotionmiscItemActivityListGetRequest) SetPageNo(_pageNo int64) error {
    r._pageNo = _pageNo
    r.Set("page_no", _pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoPromotionmiscItemActivityListGetRequest) GetPageNo() int64 {
    return r._pageNo
}
// PageSize Setter
// 每页记录数，最大支持50 。
func (r *TaobaoPromotionmiscItemActivityListGetRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoPromotionmiscItemActivityListGetRequest) GetPageSize() int64 {
    return r._pageSize
}
