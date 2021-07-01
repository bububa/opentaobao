package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除活动详情 API请求
taobao.ump.detail.delete

删除活动详情
*/
type TaobaoUmpDetailDeleteAPIRequest struct {
    model.Params
    // 活动详情id
    _detailId   int64
}

// 初始化TaobaoUmpDetailDeleteAPIRequest对象
func NewTaobaoUmpDetailDeleteRequest() *TaobaoUmpDetailDeleteAPIRequest{
    return &TaobaoUmpDetailDeleteAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoUmpDetailDeleteAPIRequest) GetApiMethodName() string {
    return "taobao.ump.detail.delete"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoUmpDetailDeleteAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DetailId Setter
// 活动详情id
func (r *TaobaoUmpDetailDeleteAPIRequest) SetDetailId(_detailId int64) error {
    r._detailId = _detailId
    r.Set("detail_id", _detailId)
    return nil
}

// DetailId Getter
func (r TaobaoUmpDetailDeleteAPIRequest) GetDetailId() int64 {
    return r._detailId
}
