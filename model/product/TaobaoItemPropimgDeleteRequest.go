package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除属性图片 APIRequest
taobao.item.propimg.delete

删除propimg_id 所指定的商品属性图片 <br/>传入的num_iid所对应的商品必须属于当前会话的用户 <br/>propimg_id对应的属性图片需要属于num_iid对应的商品
*/
type TaobaoItemPropimgDeleteRequest struct {
    model.Params

    // 商品属性图片ID
    id   int64 

    // 商品数字ID，必选
    numIid   int64 

}

func NewTaobaoItemPropimgDeleteRequest() *TaobaoItemPropimgDeleteRequest{
    return &TaobaoItemPropimgDeleteRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoItemPropimgDeleteRequest) GetApiMethodName() string {
    return "taobao.item.propimg.delete"
}

func (r TaobaoItemPropimgDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoItemPropimgDeleteRequest) SetId(id int64) error {
    r.id = id
    r.Set("id", id)
    return nil
}

func (r TaobaoItemPropimgDeleteRequest) GetId() int64 {
    return r.id
}

func (r *TaobaoItemPropimgDeleteRequest) SetNumIid(numIid int64) error {
    r.numIid = numIid
    r.Set("num_iid", numIid)
    return nil
}

func (r TaobaoItemPropimgDeleteRequest) GetNumIid() int64 {
    return r.numIid
}

