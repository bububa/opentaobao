package media

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoPictureIsreferencedGetAPIRequest) Reset() {
	r._pictureId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoPictureIsreferencedGetAPIRequest) GetApiMethodName() string {
	return "taobao.picture.isreferenced.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoPictureIsreferencedGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoPictureIsreferencedGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPictureId is PictureId Setter
// 图片id
func (r *TaobaoPictureIsreferencedGetAPIRequest) SetPictureId(_pictureId int64) error {
	r._pictureId = _pictureId
	r.Set("picture_id", _pictureId)
	return nil
}

// GetPictureId PictureId Getter
func (r TaobaoPictureIsreferencedGetAPIRequest) GetPictureId() int64 {
	return r._pictureId
}

var poolTaobaoPictureIsreferencedGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoPictureIsreferencedGetRequest()
	},
}

// GetTaobaoPictureIsreferencedGetRequest 从 sync.Pool 获取 TaobaoPictureIsreferencedGetAPIRequest
func GetTaobaoPictureIsreferencedGetAPIRequest() *TaobaoPictureIsreferencedGetAPIRequest {
	return poolTaobaoPictureIsreferencedGetAPIRequest.Get().(*TaobaoPictureIsreferencedGetAPIRequest)
}

// ReleaseTaobaoPictureIsreferencedGetAPIRequest 将 TaobaoPictureIsreferencedGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoPictureIsreferencedGetAPIRequest(v *TaobaoPictureIsreferencedGetAPIRequest) {
	v.Reset()
	poolTaobaoPictureIsreferencedGetAPIRequest.Put(v)
}
