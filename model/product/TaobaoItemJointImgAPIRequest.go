package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoItemJointImgAPIRequest 商品关联子图 API请求
// taobao.item.joint.img
//
// * 关联一张商品图片到num_iid指定的商品中&lt;br/&gt;    * 传入的num_iid所对应的商品必须属于当前会话的用户&lt;br/&gt;    * 商品图片关联在卖家身份和图片来源上的限制，卖家要是B卖家或订购了多图服务才能关联图片，并且图片要来自于卖家自己的图片空间才行&lt;br/&gt;    * 商品图片数量有限制。不管是上传的图片还是关联的图片，他们的总数不能超过一定限额
type TaobaoItemJointImgAPIRequest struct {
	model.Params
	// 图片URL,图片空间图片的相对地址，支持的文件类型：jpg,jpeg,png
	_picPath string
	// 商品数字ID，必选
	_numIid int64
	// 商品图片id(如果是更新图片，则需要传该参数)
	_id int64
	// 图片序号
	_position int64
	// 上传的图片是否关联为商品主图
	_isMajor bool
	// 是否3:4长方形图片，绑定3:4主图视频时用于上传3:4商品主图
	_isRectangle bool
}

// NewTaobaoItemJointImgRequest 初始化TaobaoItemJointImgAPIRequest对象
func NewTaobaoItemJointImgRequest() *TaobaoItemJointImgAPIRequest {
	return &TaobaoItemJointImgAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoItemJointImgAPIRequest) GetApiMethodName() string {
	return "taobao.item.joint.img"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoItemJointImgAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetPicPath is PicPath Setter
// 图片URL,图片空间图片的相对地址，支持的文件类型：jpg,jpeg,png
func (r *TaobaoItemJointImgAPIRequest) SetPicPath(_picPath string) error {
	r._picPath = _picPath
	r.Set("pic_path", _picPath)
	return nil
}

// GetPicPath PicPath Getter
func (r TaobaoItemJointImgAPIRequest) GetPicPath() string {
	return r._picPath
}

// SetNumIid is NumIid Setter
// 商品数字ID，必选
func (r *TaobaoItemJointImgAPIRequest) SetNumIid(_numIid int64) error {
	r._numIid = _numIid
	r.Set("num_iid", _numIid)
	return nil
}

// GetNumIid NumIid Getter
func (r TaobaoItemJointImgAPIRequest) GetNumIid() int64 {
	return r._numIid
}

// SetId is Id Setter
// 商品图片id(如果是更新图片，则需要传该参数)
func (r *TaobaoItemJointImgAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r TaobaoItemJointImgAPIRequest) GetId() int64 {
	return r._id
}

// SetPosition is Position Setter
// 图片序号
func (r *TaobaoItemJointImgAPIRequest) SetPosition(_position int64) error {
	r._position = _position
	r.Set("position", _position)
	return nil
}

// GetPosition Position Getter
func (r TaobaoItemJointImgAPIRequest) GetPosition() int64 {
	return r._position
}

// SetIsMajor is IsMajor Setter
// 上传的图片是否关联为商品主图
func (r *TaobaoItemJointImgAPIRequest) SetIsMajor(_isMajor bool) error {
	r._isMajor = _isMajor
	r.Set("is_major", _isMajor)
	return nil
}

// GetIsMajor IsMajor Getter
func (r TaobaoItemJointImgAPIRequest) GetIsMajor() bool {
	return r._isMajor
}

// SetIsRectangle is IsRectangle Setter
// 是否3:4长方形图片，绑定3:4主图视频时用于上传3:4商品主图
func (r *TaobaoItemJointImgAPIRequest) SetIsRectangle(_isRectangle bool) error {
	r._isRectangle = _isRectangle
	r.Set("is_rectangle", _isRectangle)
	return nil
}

// GetIsRectangle IsRectangle Getter
func (r TaobaoItemJointImgAPIRequest) GetIsRectangle() bool {
	return r._isRectangle
}
