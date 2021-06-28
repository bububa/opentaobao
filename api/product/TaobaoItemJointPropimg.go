package product

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/product"
)

/* 
商品关联属性图 
taobao.item.joint.propimg

* 关联一张商品属性图片到num_iid指定的商品中<br/>    * 传入的num_iid所对应的商品必须属于当前会话的用户<br/>    * 图片的属性必须要是颜色的属性，这个在前台显示的时候需要和sku进行关联的<br/>    * 商品图片关联在卖家身份和图片来源上的限制，卖家要是B卖家或订购了多图服务才能关联图片，并且图片要来自于卖家自己的图片空间才行<br/>    * 商品图片数量有限制。不管是上传的图片还是关联的图片，他们的总数不能超过一定限额，最多不能超过24张（每个颜色属性都有一张）
*/
func TaobaoItemJointPropimg(clt *core.SDKClient, req *product.TaobaoItemJointPropimgRequest, session string) (*product.TaobaoItemJointPropimgAPIResponse, error) {
    var resp product.TaobaoItemJointPropimgAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
