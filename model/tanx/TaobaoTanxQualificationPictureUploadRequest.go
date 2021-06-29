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
    _memberId   int64
    // dsp用户检验token
    _token   string
    // 1970年到现在的时间，毫秒
    _signTime   int64
    // File文件getByte后的二进制数组
    _fileByte   *model.File
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
func (r *TaobaoTanxQualificationPictureUploadRequest) SetMemberId(_memberId int64) error {
    r._memberId = _memberId
    r.Set("member_id", _memberId)
    return nil
}

// MemberId Getter
func (r TaobaoTanxQualificationPictureUploadRequest) GetMemberId() int64 {
    return r._memberId
}
// Token Setter
// dsp用户检验token
func (r *TaobaoTanxQualificationPictureUploadRequest) SetToken(_token string) error {
    r._token = _token
    r.Set("token", _token)
    return nil
}

// Token Getter
func (r TaobaoTanxQualificationPictureUploadRequest) GetToken() string {
    return r._token
}
// SignTime Setter
// 1970年到现在的时间，毫秒
func (r *TaobaoTanxQualificationPictureUploadRequest) SetSignTime(_signTime int64) error {
    r._signTime = _signTime
    r.Set("sign_time", _signTime)
    return nil
}

// SignTime Getter
func (r TaobaoTanxQualificationPictureUploadRequest) GetSignTime() int64 {
    return r._signTime
}
// FileByte Setter
// File文件getByte后的二进制数组
func (r *TaobaoTanxQualificationPictureUploadRequest) SetFileByte(_fileByte *model.File) error {
    r._fileByte = _fileByte
    r.Set("file_byte", _fileByte)
    return nil
}

// FileByte Getter
func (r TaobaoTanxQualificationPictureUploadRequest) GetFileByte() *model.File {
    return r._fileByte
}
