package singletreasure

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
更新单品优惠接口 APIRequest
taobao.singletreasure.activity.item.update

更新单品优惠接口
*/
type TaobaoSingletreasureActivityItemUpdateRequest struct {
    model.Params

    // 修改接口的入参对象
    itemDetailInfo   *ItemDetailInfoCreateDto 

}

func NewTaobaoSingletreasureActivityItemUpdateRequest() *TaobaoSingletreasureActivityItemUpdateRequest{
    return &TaobaoSingletreasureActivityItemUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSingletreasureActivityItemUpdateRequest) GetApiMethodName() string {
    return "taobao.singletreasure.activity.item.update"
}

func (r TaobaoSingletreasureActivityItemUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSingletreasureActivityItemUpdateRequest) SetItemDetailInfo(itemDetailInfo *ItemDetailInfoCreateDto) error {
    r.itemDetailInfo = itemDetailInfo
    r.Set("item_detail_info", itemDetailInfo)
    return nil
}

func (r TaobaoSingletreasureActivityItemUpdateRequest) GetItemDetailInfo() *ItemDetailInfoCreateDto {
    return r.itemDetailInfo
}

