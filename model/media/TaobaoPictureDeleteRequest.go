package media

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除图片空间图片 API请求
taobao.picture.delete

删除图片空间图片
*/
type TaobaoPictureDeleteRequest struct {
    model.Params
    // 图片ID字符串,可以一个也可以一组,用英文逗号间隔,如450,120,155.限制数量是100
    pictureIds   []string
}

// 初始化TaobaoPictureDeleteRequest对象
func NewTaobaoPictureDeleteRequest() *TaobaoPictureDeleteRequest{
    return &TaobaoPictureDeleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoPictureDeleteRequest) GetApiMethodName() string {
    return "taobao.picture.delete"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoPictureDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PictureIds Setter
// 图片ID字符串,可以一个也可以一组,用英文逗号间隔,如450,120,155.限制数量是100
func (r *TaobaoPictureDeleteRequest) SetPictureIds(pictureIds []string) error {
    r.pictureIds = pictureIds
    r.Set("picture_ids", pictureIds)
    return nil
}

// PictureIds Getter
func (r TaobaoPictureDeleteRequest) GetPictureIds() []string {
    return r.pictureIds
}
