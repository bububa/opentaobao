package wenyuvideo

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据人物名称查询人物列表 APIRequest
youku.wenyuvideo.persion.search

根据人物名称查询人物列表
*/
type YoukuWenyuvideoPersionSearchRequest struct {
    model.Params

    // 人物名字，搜索规则是完全匹配，即只返回同名人物列表
    personName   string 

}

func NewYoukuWenyuvideoPersionSearchRequest() *YoukuWenyuvideoPersionSearchRequest{
    return &YoukuWenyuvideoPersionSearchRequest{
        Params: model.NewParams(),
    }
}

func (r YoukuWenyuvideoPersionSearchRequest) GetApiMethodName() string {
    return "youku.wenyuvideo.persion.search"
}

func (r YoukuWenyuvideoPersionSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YoukuWenyuvideoPersionSearchRequest) SetPersonName(personName string) error {
    r.personName = personName
    r.Set("person_name", personName)
    return nil
}

func (r YoukuWenyuvideoPersionSearchRequest) GetPersonName() string {
    return r.personName
}

