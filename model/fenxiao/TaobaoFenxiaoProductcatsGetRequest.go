package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询产品线列表 APIRequest
taobao.fenxiao.productcats.get

查询供应商的所有产品线数据。根据登陆用户来查询，不需要其他入参
*/
type TaobaoFenxiaoProductcatsGetRequest struct {
    model.Params

    // 返回字段列表
    fields   string 

}

func NewTaobaoFenxiaoProductcatsGetRequest() *TaobaoFenxiaoProductcatsGetRequest{
    return &TaobaoFenxiaoProductcatsGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoFenxiaoProductcatsGetRequest) GetApiMethodName() string {
    return "taobao.fenxiao.productcats.get"
}

func (r TaobaoFenxiaoProductcatsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoFenxiaoProductcatsGetRequest) SetFields(fields string) error {
    r.fields = fields
    r.Set("fields", fields)
    return nil
}

func (r TaobaoFenxiaoProductcatsGetRequest) GetFields() string {
    return r.fields
}

