package category

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取后台供卖家发布商品的标准商品类目 APIRequest
taobao.itemcats.get

获取后台供卖家发布商品的标准商品类目。
*/
type TaobaoItemcatsGetRequest struct {
    model.Params

    // 商品所属类目ID列表，用半角逗号(,)分隔 例如:(18957,19562,) (cids、parent_cid至少传一个)
    cids   []int64 

    // 需要返回的字段列表，见ItemCat，默认返回：cid,parent_cid,name,is_parent；增量类目信息,根据fields传入的参数返回相应的结果。 features字段： 1、如果存在attr_key=freeze表示该类目被冻结了，attr_value=0,5，value可能存在2个值（也可能只有1个），用逗号分割，0表示禁编辑，5表示禁止发布
    fields   []string 

    // 父商品类目 id，0表示根节点, 传输该参数返回所有子类目。 (cids、parent_cid至少传一个)
    parentCid   int64 

}

func NewTaobaoItemcatsGetRequest() *TaobaoItemcatsGetRequest{
    return &TaobaoItemcatsGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoItemcatsGetRequest) GetApiMethodName() string {
    return "taobao.itemcats.get"
}

func (r TaobaoItemcatsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoItemcatsGetRequest) SetCids(cids []int64) error {
    r.cids = cids
    r.Set("cids", cids)
    return nil
}

func (r TaobaoItemcatsGetRequest) GetCids() []int64 {
    return r.cids
}

func (r *TaobaoItemcatsGetRequest) SetFields(fields []string) error {
    r.fields = fields
    r.Set("fields", fields)
    return nil
}

func (r TaobaoItemcatsGetRequest) GetFields() []string {
    return r.fields
}

func (r *TaobaoItemcatsGetRequest) SetParentCid(parentCid int64) error {
    r.parentCid = parentCid
    r.Set("parent_cid", parentCid)
    return nil
}

func (r TaobaoItemcatsGetRequest) GetParentCid() int64 {
    return r.parentCid
}

