package singletreasure

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批量添加商品接口 APIRequest
taobao.singletreasure.activity.item.batchadd

向活动中批量添加商品优惠
*/
type TaobaoSingletreasureActivityItemBatchaddRequest struct {
    model.Params

    // 系统入参
    itemDetailInfo   *ItemDetailInfoBatchCreateDto 

}

func NewTaobaoSingletreasureActivityItemBatchaddRequest() *TaobaoSingletreasureActivityItemBatchaddRequest{
    return &TaobaoSingletreasureActivityItemBatchaddRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSingletreasureActivityItemBatchaddRequest) GetApiMethodName() string {
    return "taobao.singletreasure.activity.item.batchadd"
}

func (r TaobaoSingletreasureActivityItemBatchaddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSingletreasureActivityItemBatchaddRequest) SetItemDetailInfo(itemDetailInfo *ItemDetailInfoBatchCreateDto) error {
    r.itemDetailInfo = itemDetailInfo
    r.Set("item_detail_info", itemDetailInfo)
    return nil
}

func (r TaobaoSingletreasureActivityItemBatchaddRequest) GetItemDetailInfo() *ItemDetailInfoBatchCreateDto {
    return r.itemDetailInfo
}

