package smartstore

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
智慧门店区域编码查询 APIRequest
taobao.istore.areas.get

查询标准地址区域代码信息。可以直接参考最新的行政区域代码：
<a href="http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/2016/index.html">http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/2016/index.html</a>
*/
type TaobaoIstoreAreasGetRequest struct {
    model.Params

    // 需返回的字段列表.可选值:Area 结构中的所有字段;多个字段之间用","分隔.如:id,type,name,parent_id,zip.
    fields   string 

}

func NewTaobaoIstoreAreasGetRequest() *TaobaoIstoreAreasGetRequest{
    return &TaobaoIstoreAreasGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoIstoreAreasGetRequest) GetApiMethodName() string {
    return "taobao.istore.areas.get"
}

func (r TaobaoIstoreAreasGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoIstoreAreasGetRequest) SetFields(fields string) error {
    r.fields = fields
    r.Set("fields", fields)
    return nil
}

func (r TaobaoIstoreAreasGetRequest) GetFields() string {
    return r.fields
}

