package media

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaopicturedeleteAPIRequest 删除图片空间图片 API请求
// taobao.picture.delete
//
// 删除图片空间图片
type TaobaopicturedeleteAPIRequest struct {
	model.Params
	// 图片ID字符串,可以一个也可以一组,用英文逗号间隔,如450,120,155.限制数量是100
	_pictureIds []string
}

// NewTaobaopicturedeleteRequest 初始化TaobaopicturedeleteAPIRequest对象
func NewTaobaopicturedeleteRequest() *TaobaopicturedeleteAPIRequest {
	return &TaobaopicturedeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaopicturedeleteAPIRequest) GetApiMethodName() string {
	return "taobao.picture.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaopicturedeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaopicturedeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPictureIds is PictureIds Setter
// 图片ID字符串,可以一个也可以一组,用英文逗号间隔,如450,120,155.限制数量是100
func (r *TaobaopicturedeleteAPIRequest) SetPictureIds(_pictureIds []string) error {
	r._pictureIds = _pictureIds
	r.Set("picture_ids", _pictureIds)
	return nil
}

// GetPictureIds PictureIds Getter
func (r TaobaopicturedeleteAPIRequest) GetPictureIds() []string {
	return r._pictureIds
}
