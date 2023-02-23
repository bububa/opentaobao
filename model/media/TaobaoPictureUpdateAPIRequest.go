package media

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPictureUpdateAPIRequest 修改图片名字 API请求
// taobao.picture.update
//
// 修改指定图片的图片名
type TaobaoPictureUpdateAPIRequest struct {
	model.Params
	// 新的图片名，最大长度50字符，不能为空
	_newName string
	// 要更改名字的图片的id
	_pictureId int64
}

// NewTaobaoPictureUpdateRequest 初始化TaobaoPictureUpdateAPIRequest对象
func NewTaobaoPictureUpdateRequest() *TaobaoPictureUpdateAPIRequest {
	return &TaobaoPictureUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoPictureUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.picture.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoPictureUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoPictureUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNewName is NewName Setter
// 新的图片名，最大长度50字符，不能为空
func (r *TaobaoPictureUpdateAPIRequest) SetNewName(_newName string) error {
	r._newName = _newName
	r.Set("new_name", _newName)
	return nil
}

// GetNewName NewName Getter
func (r TaobaoPictureUpdateAPIRequest) GetNewName() string {
	return r._newName
}

// SetPictureId is PictureId Setter
// 要更改名字的图片的id
func (r *TaobaoPictureUpdateAPIRequest) SetPictureId(_pictureId int64) error {
	r._pictureId = _pictureId
	r.Set("picture_id", _pictureId)
	return nil
}

// GetPictureId PictureId Getter
func (r TaobaoPictureUpdateAPIRequest) GetPictureId() int64 {
	return r._pictureId
}
