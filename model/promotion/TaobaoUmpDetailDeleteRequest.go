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
type TaobaoUmpDetailDeleteRequest struct {
    model.Params
    // 活动详情id
    detailId   int64
}

// 初始化TaobaoUmpDetailDeleteRequest对象
func NewTaobaoUmpDetailDeleteRequest() *TaobaoUmpDetailDeleteRequest{
    return &TaobaoUmpDetailDeleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoUmpDetailDeleteRequest) GetApiMethodName() string {
    return "taobao.ump.detail.delete"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoUmpDetailDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DetailId Setter
// 活动详情id
func (r *TaobaoUmpDetailDeleteRequest) SetDetailId(detailId int64) error {
    r.detailId = detailId
    r.Set("detail_id", detailId)
    return nil
}

// DetailId Getter
func (r TaobaoUmpDetailDeleteRequest) GetDetailId() int64 {
    return r.detailId
}
