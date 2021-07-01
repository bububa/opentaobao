package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询产品线列表 API请求
taobao.fenxiao.productcats.get

查询供应商的所有产品线数据。根据登陆用户来查询，不需要其他入参
*/
type TaobaoFenxiaoProductcatsGetAPIRequest struct {
    model.Params
    // 返回字段列表
    _fields   string
}

// 初始化TaobaoFenxiaoProductcatsGetAPIRequest对象
func NewTaobaoFenxiaoProductcatsGetRequest() *TaobaoFenxiaoProductcatsGetAPIRequest{
    return &TaobaoFenxiaoProductcatsGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFenxiaoProductcatsGetAPIRequest) GetApiMethodName() string {
    return "taobao.fenxiao.productcats.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFenxiaoProductcatsGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Fields Setter
// 返回字段列表
func (r *TaobaoFenxiaoProductcatsGetAPIRequest) SetFields(_fields string) error {
    r._fields = _fields
    r.Set("fields", _fields)
    return nil
}

// Fields Getter
func (r TaobaoFenxiaoProductcatsGetAPIRequest) GetFields() string {
    return r._fields
}
