package media

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPictureDeleteAPIRequest 删除图片空间图片 API请求
// taobao.picture.delete
//
// 删除图片空间图片
type TaobaoPictureDeleteAPIRequest struct {
	model.Params
	// 图片ID字符串,可以一个也可以一组,用英文逗号间隔,如450,120,155.限制数量是100
	_pictureIds []string
}

// NewTaobaoPictureDeleteRequest 初始化TaobaoPictureDeleteAPIRequest对象
func NewTaobaoPictureDeleteRequest() *TaobaoPictureDeleteAPIRequest {
	return &TaobaoPictureDeleteAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoPictureDeleteAPIRequest) Reset() {
	r._pictureIds = r._pictureIds[:0]
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoPictureDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.picture.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoPictureDeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoPictureDeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPictureIds is PictureIds Setter
// 图片ID字符串,可以一个也可以一组,用英文逗号间隔,如450,120,155.限制数量是100
func (r *TaobaoPictureDeleteAPIRequest) SetPictureIds(_pictureIds []string) error {
	r._pictureIds = _pictureIds
	r.Set("picture_ids", _pictureIds)
	return nil
}

// GetPictureIds PictureIds Getter
func (r TaobaoPictureDeleteAPIRequest) GetPictureIds() []string {
	return r._pictureIds
}

var poolTaobaoPictureDeleteAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoPictureDeleteRequest()
	},
}

// GetTaobaoPictureDeleteRequest 从 sync.Pool 获取 TaobaoPictureDeleteAPIRequest
func GetTaobaoPictureDeleteAPIRequest() *TaobaoPictureDeleteAPIRequest {
	return poolTaobaoPictureDeleteAPIRequest.Get().(*TaobaoPictureDeleteAPIRequest)
}

// ReleaseTaobaoPictureDeleteAPIRequest 将 TaobaoPictureDeleteAPIRequest 放入 sync.Pool
func ReleaseTaobaoPictureDeleteAPIRequest(v *TaobaoPictureDeleteAPIRequest) {
	v.Reset()
	poolTaobaoPictureDeleteAPIRequest.Put(v)
}
