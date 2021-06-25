package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
新增优惠活动 APIRequest
taobao.ump.activity.add

新增优惠活动。设置优惠活动的基本信息，比如活动时间，活动针对的对象（可以是满足某些条件的买家）
*/
type TaobaoUmpActivityAddRequest struct {
    model.Params

    // 工具id
    toolId   int64 

    // 活动内容，通过ump sdk里面的MarkeitngTool来生成，name必须属于“营销标签词库”——https://huodong.m.taobao.com/api/data/v2/5fe5e737d3314fa2973297f86f7bff3a.js?file=5fe5e737d3314fa2973297f86f7bff3a.js中的word值中的一种。
    content   string 

}

func NewTaobaoUmpActivityAddRequest() *TaobaoUmpActivityAddRequest{
    return &TaobaoUmpActivityAddRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoUmpActivityAddRequest) GetApiMethodName() string {
    return "taobao.ump.activity.add"
}

func (r TaobaoUmpActivityAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoUmpActivityAddRequest) SetToolId(toolId int64) error {
    r.toolId = toolId
    r.Set("tool_id", toolId)
    return nil
}

func (r TaobaoUmpActivityAddRequest) GetToolId() int64 {
    return r.toolId
}

func (r *TaobaoUmpActivityAddRequest) SetContent(content string) error {
    r.content = content
    r.Set("content", content)
    return nil
}

func (r TaobaoUmpActivityAddRequest) GetContent() string {
    return r.content
}

