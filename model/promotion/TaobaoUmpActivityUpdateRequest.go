package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
修改活动信息 APIRequest
taobao.ump.activity.update

修改营销活动
*/
type TaobaoUmpActivityUpdateRequest struct {
    model.Params

    // 活动id
    actId   int64 

    // 营销活动内容，json格式，通过ump sdk 的marketingTool来生成
    content   string 

}

func NewTaobaoUmpActivityUpdateRequest() *TaobaoUmpActivityUpdateRequest{
    return &TaobaoUmpActivityUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoUmpActivityUpdateRequest) GetApiMethodName() string {
    return "taobao.ump.activity.update"
}

func (r TaobaoUmpActivityUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoUmpActivityUpdateRequest) SetActId(actId int64) error {
    r.actId = actId
    r.Set("act_id", actId)
    return nil
}

func (r TaobaoUmpActivityUpdateRequest) GetActId() int64 {
    return r.actId
}

func (r *TaobaoUmpActivityUpdateRequest) SetContent(content string) error {
    r.content = content
    r.Set("content", content)
    return nil
}

func (r TaobaoUmpActivityUpdateRequest) GetContent() string {
    return r.content
}

