package media

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPictureIsreferencedGetAPIRequest 图片是否被引用 API请求
// taobao.picture.isreferenced.get
//
// 查询图片是否被引用，被引用返回true，未被引用返回false
type TaobaoPictureIsreferencedGetAPIRequest struct {
	model.Params
	// 图片id
	_pictureId int64
}

// NewTaobaoPictureIsreferencedGetRequest 初始化TaobaoPictureIsreferencedGetAPIRequest对象
func NewTaobaoPictureIsreferencedGetRequest() *TaobaoPictureIsreferencedGetAPIRequest {
	return &TaobaoPictureIsreferencedGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoPictureIsreferencedGetAPIRequest) GetApiMethodName() string {
	return "taobao.picture.isreferenced.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoPictureIsreferencedGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is PictureId Setter
// 图片id
func (r *TaobaoPictureIsreferencedGetAPIRequest) SetPictureId(_pictureId int64) error {
	r._pictureId = _pictureId
	r.Set("picture_id", _pictureId)
	return nil
}

// Get PictureId Getter
func (r TaobaoPictureIsreferencedGetAPIRequest) GetPictureId() int64 {
	return r._pictureId
}
