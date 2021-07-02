package media

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPictureReplaceAPIRequest 替换图片 API请求
// taobao.picture.replace
//
// 替换一张图片，只替换图片数据，图片名称，图片分类等保持不变。
type TaobaoPictureReplaceAPIRequest struct {
	model.Params
	// 要替换的图片的id，必须大于0
	_pictureId int64
	// 图片二进制文件流,不能为空,允许png、jpg、gif图片格式
	_imageData *model.File
}

// NewTaobaoPictureReplaceRequest 初始化TaobaoPictureReplaceAPIRequest对象
func NewTaobaoPictureReplaceRequest() *TaobaoPictureReplaceAPIRequest {
	return &TaobaoPictureReplaceAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoPictureReplaceAPIRequest) GetApiMethodName() string {
	return "taobao.picture.replace"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoPictureReplaceAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is PictureId Setter
// 要替换的图片的id，必须大于0
func (r *TaobaoPictureReplaceAPIRequest) SetPictureId(_pictureId int64) error {
	r._pictureId = _pictureId
	r.Set("picture_id", _pictureId)
	return nil
}

// Get PictureId Getter
func (r TaobaoPictureReplaceAPIRequest) GetPictureId() int64 {
	return r._pictureId
}

// Set is ImageData Setter
// 图片二进制文件流,不能为空,允许png、jpg、gif图片格式
func (r *TaobaoPictureReplaceAPIRequest) SetImageData(_imageData *model.File) error {
	r._imageData = _imageData
	r.Set("image_data", _imageData)
	return nil
}

// Get ImageData Getter
func (r TaobaoPictureReplaceAPIRequest) GetImageData() *model.File {
	return r._imageData
}
