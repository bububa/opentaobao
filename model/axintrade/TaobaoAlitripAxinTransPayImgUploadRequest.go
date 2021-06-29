package axintrade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
上传图片到支付宝图片空间接口 APIRequest
taobao.alitrip.axin.trans.pay.img.upload

阿信供销平台-上传图片到支付宝图片空间接口
*/
type TaobaoAlitripAxinTransPayImgUploadRequest struct {
    model.Params

    // 上传图片到支付宝图片空间接口入参
    axinPayImgUploadDTO   *AxinPayImgUploadDto 

    // 图片字节流
    imgContents   []*model.File 

}

func NewTaobaoAlitripAxinTransPayImgUploadRequest() *TaobaoAlitripAxinTransPayImgUploadRequest{
    return &TaobaoAlitripAxinTransPayImgUploadRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAlitripAxinTransPayImgUploadRequest) GetApiMethodName() string {
    return "taobao.alitrip.axin.trans.pay.img.upload"
}

func (r TaobaoAlitripAxinTransPayImgUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAlitripAxinTransPayImgUploadRequest) SetAxinPayImgUploadDTO(axinPayImgUploadDTO *AxinPayImgUploadDto) error {
    r.axinPayImgUploadDTO = axinPayImgUploadDTO
    r.Set("axin_pay_img_upload_d_t_o", axinPayImgUploadDTO)
    return nil
}

func (r TaobaoAlitripAxinTransPayImgUploadRequest) GetAxinPayImgUploadDTO() *AxinPayImgUploadDto {
    return r.axinPayImgUploadDTO
}

func (r *TaobaoAlitripAxinTransPayImgUploadRequest) SetImgContents(imgContents []*model.File) error {
    r.imgContents = imgContents
    r.Set("img_contents", imgContents)
    return nil
}

func (r TaobaoAlitripAxinTransPayImgUploadRequest) GetImgContents() []*model.File {
    return r.imgContents
}

