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
type TmallServicecenterReservecondUpdateAPIRequest struct {
    model.Params
    // 主动预约开通条件
    _rocList   []ReserveOpenConditionDto
}

// 初始化TmallServicecenterReservecondUpdateAPIRequest对象
func NewTmallServicecenterReservecondUpdateRequest() *TmallServicecenterReservecondUpdateAPIRequest{
    return &TmallServicecenterReservecondUpdateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallServicecenterReservecondUpdateAPIRequest) GetApiMethodName() string {
    return "tmall.servicecenter.reservecond.update"
}

// IRequest interface 方法, 获取API参数
func (r TmallServicecenterReservecondUpdateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RocList Setter
// 主动预约开通条件
func (r *TmallServicecenterReservecondUpdateAPIRequest) SetRocList(_rocList []ReserveOpenConditionDto) error {
    r._rocList = _rocList
    r.Set("roc_list", _rocList)
    return nil
}

// RocList Getter
func (r TmallServicecenterReservecondUpdateAPIRequest) GetRocList() []ReserveOpenConditionDto {
    return r._rocList
}
