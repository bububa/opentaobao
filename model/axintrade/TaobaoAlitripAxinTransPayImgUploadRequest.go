package axintrade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
上传图片到支付宝图片空间接口 API请求
taobao.alitrip.axin.trans.pay.img.upload

阿信供销平台-上传图片到支付宝图片空间接口
*/
type TaobaoAlitripAxinTransPayImgUploadRequest struct {
    model.Params
    // 上传图片到支付宝图片空间接口入参
    _axinPayImgUploadDTO   *AxinPayImgUploadDto
    // 图片字节流
    _imgContents   []*model.File
}

// 初始化TaobaoAlitripAxinTransPayImgUploadRequest对象
func NewTaobaoAlitripAxinTransPayImgUploadRequest() *TaobaoAlitripAxinTransPayImgUploadRequest{
    return &TaobaoAlitripAxinTransPayImgUploadRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAlitripAxinTransPayImgUploadRequest) GetApiMethodName() string {
    return "taobao.alitrip.axin.trans.pay.img.upload"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAlitripAxinTransPayImgUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AxinPayImgUploadDTO Setter
// 上传图片到支付宝图片空间接口入参
func (r *TaobaoAlitripAxinTransPayImgUploadRequest) SetAxinPayImgUploadDTO(_axinPayImgUploadDTO *AxinPayImgUploadDto) error {
    r._axinPayImgUploadDTO = _axinPayImgUploadDTO
    r.Set("axin_pay_img_upload_d_t_o", _axinPayImgUploadDTO)
    return nil
}

// AxinPayImgUploadDTO Getter
func (r TaobaoAlitripAxinTransPayImgUploadRequest) GetAxinPayImgUploadDTO() *AxinPayImgUploadDto {
    return r._axinPayImgUploadDTO
}
// ImgContents Setter
// 图片字节流
func (r *TaobaoAlitripAxinTransPayImgUploadRequest) SetImgContents(_imgContents []*model.File) error {
    r._imgContents = _imgContents
    r.Set("img_contents", _imgContents)
    return nil
}

// ImgContents Getter
func (r TaobaoAlitripAxinTransPayImgUploadRequest) GetImgContents() []*model.File {
    return r._imgContents
}
