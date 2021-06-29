package topoaid

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
OAID解密 API请求
taobao.top.oaid.decrypt

解码OAID(Open Addressee ID)，返回收件人信息。
*/
type TaobaoTopOaidDecryptRequest struct {
    model.Params
    // 解密请求列表，最多支持20个。
    _queryList   []ReceiverQuery
}

// 初始化TaobaoTopOaidDecryptRequest对象
func NewTaobaoTopOaidDecryptRequest() *TaobaoTopOaidDecryptRequest{
    return &TaobaoTopOaidDecryptRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTopOaidDecryptRequest) GetApiMethodName() string {
    return "taobao.top.oaid.decrypt"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTopOaidDecryptRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// QueryList Setter
// 解密请求列表，最多支持20个。
func (r *TaobaoTopOaidDecryptRequest) SetQueryList(_queryList []ReceiverQuery) error {
    r._queryList = _queryList
    r.Set("query_list", _queryList)
    return nil
}

// QueryList Getter
func (r TaobaoTopOaidDecryptRequest) GetQueryList() []ReceiverQuery {
    return r._queryList
}
