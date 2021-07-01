package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询活动详情 API请求
taobao.ump.detail.get

查询活动详情
*/
type TaobaoUmpDetailGetAPIRequest struct {
    model.Params
    // 活动详情的id
    _detailId   int64
}

// 初始化TaobaoUmpDetailGetAPIRequest对象
func NewTaobaoUmpDetailGetRequest() *TaobaoUmpDetailGetAPIRequest{
    return &TaobaoUmpDetailGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoUmpDetailGetAPIRequest) GetApiMethodName() string {
    return "taobao.ump.detail.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoUmpDetailGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DetailId Setter
// 活动详情的id
func (r *TaobaoUmpDetailGetAPIRequest) SetDetailId(_detailId int64) error {
    r._detailId = _detailId
    r.Set("detail_id", _detailId)
    return nil
}

// DetailId Getter
func (r TaobaoUmpDetailGetAPIRequest) GetDetailId() int64 {
    return r._detailId
}
