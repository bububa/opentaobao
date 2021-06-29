package singletreasure

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批量修改商品接口 APIRequest
taobao.singletreasure.activity.item.batchupdate

批量修改商品优惠接口
*/
type TaobaoSingletreasureActivityItemBatchupdateRequest struct {
    model.Params

    // 系统入参
    itemDetailInfo   *ItemDetailInfoBatchCreateDto 

}

func NewTaobaoSingletreasureActivityItemBatchupdateRequest() *TaobaoSingletreasureActivityItemBatchupdateRequest{
    return &TaobaoSingletreasureActivityItemBatchupdateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSingletreasureActivityItemBatchupdateRequest) GetApiMethodName() string {
    return "taobao.singletreasure.activity.item.batchupdate"
}

func (r TaobaoSingletreasureActivityItemBatchupdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSingletreasureActivityItemBatchupdateRequest) SetItemDetailInfo(itemDetailInfo *ItemDetailInfoBatchCreateDto) error {
    r.itemDetailInfo = itemDetailInfo
    r.Set("item_detail_info", itemDetailInfo)
    return nil
}

func (r TaobaoSingletreasureActivityItemBatchupdateRequest) GetItemDetailInfo() *ItemDetailInfoBatchCreateDto {
    return r.itemDetailInfo
}

