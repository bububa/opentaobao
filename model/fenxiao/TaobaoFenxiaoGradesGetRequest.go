package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
分销商等级查询 APIRequest
taobao.fenxiao.grades.get

根据供应商ID，查询他的分销商等级信息
*/
type TaobaoFenxiaoGradesGetRequest struct {
    model.Params

}

func NewTaobaoFenxiaoGradesGetRequest() *TaobaoFenxiaoGradesGetRequest{
    return &TaobaoFenxiaoGradesGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoFenxiaoGradesGetRequest) GetApiMethodName() string {
    return "taobao.fenxiao.grades.get"
}

func (r TaobaoFenxiaoGradesGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


