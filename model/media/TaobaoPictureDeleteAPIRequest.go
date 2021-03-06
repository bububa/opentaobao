package media

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoPictureDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.picture.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoPictureDeleteAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
