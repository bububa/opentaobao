package iot

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取对话流列表 APIRequest
taobao.ailab.aicloud.top.feedlist.get

获取指定应用的对话流信息
*/
type TaobaoAilabAicloudTopFeedlistGetRequest struct {
    model.Params

    // 用户信息
    param0   *OpenBaseInfo 

    // 设备id
    param1   string 

    // 最后一条对话的key
    param2   string 

    // 单页的条目数，注意，是String类型！
    param3   string 

}

func NewTaobaoAilabAicloudTopFeedlistGetRequest() *TaobaoAilabAicloudTopFeedlistGetRequest{
    return &TaobaoAilabAicloudTopFeedlistGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAilabAicloudTopFeedlistGetRequest) GetApiMethodName() string {
    return "taobao.ailab.aicloud.top.feedlist.get"
}

func (r TaobaoAilabAicloudTopFeedlistGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAilabAicloudTopFeedlistGetRequest) SetParam0(param0 *OpenBaseInfo) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

func (r TaobaoAilabAicloudTopFeedlistGetRequest) GetParam0() *OpenBaseInfo {
    return r.param0
}

func (r *TaobaoAilabAicloudTopFeedlistGetRequest) SetParam1(param1 string) error {
    r.param1 = param1
    r.Set("param1", param1)
    return nil
}

func (r TaobaoAilabAicloudTopFeedlistGetRequest) GetParam1() string {
    return r.param1
}

func (r *TaobaoAilabAicloudTopFeedlistGetRequest) SetParam2(param2 string) error {
    r.param2 = param2
    r.Set("param2", param2)
    return nil
}

func (r TaobaoAilabAicloudTopFeedlistGetRequest) GetParam2() string {
    return r.param2
}

func (r *TaobaoAilabAicloudTopFeedlistGetRequest) SetParam3(param3 string) error {
    r.param3 = param3
    r.Set("param3", param3)
    return nil
}

func (r TaobaoAilabAicloudTopFeedlistGetRequest) GetParam3() string {
    return r.param3
}

