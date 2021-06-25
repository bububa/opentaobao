package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
新增活动详情 APIRequest
taobao.ump.detail.add

增加活动详情。活动详情里面包括活动的范围（店铺，商品），活动的参数（比如具体的折扣），参与类型（全部，部分，部分不参加）等信息。当参与类型为部分或部分不参加的时候需要和taobao.ump.range.add来配合使用。
*/
type TaobaoUmpDetailAddRequest struct {
    model.Params

    // 增加工具详情
    actId   int64 

    // 活动详情内容，json格式，可以通过ump sdk中的MarketingTool来进行处理
    content   string 

}

func NewTaobaoUmpDetailAddRequest() *TaobaoUmpDetailAddRequest{
    return &TaobaoUmpDetailAddRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoUmpDetailAddRequest) GetApiMethodName() string {
    return "taobao.ump.detail.add"
}

func (r TaobaoUmpDetailAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoUmpDetailAddRequest) SetActId(actId int64) error {
    r.actId = actId
    r.Set("act_id", actId)
    return nil
}

func (r TaobaoUmpDetailAddRequest) GetActId() int64 {
    return r.actId
}

func (r *TaobaoUmpDetailAddRequest) SetContent(content string) error {
    r.content = content
    r.Set("content", content)
    return nil
}

func (r TaobaoUmpDetailAddRequest) GetContent() string {
    return r.content
}

