package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询通用单品优惠活动列表 APIRequest
taobao.promotionmisc.common.item.activity.list.get

查询通用单品优惠活动列表。
*/
type TaobaoPromotionmiscCommonItemActivityListGetRequest struct {
    model.Params

    // 分页页码，页码从1开始
    pageNo   int64 

    // 分页大小，不能超过50
    pageSize   int64 

}

func NewTaobaoPromotionmiscCommonItemActivityListGetRequest() *TaobaoPromotionmiscCommonItemActivityListGetRequest{
    return &TaobaoPromotionmiscCommonItemActivityListGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoPromotionmiscCommonItemActivityListGetRequest) GetApiMethodName() string {
    return "taobao.promotionmisc.common.item.activity.list.get"
}

func (r TaobaoPromotionmiscCommonItemActivityListGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoPromotionmiscCommonItemActivityListGetRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

func (r TaobaoPromotionmiscCommonItemActivityListGetRequest) GetPageNo() int64 {
    return r.pageNo
}

func (r *TaobaoPromotionmiscCommonItemActivityListGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoPromotionmiscCommonItemActivityListGetRequest) GetPageSize() int64 {
    return r.pageSize
}

