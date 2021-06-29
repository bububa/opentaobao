package user

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
Open Account索引查询 API请求
taobao.open.account.index.find

Open Account索引查询
*/
type TaobaoOpenAccountIndexFindRequest struct {
    model.Params
    // int MOBILE         = 1;int EMAIL          = 2;int ISV_ACCOUNT_ID = 3;int LOGIN_ID       = 4;int OPEN_ID        = 5;
    _indexType   int64
    // 具体值，当索引类型是 OPEN_ID 是，格式为 oauthPlatform|openId，即使用竖线分隔的组合值
    _indexValue   string
}

// 初始化TaobaoOpenAccountIndexFindRequest对象
func NewTaobaoOpenAccountIndexFindRequest() *TaobaoOpenAccountIndexFindRequest{
    return &TaobaoOpenAccountIndexFindRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOpenAccountIndexFindRequest) GetApiMethodName() string {
    return "taobao.open.account.index.find"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOpenAccountIndexFindRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// IndexType Setter
// int MOBILE         = 1;int EMAIL          = 2;int ISV_ACCOUNT_ID = 3;int LOGIN_ID       = 4;int OPEN_ID        = 5;
func (r *TaobaoOpenAccountIndexFindRequest) SetIndexType(_indexType int64) error {
    r._indexType = _indexType
    r.Set("index_type", _indexType)
    return nil
}

// IndexType Getter
func (r TaobaoOpenAccountIndexFindRequest) GetIndexType() int64 {
    return r._indexType
}
// IndexValue Setter
// 具体值，当索引类型是 OPEN_ID 是，格式为 oauthPlatform|openId，即使用竖线分隔的组合值
func (r *TaobaoOpenAccountIndexFindRequest) SetIndexValue(_indexValue string) error {
    r._indexValue = _indexValue
    r.Set("index_value", _indexValue)
    return nil
}

// IndexValue Getter
func (r TaobaoOpenAccountIndexFindRequest) GetIndexValue() string {
    return r._indexValue
}
