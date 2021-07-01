package media

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
修改图片名字 API请求
taobao.picture.update

修改指定图片的图片名
*/
type TaobaoPictureUpdateAPIRequest struct {
    model.Params
    // 要更改名字的图片的id
    _pictureId   int64
    // 新的图片名，最大长度50字符，不能为空
    _newName   string
}

// 初始化TaobaoPictureUpdateAPIRequest对象
func NewTaobaoPictureUpdateRequest() *TaobaoPictureUpdateAPIRequest{
    return &TaobaoPictureUpdateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoPictureUpdateAPIRequest) GetApiMethodName() string {
    return "taobao.picture.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoPictureUpdateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PictureId Setter
// 要更改名字的图片的id
func (r *TaobaoPictureUpdateAPIRequest) SetPictureId(_pictureId int64) error {
    r._pictureId = _pictureId
    r.Set("picture_id", _pictureId)
    return nil
}

// PictureId Getter
func (r TaobaoPictureUpdateAPIRequest) GetPictureId() int64 {
    return r._pictureId
}
// NewName Setter
// 新的图片名，最大长度50字符，不能为空
func (r *TaobaoPictureUpdateAPIRequest) SetNewName(_newName string) error {
    r._newName = _newName
    r.Set("new_name", _newName)
    return nil
}

// NewName Getter
func (r TaobaoPictureUpdateAPIRequest) GetNewName() string {
    return r._newName
}
