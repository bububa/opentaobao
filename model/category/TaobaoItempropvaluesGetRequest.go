package category

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取标准类目属性值 API请求
taobao.itempropvalues.get

获取标准类目属性值
*/
type TaobaoItempropvaluesGetRequest struct {
    model.Params
    // 需要返回的字段。目前支持有：cid,pid,prop_name,vid,name,name_alias,status,sort_order
    fields   []string
    // 叶子类目ID ,通过taobao.itemcats.get获得叶子类目ID
    cid   int64
    // 属性和属性值 id串，格式例如(pid1;pid2)或(pid1:vid1;pid2:vid2)或(pid1;pid2:vid2)
    pvs   string
    // 获取类目的类型：1代表集市、2代表天猫
    type   int64
    // 属性的Key，支持多条，以“,”分隔
    attrKeys   []string
}

// 初始化TaobaoItempropvaluesGetRequest对象
func NewTaobaoItempropvaluesGetRequest() *TaobaoItempropvaluesGetRequest{
    return &TaobaoItempropvaluesGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoItempropvaluesGetRequest) GetApiMethodName() string {
    return "taobao.itempropvalues.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoItempropvaluesGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Fields Setter
// 需要返回的字段。目前支持有：cid,pid,prop_name,vid,name,name_alias,status,sort_order
func (r *TaobaoItempropvaluesGetRequest) SetFields(fields []string) error {
    r.fields = fields
    r.Set("fields", fields)
    return nil
}

// Fields Getter
func (r TaobaoItempropvaluesGetRequest) GetFields() []string {
    return r.fields
}
// Cid Setter
// 叶子类目ID ,通过taobao.itemcats.get获得叶子类目ID
func (r *TaobaoItempropvaluesGetRequest) SetCid(cid int64) error {
    r.cid = cid
    r.Set("cid", cid)
    return nil
}

// Cid Getter
func (r TaobaoItempropvaluesGetRequest) GetCid() int64 {
    return r.cid
}
// Pvs Setter
// 属性和属性值 id串，格式例如(pid1;pid2)或(pid1:vid1;pid2:vid2)或(pid1;pid2:vid2)
func (r *TaobaoItempropvaluesGetRequest) SetPvs(pvs string) error {
    r.pvs = pvs
    r.Set("pvs", pvs)
    return nil
}

// Pvs Getter
func (r TaobaoItempropvaluesGetRequest) GetPvs() string {
    return r.pvs
}
// Type Setter
// 获取类目的类型：1代表集市、2代表天猫
func (r *TaobaoItempropvaluesGetRequest) SetType(type int64) error {
    r.type = type
    r.Set("type", type)
    return nil
}

// Type Getter
func (r TaobaoItempropvaluesGetRequest) GetType() int64 {
    return r.type
}
// AttrKeys Setter
// 属性的Key，支持多条，以“,”分隔
func (r *TaobaoItempropvaluesGetRequest) SetAttrKeys(attrKeys []string) error {
    r.attrKeys = attrKeys
    r.Set("attr_keys", attrKeys)
    return nil
}

// AttrKeys Getter
func (r TaobaoItempropvaluesGetRequest) GetAttrKeys() []string {
    return r.attrKeys
}
