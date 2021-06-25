package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商品关联子图 APIRequest
taobao.item.joint.img

* 关联一张商品图片到num_iid指定的商品中<br/>    * 传入的num_iid所对应的商品必须属于当前会话的用户<br/>    * 商品图片关联在卖家身份和图片来源上的限制，卖家要是B卖家或订购了多图服务才能关联图片，并且图片要来自于卖家自己的图片空间才行<br/>    * 商品图片数量有限制。不管是上传的图片还是关联的图片，他们的总数不能超过一定限额
*/
type TaobaoItemJointImgRequest struct {
    model.Params

    // 商品图片id(如果是更新图片，则需要传该参数)
    id   int64 

    // 商品数字ID，必选
    numIid   int64 

    // 图片URL,图片空间图片的相对地址，支持的文件类型：jpg,jpeg,png
    picPath   string 

    // 上传的图片是否关联为商品主图
    isMajor   bool 

    // 图片序号
    position   int64 

    // 是否3:4长方形图片，绑定3:4主图视频时用于上传3:4商品主图
    isRectangle   bool 

}

func NewTaobaoItemJointImgRequest() *TaobaoItemJointImgRequest{
    return &TaobaoItemJointImgRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoItemJointImgRequest) GetApiMethodName() string {
    return "taobao.item.joint.img"
}

func (r TaobaoItemJointImgRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoItemJointImgRequest) SetId(id int64) error {
    r.id = id
    r.Set("id", id)
    return nil
}

func (r TaobaoItemJointImgRequest) GetId() int64 {
    return r.id
}

func (r *TaobaoItemJointImgRequest) SetNumIid(numIid int64) error {
    r.numIid = numIid
    r.Set("num_iid", numIid)
    return nil
}

func (r TaobaoItemJointImgRequest) GetNumIid() int64 {
    return r.numIid
}

func (r *TaobaoItemJointImgRequest) SetPicPath(picPath string) error {
    r.picPath = picPath
    r.Set("pic_path", picPath)
    return nil
}

func (r TaobaoItemJointImgRequest) GetPicPath() string {
    return r.picPath
}

func (r *TaobaoItemJointImgRequest) SetIsMajor(isMajor bool) error {
    r.isMajor = isMajor
    r.Set("is_major", isMajor)
    return nil
}

func (r TaobaoItemJointImgRequest) GetIsMajor() bool {
    return r.isMajor
}

func (r *TaobaoItemJointImgRequest) SetPosition(position int64) error {
    r.position = position
    r.Set("position", position)
    return nil
}

func (r TaobaoItemJointImgRequest) GetPosition() int64 {
    return r.position
}

func (r *TaobaoItemJointImgRequest) SetIsRectangle(isRectangle bool) error {
    r.isRectangle = isRectangle
    r.Set("is_rectangle", isRectangle)
    return nil
}

func (r TaobaoItemJointImgRequest) GetIsRectangle() bool {
    return r.isRectangle
}

