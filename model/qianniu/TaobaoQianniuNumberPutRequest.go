package qianniu

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
ISV上传数据接口 APIRequest
taobao.qianniu.number.put

ISV提供给卖家使用的业务数据，需要通过这个接口上传到千牛数据中心。
*/
type TaobaoQianniuNumberPutRequest struct {
    model.Params

    // 考虑到稳定性，建议一次卖家最多为200个。
标准json格式的数组构成的字符串。每个元素为{user_id:****,field:"****",value:"****"}分别是用户的userid，数据的名称，以及数据的值。
    data   string 

}

func NewTaobaoQianniuNumberPutRequest() *TaobaoQianniuNumberPutRequest{
    return &TaobaoQianniuNumberPutRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoQianniuNumberPutRequest) GetApiMethodName() string {
    return "taobao.qianniu.number.put"
}

func (r TaobaoQianniuNumberPutRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoQianniuNumberPutRequest) SetData(data string) error {
    r.data = data
    r.Set("data", data)
    return nil
}

func (r TaobaoQianniuNumberPutRequest) GetData() string {
    return r.data
}

