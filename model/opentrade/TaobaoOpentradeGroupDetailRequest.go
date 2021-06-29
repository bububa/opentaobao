package opentrade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
组团购场景查询团详情 API请求
taobao.opentrade.group.detail

组团购场景下，查询团详情
*/
type TaobaoOpentradeGroupDetailRequest struct {
    model.Params
    // 团id
    _groupId   int64
}

// 初始化TaobaoOpentradeGroupDetailRequest对象
func NewTaobaoOpentradeGroupDetailRequest() *TaobaoOpentradeGroupDetailRequest{
    return &TaobaoOpentradeGroupDetailRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOpentradeGroupDetailRequest) GetApiMethodName() string {
    return "taobao.opentrade.group.detail"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOpentradeGroupDetailRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// GroupId Setter
// 团id
func (r *TaobaoOpentradeGroupDetailRequest) SetGroupId(_groupId int64) error {
    r._groupId = _groupId
    r.Set("group_id", _groupId)
    return nil
}

// GroupId Getter
func (r TaobaoOpentradeGroupDetailRequest) GetGroupId() int64 {
    return r._groupId
}
