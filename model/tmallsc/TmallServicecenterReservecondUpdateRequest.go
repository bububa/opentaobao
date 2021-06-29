package tmallsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
主动预约条件更新 API请求
tmall.servicecenter.reservecond.update

1、设置主动预约开通条件
*/
type TmallServicecenterReservecondUpdateRequest struct {
    model.Params
    // 主动预约开通条件
    _rocList   []ReserveOpenConditionDTO
}

// 初始化TmallServicecenterReservecondUpdateRequest对象
func NewTmallServicecenterReservecondUpdateRequest() *TmallServicecenterReservecondUpdateRequest{
    return &TmallServicecenterReservecondUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallServicecenterReservecondUpdateRequest) GetApiMethodName() string {
    return "tmall.servicecenter.reservecond.update"
}

// IRequest interface 方法, 获取API参数
func (r TmallServicecenterReservecondUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RocList Setter
// 主动预约开通条件
func (r *TmallServicecenterReservecondUpdateRequest) SetRocList(_rocList []ReserveOpenConditionDTO) error {
    r._rocList = _rocList
    r.Set("roc_list", _rocList)
    return nil
}

// RocList Getter
func (r TmallServicecenterReservecondUpdateRequest) GetRocList() []ReserveOpenConditionDTO {
    return r._rocList
}
