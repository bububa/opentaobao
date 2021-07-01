package singletreasure

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
更新单品优惠接口 API请求
taobao.singletreasure.activity.item.update

更新单品优惠接口
*/
type TaobaoSingletreasureActivityItemUpdateAPIRequest struct {
    model.Params
    // 修改接口的入参对象
    _itemDetailInfo   *ItemDetailInfoCreateDTO
}

// 初始化TaobaoSingletreasureActivityItemUpdateAPIRequest对象
func NewTaobaoSingletreasureActivityItemUpdateRequest() *TaobaoSingletreasureActivityItemUpdateAPIRequest{
    return &TaobaoSingletreasureActivityItemUpdateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSingletreasureActivityItemUpdateAPIRequest) GetApiMethodName() string {
    return "taobao.singletreasure.activity.item.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSingletreasureActivityItemUpdateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemDetailInfo Setter
// 修改接口的入参对象
func (r *TaobaoSingletreasureActivityItemUpdateAPIRequest) SetItemDetailInfo(_itemDetailInfo *ItemDetailInfoCreateDTO) error {
    r._itemDetailInfo = _itemDetailInfo
    r.Set("item_detail_info", _itemDetailInfo)
    return nil
}

// ItemDetailInfo Getter
func (r TaobaoSingletreasureActivityItemUpdateAPIRequest) GetItemDetailInfo() *ItemDetailInfoCreateDTO {
    return r._itemDetailInfo
}
