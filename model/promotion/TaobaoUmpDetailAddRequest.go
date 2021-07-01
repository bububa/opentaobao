package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
新增活动详情 API请求
taobao.ump.detail.add

增加活动详情。活动详情里面包括活动的范围（店铺，商品），活动的参数（比如具体的折扣），参与类型（全部，部分，部分不参加）等信息。当参与类型为部分或部分不参加的时候需要和taobao.ump.range.add来配合使用。
*/
type TaobaoUmpDetailAddAPIRequest struct {
    model.Params
    // 增加工具详情
    _actId   int64
    // 活动详情内容，json格式，可以通过ump sdk中的MarketingTool来进行处理
    _content   string
}

// 初始化TaobaoUmpDetailAddAPIRequest对象
func NewTaobaoUmpDetailAddRequest() *TaobaoUmpDetailAddAPIRequest{
    return &TaobaoUmpDetailAddAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoUmpDetailAddAPIRequest) GetApiMethodName() string {
    return "taobao.ump.detail.add"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoUmpDetailAddAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ActId Setter
// 增加工具详情
func (r *TaobaoUmpDetailAddAPIRequest) SetActId(_actId int64) error {
    r._actId = _actId
    r.Set("act_id", _actId)
    return nil
}

// ActId Getter
func (r TaobaoUmpDetailAddAPIRequest) GetActId() int64 {
    return r._actId
}
// Content Setter
// 活动详情内容，json格式，可以通过ump sdk中的MarketingTool来进行处理
func (r *TaobaoUmpDetailAddAPIRequest) SetContent(_content string) error {
    r._content = _content
    r.Set("content", _content)
    return nil
}

// Content Getter
func (r TaobaoUmpDetailAddAPIRequest) GetContent() string {
    return r._content
}
