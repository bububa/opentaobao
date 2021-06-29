package singletreasure

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批量添加商品接口 API请求
taobao.singletreasure.activity.item.batchadd

向活动中批量添加商品优惠
*/
type TaobaoSingletreasureActivityItemBatchaddRequest struct {
    model.Params
    // 系统入参
    _itemDetailInfo   *ItemDetailInfoBatchCreateDto
}

// 初始化TaobaoSingletreasureActivityItemBatchaddRequest对象
func NewTaobaoSingletreasureActivityItemBatchaddRequest() *TaobaoSingletreasureActivityItemBatchaddRequest{
    return &TaobaoSingletreasureActivityItemBatchaddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSingletreasureActivityItemBatchaddRequest) GetApiMethodName() string {
    return "taobao.singletreasure.activity.item.batchadd"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSingletreasureActivityItemBatchaddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemDetailInfo Setter
// 系统入参
func (r *TaobaoSingletreasureActivityItemBatchaddRequest) SetItemDetailInfo(_itemDetailInfo *ItemDetailInfoBatchCreateDto) error {
    r._itemDetailInfo = _itemDetailInfo
    r.Set("item_detail_info", _itemDetailInfo)
    return nil
}

// ItemDetailInfo Getter
func (r TaobaoSingletreasureActivityItemBatchaddRequest) GetItemDetailInfo() *ItemDetailInfoBatchCreateDto {
    return r._itemDetailInfo
}