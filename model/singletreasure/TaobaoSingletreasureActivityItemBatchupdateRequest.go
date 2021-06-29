package singletreasure

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批量修改商品接口 API请求
taobao.singletreasure.activity.item.batchupdate

批量修改商品优惠接口
*/
type TaobaoSingletreasureActivityItemBatchupdateRequest struct {
    model.Params
    // 系统入参
    _itemDetailInfo   *ItemDetailInfoBatchCreateDTO
}

// 初始化TaobaoSingletreasureActivityItemBatchupdateRequest对象
func NewTaobaoSingletreasureActivityItemBatchupdateRequest() *TaobaoSingletreasureActivityItemBatchupdateRequest{
    return &TaobaoSingletreasureActivityItemBatchupdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSingletreasureActivityItemBatchupdateRequest) GetApiMethodName() string {
    return "taobao.singletreasure.activity.item.batchupdate"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSingletreasureActivityItemBatchupdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemDetailInfo Setter
// 系统入参
func (r *TaobaoSingletreasureActivityItemBatchupdateRequest) SetItemDetailInfo(_itemDetailInfo *ItemDetailInfoBatchCreateDTO) error {
    r._itemDetailInfo = _itemDetailInfo
    r.Set("item_detail_info", _itemDetailInfo)
    return nil
}

// ItemDetailInfo Getter
func (r TaobaoSingletreasureActivityItemBatchupdateRequest) GetItemDetailInfo() *ItemDetailInfoBatchCreateDTO {
    return r._itemDetailInfo
}
