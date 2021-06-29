package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
分销商等级查询 API请求
taobao.fenxiao.grades.get

根据供应商ID，查询他的分销商等级信息
*/
type TaobaoFenxiaoGradesGetRequest struct {
    model.Params
}

// 初始化TaobaoFenxiaoGradesGetRequest对象
func NewTaobaoFenxiaoGradesGetRequest() *TaobaoFenxiaoGradesGetRequest{
    return &TaobaoFenxiaoGradesGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFenxiaoGradesGetRequest) GetApiMethodName() string {
    return "taobao.fenxiao.grades.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFenxiaoGradesGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
