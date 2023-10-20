package media

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaopictureisreferencedgetAPIRequest 图片是否被引用 API请求
// taobao.picture.isreferenced.get
//
// 查询图片是否被引用，被引用返回true，未被引用返回false
type TaobaopictureisreferencedgetAPIRequest struct {
	model.Params
	// 图片id
	_pictureId int64
}

// NewTaobaopictureisreferencedgetRequest 初始化TaobaopictureisreferencedgetAPIRequest对象
func NewTaobaopictureisreferencedgetRequest() *TaobaopictureisreferencedgetAPIRequest {
	return &TaobaopictureisreferencedgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaopictureisreferencedgetAPIRequest) GetApiMethodName() string {
	return "taobao.picture.isreferenced.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaopictureisreferencedgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaopictureisreferencedgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPictureId is PictureId Setter
// 图片id
func (r *TaobaopictureisreferencedgetAPIRequest) SetPictureId(_pictureId int64) error {
	r._pictureId = _pictureId
	r.Set("picture_id", _pictureId)
	return nil
}

// GetPictureId PictureId Getter
func (r TaobaopictureisreferencedgetAPIRequest) GetPictureId() int64 {
	return r._pictureId
}
