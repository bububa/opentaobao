package user

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
open account数据搜索 APIRequest
taobao.open.account.search

open account数据搜索
*/
type TaobaoOpenAccountSearchRequest struct {
    model.Params

    // 基于阿里云OpenSearch服务，openSearch查询语法:https://help.aliyun.com/document_detail/29157.html，搜索服务能够基于id，loginId，mobile，email，isvAccountId，display_name,create_app_key,做搜索查询，示例中mobile可以基于模糊搜素，搜索135的账号，该搜索是分页返回，start表示开始行，hit表示一页返回值，最大500
    query   string 

}

func NewTaobaoOpenAccountSearchRequest() *TaobaoOpenAccountSearchRequest{
    return &TaobaoOpenAccountSearchRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOpenAccountSearchRequest) GetApiMethodName() string {
    return "taobao.open.account.search"
}

func (r TaobaoOpenAccountSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOpenAccountSearchRequest) SetQuery(query string) error {
    r.query = query
    r.Set("query", query)
    return nil
}

func (r TaobaoOpenAccountSearchRequest) GetQuery() string {
    return r.query
}

