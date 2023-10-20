package tbitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoitemjointpropimgAPIRequest 商品关联属性图 API请求
// taobao.item.joint.propimg
//
// * 关联一张商品属性图片到num_iid指定的商品中&lt;br/&gt;    * 传入的num_iid所对应的商品必须属于当前会话的用户&lt;br/&gt;    * 图片的属性必须要是颜色的属性，这个在前台显示的时候需要和sku进行关联的&lt;br/&gt;    * 商品图片关联在卖家身份和图片来源上的限制，卖家要是B卖家或订购了多图服务才能关联图片，并且图片要来自于卖家自己的图片空间才行&lt;br/&gt;    * 商品图片数量有限制。不管是上传的图片还是关联的图片，他们的总数不能超过一定限额，最多不能超过24张（每个颜色属性都有一张）
type TaobaoitemjointpropimgAPIRequest struct {
	model.Params
	// 属性列表。调用taobao.itemprops.get获取，属性必须是颜色属性，格式:pid:vid。
	_properties string
	// 图片地址(传入图片相对地址即可,即不需包含 http://img02.taobao.net/bao/uploaded )
	_picPath string
	// 商品数字ID，必选
	_numIid int64
	// 属性图片ID。如果是新增不需要填写
	_id int64
	// 图片序号
	_position int64
}

// NewTaobaoitemjointpropimgRequest 初始化TaobaoitemjointpropimgAPIRequest对象
func NewTaobaoitemjointpropimgRequest() *TaobaoitemjointpropimgAPIRequest {
	return &TaobaoitemjointpropimgAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoitemjointpropimgAPIRequest) GetApiMethodName() string {
	return "taobao.item.joint.propimg"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoitemjointpropimgAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoitemjointpropimgAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetProperties is Properties Setter
// 属性列表。调用taobao.itemprops.get获取，属性必须是颜色属性，格式:pid:vid。
func (r *TaobaoitemjointpropimgAPIRequest) SetProperties(_properties string) error {
	r._properties = _properties
	r.Set("properties", _properties)
	return nil
}

// GetProperties Properties Getter
func (r TaobaoitemjointpropimgAPIRequest) GetProperties() string {
	return r._properties
}

// SetPicPath is PicPath Setter
// 图片地址(传入图片相对地址即可,即不需包含 http://img02.taobao.net/bao/uploaded )
func (r *TaobaoitemjointpropimgAPIRequest) SetPicPath(_picPath string) error {
	r._picPath = _picPath
	r.Set("pic_path", _picPath)
	return nil
}

// GetPicPath PicPath Getter
func (r TaobaoitemjointpropimgAPIRequest) GetPicPath() string {
	return r._picPath
}

// SetNumIid is NumIid Setter
// 商品数字ID，必选
func (r *TaobaoitemjointpropimgAPIRequest) SetNumIid(_numIid int64) error {
	r._numIid = _numIid
	r.Set("num_iid", _numIid)
	return nil
}

// GetNumIid NumIid Getter
func (r TaobaoitemjointpropimgAPIRequest) GetNumIid() int64 {
	return r._numIid
}

// SetId is Id Setter
// 属性图片ID。如果是新增不需要填写
func (r *TaobaoitemjointpropimgAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r TaobaoitemjointpropimgAPIRequest) GetId() int64 {
	return r._id
}

// SetPosition is Position Setter
// 图片序号
func (r *TaobaoitemjointpropimgAPIRequest) SetPosition(_position int64) error {
	r._position = _position
	r.Set("position", _position)
	return nil
}

// GetPosition Position Getter
func (r TaobaoitemjointpropimgAPIRequest) GetPosition() int64 {
	return r._position
}
