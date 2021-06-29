package topoaid

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
OAID解密 APIRequest
taobao.top.oaid.decrypt

解码OAID(Open Addressee ID)，返回收件人信息。
*/
type TaobaoTopOaidDecryptRequest struct {
    model.Params

    // 解密请求列表，最多支持20个。
    queryList   []ReceiverQuery 

}

func NewTaobaoTopOaidDecryptRequest() *TaobaoTopOaidDecryptRequest{
    return &TaobaoTopOaidDecryptRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTopOaidDecryptRequest) GetApiMethodName() string {
    return "taobao.top.oaid.decrypt"
}

func (r TaobaoTopOaidDecryptRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTopOaidDecryptRequest) SetQueryList(queryList []ReceiverQuery) error {
    r.queryList = queryList
    r.Set("query_list", queryList)
    return nil
}

func (r TaobaoTopOaidDecryptRequest) GetQueryList() []ReceiverQuery {
    return r.queryList
}

