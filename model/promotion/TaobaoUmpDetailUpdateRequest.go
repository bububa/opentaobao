package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
修改活动详情 APIRequest
taobao.ump.detail.update

更新活动详情
*/
type TaobaoUmpDetailUpdateRequest struct {
    model.Params

    // 活动详情id
    detailId   int64 

    // 活动详情内容，可以通过ump sdk中的MarketingTool来生成这个内容
    content   string 

}

func NewTaobaoUmpDetailUpdateRequest() *TaobaoUmpDetailUpdateRequest{
    return &TaobaoUmpDetailUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoUmpDetailUpdateRequest) GetApiMethodName() string {
    return "taobao.ump.detail.update"
}

func (r TaobaoUmpDetailUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoUmpDetailUpdateRequest) SetDetailId(detailId int64) error {
    r.detailId = detailId
    r.Set("detail_id", detailId)
    return nil
}

func (r TaobaoUmpDetailUpdateRequest) GetDetailId() int64 {
    return r.detailId
}

func (r *TaobaoUmpDetailUpdateRequest) SetContent(content string) error {
    r.content = content
    r.Set("content", content)
    return nil
}

func (r TaobaoUmpDetailUpdateRequest) GetContent() string {
    return r.content
}

