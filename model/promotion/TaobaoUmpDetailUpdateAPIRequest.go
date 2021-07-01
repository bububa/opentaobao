package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
修改活动详情 API请求
taobao.ump.detail.update

更新活动详情
*/
type TaobaoUmpDetailUpdateAPIRequest struct {
    model.Params
    // 活动详情id
    _detailId   int64
    // 活动详情内容，可以通过ump sdk中的MarketingTool来生成这个内容
    _content   string
}

// 初始化TaobaoUmpDetailUpdateAPIRequest对象
func NewTaobaoUmpDetailUpdateRequest() *TaobaoUmpDetailUpdateAPIRequest{
    return &TaobaoUmpDetailUpdateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoUmpDetailUpdateAPIRequest) GetApiMethodName() string {
    return "taobao.ump.detail.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoUmpDetailUpdateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DetailId Setter
// 活动详情id
func (r *TaobaoUmpDetailUpdateAPIRequest) SetDetailId(_detailId int64) error {
    r._detailId = _detailId
    r.Set("detail_id", _detailId)
    return nil
}

// DetailId Getter
func (r TaobaoUmpDetailUpdateAPIRequest) GetDetailId() int64 {
    return r._detailId
}
// Content Setter
// 活动详情内容，可以通过ump sdk中的MarketingTool来生成这个内容
func (r *TaobaoUmpDetailUpdateAPIRequest) SetContent(_content string) error {
    r._content = _content
    r.Set("content", _content)
    return nil
}

// Content Getter
func (r TaobaoUmpDetailUpdateAPIRequest) GetContent() string {
    return r._content
}
