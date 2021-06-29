package tmallsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建主动预约开通条件 API请求
tmall.servicecenter.reservecond.create

1、设置主动预约开通条件
*/
type TmallServicecenterReservecondCreateRequest struct {
    model.Params
    // 主动预约开通条件
    _rocList   []ReserveOpenConditionDTO
}

// 初始化TmallServicecenterReservecondCreateRequest对象
func NewTmallServicecenterReservecondCreateRequest() *TmallServicecenterReservecondCreateRequest{
    return &TmallServicecenterReservecondCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallServicecenterReservecondCreateRequest) GetApiMethodName() string {
    return "tmall.servicecenter.reservecond.create"
}

// IRequest interface 方法, 获取API参数
func (r TmallServicecenterReservecondCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RocList Setter
// 主动预约开通条件
func (r *TmallServicecenterReservecondCreateRequest) SetRocList(_rocList []ReserveOpenConditionDTO) error {
    r._rocList = _rocList
    r.Set("roc_list", _rocList)
    return nil
}

// RocList Getter
func (r TmallServicecenterReservecondCreateRequest) GetRocList() []ReserveOpenConditionDTO {
    return r._rocList
}
