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
type TaobaoSingletreasureActivityItemUpdateRequest struct {
    model.Params
    // 修改接口的入参对象
    _itemDetailInfo   *ItemDetailInfoCreateDTO
}

// 初始化TaobaoSingletreasureActivityItemUpdateRequest对象
func NewTaobaoSingletreasureActivityItemUpdateRequest() *TaobaoSingletreasureActivityItemUpdateRequest{
    return &TaobaoSingletreasureActivityItemUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSingletreasureActivityItemUpdateRequest) GetApiMethodName() string {
    return "taobao.singletreasure.activity.item.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSingletreasureActivityItemUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemDetailInfo Setter
// 修改接口的入参对象
func (r *TaobaoSingletreasureActivityItemUpdateRequest) SetItemDetailInfo(_itemDetailInfo *ItemDetailInfoCreateDTO) error {
    r._itemDetailInfo = _itemDetailInfo
    r.Set("item_detail_info", _itemDetailInfo)
    return nil
}

// ItemDetailInfo Getter
func (r TaobaoSingletreasureActivityItemUpdateRequest) GetItemDetailInfo() *ItemDetailInfoCreateDTO {
    return r._itemDetailInfo
}
