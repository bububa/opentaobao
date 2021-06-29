package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商品关联属性图 API请求
taobao.item.joint.propimg

* 关联一张商品属性图片到num_iid指定的商品中<br/>    * 传入的num_iid所对应的商品必须属于当前会话的用户<br/>    * 图片的属性必须要是颜色的属性，这个在前台显示的时候需要和sku进行关联的<br/>    * 商品图片关联在卖家身份和图片来源上的限制，卖家要是B卖家或订购了多图服务才能关联图片，并且图片要来自于卖家自己的图片空间才行<br/>    * 商品图片数量有限制。不管是上传的图片还是关联的图片，他们的总数不能超过一定限额，最多不能超过24张（每个颜色属性都有一张）
*/
type TaobaoItemJointPropimgRequest struct {
    model.Params
    // 属性列表。调用taobao.itemprops.get获取，属性必须是颜色属性，格式:pid:vid。
    properties   string
    // 属性图片ID。如果是新增不需要填写
    id   int64
    // 图片地址(传入图片相对地址即可,即不需包含 http://img02.taobao.net/bao/uploaded )
    picPath   string
    // 商品数字ID，必选
    numIid   int64
    // 图片序号
    position   int64
}

// 初始化TaobaoItemJointPropimgRequest对象
func NewTaobaoItemJointPropimgRequest() *TaobaoItemJointPropimgRequest{
    return &TaobaoItemJointPropimgRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoItemJointPropimgRequest) GetApiMethodName() string {
    return "taobao.item.joint.propimg"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoItemJointPropimgRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Properties Setter
// 属性列表。调用taobao.itemprops.get获取，属性必须是颜色属性，格式:pid:vid。
func (r *TaobaoItemJointPropimgRequest) SetProperties(properties string) error {
    r.properties = properties
    r.Set("properties", properties)
    return nil
}

// Properties Getter
func (r TaobaoItemJointPropimgRequest) GetProperties() string {
    return r.properties
}
// Id Setter
// 属性图片ID。如果是新增不需要填写
func (r *TaobaoItemJointPropimgRequest) SetId(id int64) error {
    r.id = id
    r.Set("id", id)
    return nil
}

// Id Getter
func (r TaobaoItemJointPropimgRequest) GetId() int64 {
    return r.id
}
// PicPath Setter
// 图片地址(传入图片相对地址即可,即不需包含 http://img02.taobao.net/bao/uploaded )
func (r *TaobaoItemJointPropimgRequest) SetPicPath(picPath string) error {
    r.picPath = picPath
    r.Set("pic_path", picPath)
    return nil
}

// PicPath Getter
func (r TaobaoItemJointPropimgRequest) GetPicPath() string {
    return r.picPath
}
// NumIid Setter
// 商品数字ID，必选
func (r *TaobaoItemJointPropimgRequest) SetNumIid(numIid int64) error {
    r.numIid = numIid
    r.Set("num_iid", numIid)
    return nil
}

// NumIid Getter
func (r TaobaoItemJointPropimgRequest) GetNumIid() int64 {
    return r.numIid
}
// Position Setter
// 图片序号
func (r *TaobaoItemJointPropimgRequest) SetPosition(position int64) error {
    r.position = position
    r.Set("position", position)
    return nil
}

// Position Getter
func (r TaobaoItemJointPropimgRequest) GetPosition() int64 {
    return r.position
}
