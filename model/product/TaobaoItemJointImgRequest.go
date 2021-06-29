package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商品关联子图 API请求
taobao.item.joint.img

* 关联一张商品图片到num_iid指定的商品中<br/>    * 传入的num_iid所对应的商品必须属于当前会话的用户<br/>    * 商品图片关联在卖家身份和图片来源上的限制，卖家要是B卖家或订购了多图服务才能关联图片，并且图片要来自于卖家自己的图片空间才行<br/>    * 商品图片数量有限制。不管是上传的图片还是关联的图片，他们的总数不能超过一定限额
*/
type TaobaoItemJointImgRequest struct {
    model.Params
    // 商品图片id(如果是更新图片，则需要传该参数)
    _id   int64
    // 商品数字ID，必选
    _numIid   int64
    // 图片URL,图片空间图片的相对地址，支持的文件类型：jpg,jpeg,png
    _picPath   string
    // 上传的图片是否关联为商品主图
    _isMajor   bool
    // 图片序号
    _position   int64
    // 是否3:4长方形图片，绑定3:4主图视频时用于上传3:4商品主图
    _isRectangle   bool
}

// 初始化TaobaoItemJointImgRequest对象
func NewTaobaoItemJointImgRequest() *TaobaoItemJointImgRequest{
    return &TaobaoItemJointImgRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoItemJointImgRequest) GetApiMethodName() string {
    return "taobao.item.joint.img"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoItemJointImgRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Id Setter
// 商品图片id(如果是更新图片，则需要传该参数)
func (r *TaobaoItemJointImgRequest) SetId(_id int64) error {
    r._id = _id
    r.Set("id", _id)
    return nil
}

// Id Getter
func (r TaobaoItemJointImgRequest) GetId() int64 {
    return r._id
}
// NumIid Setter
// 商品数字ID，必选
func (r *TaobaoItemJointImgRequest) SetNumIid(_numIid int64) error {
    r._numIid = _numIid
    r.Set("num_iid", _numIid)
    return nil
}

// NumIid Getter
func (r TaobaoItemJointImgRequest) GetNumIid() int64 {
    return r._numIid
}
// PicPath Setter
// 图片URL,图片空间图片的相对地址，支持的文件类型：jpg,jpeg,png
func (r *TaobaoItemJointImgRequest) SetPicPath(_picPath string) error {
    r._picPath = _picPath
    r.Set("pic_path", _picPath)
    return nil
}

// PicPath Getter
func (r TaobaoItemJointImgRequest) GetPicPath() string {
    return r._picPath
}
// IsMajor Setter
// 上传的图片是否关联为商品主图
func (r *TaobaoItemJointImgRequest) SetIsMajor(_isMajor bool) error {
    r._isMajor = _isMajor
    r.Set("is_major", _isMajor)
    return nil
}

// IsMajor Getter
func (r TaobaoItemJointImgRequest) GetIsMajor() bool {
    return r._isMajor
}
// Position Setter
// 图片序号
func (r *TaobaoItemJointImgRequest) SetPosition(_position int64) error {
    r._position = _position
    r.Set("position", _position)
    return nil
}

// Position Getter
func (r TaobaoItemJointImgRequest) GetPosition() int64 {
    return r._position
}
// IsRectangle Setter
// 是否3:4长方形图片，绑定3:4主图视频时用于上传3:4商品主图
func (r *TaobaoItemJointImgRequest) SetIsRectangle(_isRectangle bool) error {
    r._isRectangle = _isRectangle
    r.Set("is_rectangle", _isRectangle)
    return nil
}

// IsRectangle Getter
func (r TaobaoItemJointImgRequest) GetIsRectangle() bool {
    return r._isRectangle
}
