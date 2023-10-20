package tbitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoitempropimguploadAPIRequest 添加或修改属性图片 API请求
// taobao.item.propimg.upload
//
// 添加一张商品属性图片到num_iid指定的商品中 &lt;br/&gt;传入的num_iid所对应的商品必须属于当前会话的用户 &lt;br/&gt;图片的属性必须要是颜色的属性，这个在前台显示的时候需要和sku进行关联的 &lt;br/&gt;商品属性图片只有享有服务的卖家（如：淘宝大卖家、订购了淘宝多图服务的卖家）才能上传 &lt;br/&gt;商品属性图片有数量和大小上的限制，最多不能超过24张（每个颜色属性都有一张）。
type TaobaoitempropimguploadAPIRequest struct {
	model.Params
	// 属性列表。调用taobao.itemprops.get获取类目属性，属性必须是颜色属性，再用taobao.itempropvalues.get取得vid。格式:pid:vid。
	_properties string
	// 商品数字ID，必选
	_numIid int64
	// 属性图片ID。如果是新增不需要填写
	_id int64
	// 图片位置
	_position int64
	// 属性图片内容。类型:JPG,GIF;图片大小不超过:3M
	_image *model.File
}

// NewTaobaoitempropimguploadRequest 初始化TaobaoitempropimguploadAPIRequest对象
func NewTaobaoitempropimguploadRequest() *TaobaoitempropimguploadAPIRequest {
	return &TaobaoitempropimguploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoitempropimguploadAPIRequest) GetApiMethodName() string {
	return "taobao.item.propimg.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoitempropimguploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoitempropimguploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetProperties is Properties Setter
// 属性列表。调用taobao.itemprops.get获取类目属性，属性必须是颜色属性，再用taobao.itempropvalues.get取得vid。格式:pid:vid。
func (r *TaobaoitempropimguploadAPIRequest) SetProperties(_properties string) error {
	r._properties = _properties
	r.Set("properties", _properties)
	return nil
}

// GetProperties Properties Getter
func (r TaobaoitempropimguploadAPIRequest) GetProperties() string {
	return r._properties
}

// SetNumIid is NumIid Setter
// 商品数字ID，必选
func (r *TaobaoitempropimguploadAPIRequest) SetNumIid(_numIid int64) error {
	r._numIid = _numIid
	r.Set("num_iid", _numIid)
	return nil
}

// GetNumIid NumIid Getter
func (r TaobaoitempropimguploadAPIRequest) GetNumIid() int64 {
	return r._numIid
}

// SetId is Id Setter
// 属性图片ID。如果是新增不需要填写
func (r *TaobaoitempropimguploadAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r TaobaoitempropimguploadAPIRequest) GetId() int64 {
	return r._id
}

// SetPosition is Position Setter
// 图片位置
func (r *TaobaoitempropimguploadAPIRequest) SetPosition(_position int64) error {
	r._position = _position
	r.Set("position", _position)
	return nil
}

// GetPosition Position Getter
func (r TaobaoitempropimguploadAPIRequest) GetPosition() int64 {
	return r._position
}

// SetImage is Image Setter
// 属性图片内容。类型:JPG,GIF;图片大小不超过:3M
func (r *TaobaoitempropimguploadAPIRequest) SetImage(_image *model.File) error {
	r._image = _image
	r.Set("image", _image)
	return nil
}

// GetImage Image Getter
func (r TaobaoitempropimguploadAPIRequest) GetImage() *model.File {
	return r._image
}
