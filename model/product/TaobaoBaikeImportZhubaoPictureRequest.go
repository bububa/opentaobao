package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
百科图片数据导入 API请求
taobao.baike.import.zhubao.picture

用于接入外部--图片--录入到商品百科中
*/
type TaobaoBaikeImportZhubaoPictureRequest struct {
    model.Params
    // 图片二进制数据
    _picture   *model.File
}

// 初始化TaobaoBaikeImportZhubaoPictureRequest对象
func NewTaobaoBaikeImportZhubaoPictureRequest() *TaobaoBaikeImportZhubaoPictureRequest{
    return &TaobaoBaikeImportZhubaoPictureRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoBaikeImportZhubaoPictureRequest) GetApiMethodName() string {
    return "taobao.baike.import.zhubao.picture"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoBaikeImportZhubaoPictureRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Picture Setter
// 图片二进制数据
func (r *TaobaoBaikeImportZhubaoPictureRequest) SetPicture(_picture *model.File) error {
    r._picture = _picture
    r.Set("picture", _picture)
    return nil
}

// Picture Getter
func (r TaobaoBaikeImportZhubaoPictureRequest) GetPicture() *model.File {
    return r._picture
}
