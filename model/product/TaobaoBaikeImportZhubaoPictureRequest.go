package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
百科图片数据导入 APIRequest
taobao.baike.import.zhubao.picture

用于接入外部--图片--录入到商品百科中
*/
type TaobaoBaikeImportZhubaoPictureRequest struct {
    model.Params

    // 图片二进制数据
    picture   []byte 

}

func NewTaobaoBaikeImportZhubaoPictureRequest() *TaobaoBaikeImportZhubaoPictureRequest{
    return &TaobaoBaikeImportZhubaoPictureRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoBaikeImportZhubaoPictureRequest) GetApiMethodName() string {
    return "taobao.baike.import.zhubao.picture"
}

func (r TaobaoBaikeImportZhubaoPictureRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoBaikeImportZhubaoPictureRequest) SetPicture(picture []byte) error {
    r.picture = picture
    r.Set("picture", picture)
    return nil
}

func (r TaobaoBaikeImportZhubaoPictureRequest) GetPicture() []byte {
    return r.picture
}

