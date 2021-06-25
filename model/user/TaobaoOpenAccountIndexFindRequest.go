package user

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
Open Account索引查询 APIRequest
taobao.open.account.index.find

Open Account索引查询
*/
type TaobaoOpenAccountIndexFindRequest struct {
    model.Params

    // int MOBILE         = 1;int EMAIL          = 2;int ISV_ACCOUNT_ID = 3;int LOGIN_ID       = 4;int OPEN_ID        = 5;
    indexType   int64 

    // 具体值，当索引类型是 OPEN_ID 是，格式为 oauthPlatform|openId，即使用竖线分隔的组合值
    indexValue   string 

}

func NewTaobaoOpenAccountIndexFindRequest() *TaobaoOpenAccountIndexFindRequest{
    return &TaobaoOpenAccountIndexFindRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOpenAccountIndexFindRequest) GetApiMethodName() string {
    return "taobao.open.account.index.find"
}

func (r TaobaoOpenAccountIndexFindRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOpenAccountIndexFindRequest) SetIndexType(indexType int64) error {
    r.indexType = indexType
    r.Set("index_type", indexType)
    return nil
}

func (r TaobaoOpenAccountIndexFindRequest) GetIndexType() int64 {
    return r.indexType
}

func (r *TaobaoOpenAccountIndexFindRequest) SetIndexValue(indexValue string) error {
    r.indexValue = indexValue
    r.Set("index_value", indexValue)
    return nil
}

func (r TaobaoOpenAccountIndexFindRequest) GetIndexValue() string {
    return r.indexValue
}

