package user

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
open account数据搜索 API请求
taobao.open.account.search

open account数据搜索
*/
type TaobaoOpenAccountSearchRequest struct {
    model.Params
    // 基于阿里云OpenSearch服务，openSearch查询语法:https://help.aliyun.com/document_detail/29157.html，搜索服务能够基于id，loginId，mobile，email，isvAccountId，display_name,create_app_key,做搜索查询，示例中mobile可以基于模糊搜素，搜索135的账号，该搜索是分页返回，start表示开始行，hit表示一页返回值，最大500
    _query   string
}

// 初始化TaobaoOpenAccountSearchRequest对象
func NewTaobaoOpenAccountSearchRequest() *TaobaoOpenAccountSearchRequest{
    return &TaobaoOpenAccountSearchRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOpenAccountSearchRequest) GetApiMethodName() string {
    return "taobao.open.account.search"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOpenAccountSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Query Setter
// 基于阿里云OpenSearch服务，openSearch查询语法:https://help.aliyun.com/document_detail/29157.html，搜索服务能够基于id，loginId，mobile，email，isvAccountId，display_name,create_app_key,做搜索查询，示例中mobile可以基于模糊搜素，搜索135的账号，该搜索是分页返回，start表示开始行，hit表示一页返回值，最大500
func (r *TaobaoOpenAccountSearchRequest) SetQuery(_query string) error {
    r._query = _query
    r.Set("query", _query)
    return nil
}

// Query Getter
func (r TaobaoOpenAccountSearchRequest) GetQuery() string {
    return r._query
}
