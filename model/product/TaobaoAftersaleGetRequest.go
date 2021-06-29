package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询用户售后服务模板 API请求
taobao.aftersale.get

查询用户设置的售后服务模板，仅返回标题和id
*/
type TaobaoAftersaleGetRequest struct {
    model.Params
}

// 初始化TaobaoAftersaleGetRequest对象
func NewTaobaoAftersaleGetRequest() *TaobaoAftersaleGetRequest{
    return &TaobaoAftersaleGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAftersaleGetRequest) GetApiMethodName() string {
    return "taobao.aftersale.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAftersaleGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
