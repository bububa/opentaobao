package tanx

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
资质图片上传接口 APIRequest
taobao.tanx.qualification.picture.upload

资质图片上传接口
*/
type TaobaoTanxQualificationPictureUploadRequest struct {
    model.Params

    // dsp用户id
    memberId   int64 

    // dsp用户检验token
    token   string 

    // 1970年到现在的时间，毫秒
    signTime   int64 

    // File文件getByte后的二进制数组
    fileByte   []byte 

}

func NewTaobaoTanxQualificationPictureUploadRequest() *TaobaoTanxQualificationPictureUploadRequest{
    return &TaobaoTanxQualificationPictureUploadRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTanxQualificationPictureUploadRequest) GetApiMethodName() string {
    return "taobao.tanx.qualification.picture.upload"
}

func (r TaobaoTanxQualificationPictureUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTanxQualificationPictureUploadRequest) SetMemberId(memberId int64) error {
    r.memberId = memberId
    r.Set("member_id", memberId)
    return nil
}

func (r TaobaoTanxQualificationPictureUploadRequest) GetMemberId() int64 {
    return r.memberId
}

func (r *TaobaoTanxQualificationPictureUploadRequest) SetToken(token string) error {
    r.token = token
    r.Set("token", token)
    return nil
}

func (r TaobaoTanxQualificationPictureUploadRequest) GetToken() string {
    return r.token
}

func (r *TaobaoTanxQualificationPictureUploadRequest) SetSignTime(signTime int64) error {
    r.signTime = signTime
    r.Set("sign_time", signTime)
    return nil
}

func (r TaobaoTanxQualificationPictureUploadRequest) GetSignTime() int64 {
    return r.signTime
}

func (r *TaobaoTanxQualificationPictureUploadRequest) SetFileByte(fileByte []byte) error {
    r.fileByte = fileByte
    r.Set("file_byte", fileByte)
    return nil
}

func (r TaobaoTanxQualificationPictureUploadRequest) GetFileByte() []byte {
    return r.fileByte
}

