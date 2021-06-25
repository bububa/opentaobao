package category

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取标准类目属性值 APIRequest
taobao.itempropvalues.get

获取标准类目属性值
*/
type TaobaoItempropvaluesGetRequest struct {
    model.Params

    // 需要返回的字段。目前支持有：cid,pid,prop_name,vid,name,name_alias,status,sort_order
    fields   []String 

    // 叶子类目ID ,通过taobao.itemcats.get获得叶子类目ID
    cid   int64 

    // 属性和属性值 id串，格式例如(pid1;pid2)或(pid1:vid1;pid2:vid2)或(pid1;pid2:vid2)
    pvs   string 

    // 获取类目的类型：1代表集市、2代表天猫
    type   int64 

    // 属性的Key，支持多条，以“,”分隔
    attrKeys   []String 

}

func NewTaobaoItempropvaluesGetRequest() *TaobaoItempropvaluesGetRequest{
    return &TaobaoItempropvaluesGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoItempropvaluesGetRequest) GetApiMethodName() string {
    return "taobao.itempropvalues.get"
}

func (r TaobaoItempropvaluesGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoItempropvaluesGetRequest) SetFields(fields []String) error {
    r.fields = fields
    r.Set("fields", fields)
    return nil
}

func (r TaobaoItempropvaluesGetRequest) GetFields() []String {
    return r.fields
}

func (r *TaobaoItempropvaluesGetRequest) SetCid(cid int64) error {
    r.cid = cid
    r.Set("cid", cid)
    return nil
}

func (r TaobaoItempropvaluesGetRequest) GetCid() int64 {
    return r.cid
}

func (r *TaobaoItempropvaluesGetRequest) SetPvs(pvs string) error {
    r.pvs = pvs
    r.Set("pvs", pvs)
    return nil
}

func (r TaobaoItempropvaluesGetRequest) GetPvs() string {
    return r.pvs
}

func (r *TaobaoItempropvaluesGetRequest) SetType(type int64) error {
    r.type = type
    r.Set("type", type)
    return nil
}

func (r TaobaoItempropvaluesGetRequest) GetType() int64 {
    return r.type
}

func (r *TaobaoItempropvaluesGetRequest) SetAttrKeys(attrKeys []String) error {
    r.attrKeys = attrKeys
    r.Set("attr_keys", attrKeys)
    return nil
}

func (r TaobaoItempropvaluesGetRequest) GetAttrKeys() []String {
    return r.attrKeys
}

