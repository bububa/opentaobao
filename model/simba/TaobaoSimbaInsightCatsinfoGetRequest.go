package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
类目信息获取 API请求
taobao.simba.insight.catsinfo.get

获取类目信息，此接口既提供所有顶级类目的查询，又提供给定类目id自身信息和子类目信息的查询，所以可以根据此接口逐层获取所有的类目信息
*/
type TaobaoSimbaInsightCatsinfoGetRequest struct {
    model.Params
    // 表示请求的类型：0表示请求所有顶级类目的信息，这时可以忽略第二个参数，1表示获取给定的类目id的详细信息，2表示获取给定类目id的所有子类目的详细信息
    _type   int64
    // 需要查询的类目id
    _categoryIdList   []string
}

// 初始化TaobaoSimbaInsightCatsinfoGetRequest对象
func NewTaobaoSimbaInsightCatsinfoGetRequest() *TaobaoSimbaInsightCatsinfoGetRequest{
    return &TaobaoSimbaInsightCatsinfoGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaInsightCatsinfoGetRequest) GetApiMethodName() string {
    return "taobao.simba.insight.catsinfo.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaInsightCatsinfoGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Type Setter
// 表示请求的类型：0表示请求所有顶级类目的信息，这时可以忽略第二个参数，1表示获取给定的类目id的详细信息，2表示获取给定类目id的所有子类目的详细信息
func (r *TaobaoSimbaInsightCatsinfoGetRequest) SetType(_type int64) error {
    r._type = _type
    r.Set("type", _type)
    return nil
}

// Type Getter
func (r TaobaoSimbaInsightCatsinfoGetRequest) GetType() int64 {
    return r._type
}
// CategoryIdList Setter
// 需要查询的类目id
func (r *TaobaoSimbaInsightCatsinfoGetRequest) SetCategoryIdList(_categoryIdList []string) error {
    r._categoryIdList = _categoryIdList
    r.Set("category_id_list", _categoryIdList)
    return nil
}

// CategoryIdList Getter
func (r TaobaoSimbaInsightCatsinfoGetRequest) GetCategoryIdList() []string {
    return r._categoryIdList
}
