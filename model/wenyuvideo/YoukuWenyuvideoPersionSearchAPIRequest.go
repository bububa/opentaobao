package wenyuvideo

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据人物名称查询人物列表 API请求
youku.wenyuvideo.persion.search

根据人物名称查询人物列表
*/
type YoukuWenyuvideoPersionSearchAPIRequest struct {
    model.Params
    // 人物名字，搜索规则是完全匹配，即只返回同名人物列表
    _personName   string
}

// 初始化YoukuWenyuvideoPersionSearchAPIRequest对象
func NewYoukuWenyuvideoPersionSearchRequest() *YoukuWenyuvideoPersionSearchAPIRequest{
    return &YoukuWenyuvideoPersionSearchAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YoukuWenyuvideoPersionSearchAPIRequest) GetApiMethodName() string {
    return "youku.wenyuvideo.persion.search"
}

// IRequest interface 方法, 获取API参数
func (r YoukuWenyuvideoPersionSearchAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PersonName Setter
// 人物名字，搜索规则是完全匹配，即只返回同名人物列表
func (r *YoukuWenyuvideoPersionSearchAPIRequest) SetPersonName(_personName string) error {
    r._personName = _personName
    r.Set("person_name", _personName)
    return nil
}

// PersonName Getter
func (r YoukuWenyuvideoPersionSearchAPIRequest) GetPersonName() string {
    return r._personName
}
