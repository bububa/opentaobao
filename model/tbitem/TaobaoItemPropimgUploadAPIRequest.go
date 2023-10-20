package tbitem

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoItemPropimgUploadAPIRequest 添加或修改属性图片 API请求
// taobao.item.propimg.upload
//
// 添加一张商品属性图片到num_iid指定的商品中 &lt;br/&gt;传入的num_iid所对应的商品必须属于当前会话的用户 &lt;br/&gt;图片的属性必须要是颜色的属性，这个在前台显示的时候需要和sku进行关联的 &lt;br/&gt;商品属性图片只有享有服务的卖家（如：淘宝大卖家、订购了淘宝多图服务的卖家）才能上传 &lt;br/&gt;商品属性图片有数量和大小上的限制，最多不能超过24张（每个颜色属性都有一张）。
type TaobaoItemPropimgUploadAPIRequest struct {
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

// NewTaobaoItemPropimgUploadRequest 初始化TaobaoItemPropimgUploadAPIRequest对象
func NewTaobaoItemPropimgUploadRequest() *TaobaoItemPropimgUploadAPIRequest {
	return &TaobaoItemPropimgUploadAPIRequest{
		Params: model.NewParams(5),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoItemPropimgUploadAPIRequest) Reset() {
	r._properties = ""
	r._numIid = 0
	r._id = 0
	r._position = 0
	r._image = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoItemPropimgUploadAPIRequest) GetApiMethodName() string {
	return "taobao.item.propimg.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoItemPropimgUploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoItemPropimgUploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetProperties is Properties Setter
// 属性列表。调用taobao.itemprops.get获取类目属性，属性必须是颜色属性，再用taobao.itempropvalues.get取得vid。格式:pid:vid。
func (r *TaobaoItemPropimgUploadAPIRequest) SetProperties(_properties string) error {
	r._properties = _properties
	r.Set("properties", _properties)
	return nil
}

// GetProperties Properties Getter
func (r TaobaoItemPropimgUploadAPIRequest) GetProperties() string {
	return r._properties
}

// SetNumIid is NumIid Setter
// 商品数字ID，必选
func (r *TaobaoItemPropimgUploadAPIRequest) SetNumIid(_numIid int64) error {
	r._numIid = _numIid
	r.Set("num_iid", _numIid)
	return nil
}

// GetNumIid NumIid Getter
func (r TaobaoItemPropimgUploadAPIRequest) GetNumIid() int64 {
	return r._numIid
}

// SetId is Id Setter
// 属性图片ID。如果是新增不需要填写
func (r *TaobaoItemPropimgUploadAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r TaobaoItemPropimgUploadAPIRequest) GetId() int64 {
	return r._id
}

// SetPosition is Position Setter
// 图片位置
func (r *TaobaoItemPropimgUploadAPIRequest) SetPosition(_position int64) error {
	r._position = _position
	r.Set("position", _position)
	return nil
}

// GetPosition Position Getter
func (r TaobaoItemPropimgUploadAPIRequest) GetPosition() int64 {
	return r._position
}

// SetImage is Image Setter
// 属性图片内容。类型:JPG,GIF;图片大小不超过:3M
func (r *TaobaoItemPropimgUploadAPIRequest) SetImage(_image *model.File) error {
	r._image = _image
	r.Set("image", _image)
	return nil
}

// GetImage Image Getter
func (r TaobaoItemPropimgUploadAPIRequest) GetImage() *model.File {
	return r._image
}

var poolTaobaoItemPropimgUploadAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoItemPropimgUploadRequest()
	},
}

// GetTaobaoItemPropimgUploadRequest 从 sync.Pool 获取 TaobaoItemPropimgUploadAPIRequest
func GetTaobaoItemPropimgUploadAPIRequest() *TaobaoItemPropimgUploadAPIRequest {
	return poolTaobaoItemPropimgUploadAPIRequest.Get().(*TaobaoItemPropimgUploadAPIRequest)
}

// ReleaseTaobaoItemPropimgUploadAPIRequest 将 TaobaoItemPropimgUploadAPIRequest 放入 sync.Pool
func ReleaseTaobaoItemPropimgUploadAPIRequest(v *TaobaoItemPropimgUploadAPIRequest) {
	v.Reset()
	poolTaobaoItemPropimgUploadAPIRequest.Put(v)
}
