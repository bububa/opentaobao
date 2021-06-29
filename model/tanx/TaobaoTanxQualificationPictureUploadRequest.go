package tanx

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
资质图片上传接口 API请求
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
    fileByte   []*model.File
}

// 初始化TaobaoTanxQualificationPictureUploadRequest对象
func NewTaobaoTanxQualificationPictureUploadRequest() *TaobaoTanxQualificationPictureUploadRequest{
    return &TaobaoTanxQualificationPictureUploadRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTanxQualificationPictureUploadRequest) GetApiMethodName() string {
    return "taobao.tanx.qualification.picture.upload"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTanxQualificationPictureUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// MemberId Setter
// dsp用户id
func (r *TaobaoTanxQualificationPictureUploadRequest) SetMemberId(memberId int64) error {
    r.memberId = memberId
    r.Set("member_id", memberId)
    return nil
}

// MemberId Getter
func (r TaobaoTanxQualificationPictureUploadRequest) GetMemberId() int64 {
    return r.memberId
}
// Token Setter
// dsp用户检验token
func (r *TaobaoTanxQualificationPictureUploadRequest) SetToken(token string) error {
    r.token = token
    r.Set("token", token)
    return nil
}

// Token Getter
func (r TaobaoTanxQualificationPictureUploadRequest) GetToken() string {
    return r.token
}
// SignTime Setter
// 1970年到现在的时间，毫秒
func (r *TaobaoTanxQualificationPictureUploadRequest) SetSignTime(signTime int64) error {
    r.signTime = signTime
    r.Set("sign_time", signTime)
    return nil
}

// SignTime Getter
func (r TaobaoTanxQualificationPictureUploadRequest) GetSignTime() int64 {
    return r.signTime
}
// FileByte Setter
// File文件getByte后的二进制数组
func (r *TaobaoTanxQualificationPictureUploadRequest) SetFileByte(fileByte []*model.File) error {
    r.fileByte = fileByte
    r.Set("file_byte", fileByte)
    return nil
}

// FileByte Getter
func (r TaobaoTanxQualificationPictureUploadRequest) GetFileByte() []*model.File {
    return r.fileByte
}
