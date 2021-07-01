package smartstore

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
智慧门店区域编码查询 API请求
taobao.istore.areas.get

查询标准地址区域代码信息。可以直接参考最新的行政区域代码：
<a href="http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/2016/index.html">http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/2016/index.html</a>
*/
type TaobaoIstoreAreasGetAPIRequest struct {
    model.Params
    // 需返回的字段列表.可选值:Area 结构中的所有字段;多个字段之间用","分隔.如:id,type,name,parent_id,zip.
    _fields   string
}

// 初始化TaobaoIstoreAreasGetAPIRequest对象
func NewTaobaoIstoreAreasGetRequest() *TaobaoIstoreAreasGetAPIRequest{
    return &TaobaoIstoreAreasGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoIstoreAreasGetAPIRequest) GetApiMethodName() string {
    return "taobao.istore.areas.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoIstoreAreasGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Fields Setter
// 需返回的字段列表.可选值:Area 结构中的所有字段;多个字段之间用","分隔.如:id,type,name,parent_id,zip.
func (r *TaobaoIstoreAreasGetAPIRequest) SetFields(_fields string) error {
    r._fields = _fields
    r.Set("fields", _fields)
    return nil
}

// Fields Getter
func (r TaobaoIstoreAreasGetAPIRequest) GetFields() string {
    return r._fields
}
