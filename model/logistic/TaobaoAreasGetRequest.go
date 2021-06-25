package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询地址区域 APIRequest
taobao.areas.get

查询标准地址区域代码信息。可以直接参考最新的行政区域代码：
<a href="http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/" target="_blank"> http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/</a>
*/
type TaobaoAreasGetRequest struct {
    model.Params

    // 需返回的字段列表.可选值:Area 结构中的所有字段;多个字段之间用","分隔.如:id,type,name,parent_id,zip.
    fields   string 

}

func NewTaobaoAreasGetRequest() *TaobaoAreasGetRequest{
    return &TaobaoAreasGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAreasGetRequest) GetApiMethodName() string {
    return "taobao.areas.get"
}

func (r TaobaoAreasGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAreasGetRequest) SetFields(fields string) error {
    r.fields = fields
    r.Set("fields", fields)
    return nil
}

func (r TaobaoAreasGetRequest) GetFields() string {
    return r.fields
}

