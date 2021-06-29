package alihealthmdeer

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
合作伙伴视频同步给医知鹿接口 API请求
alibaba.alihealth.mdeer.video.sync

合伙做伴内容同步接口，用来视频同步
*/
type AlibabaAlihealthMdeerVideoSyncRequest struct {
    model.Params
    // 合作方头像url
    _partnerPortraitUrl   string
    // 作者电话
    _phoneNumber   string
    // 作者简介
    _authorIntroduction   string
    // 作者科室
    _authorDepartment   string
    // 作者级别
    _authorLevel   string
    // 医院级别
    _hospitalLevel   string
    // 医院名称
    _hospitalName   string
    // 作者头像
    _portraitUrl   string
    // 作者名称
    _authorName   string
    // 作者id
    _authorId   string
    // 合作方主页
    _partnerHomepage   string
    // 合作方名称
    _partnerName   string
    // 发布日期
    _releaseDate   string
    // 视频文件url
    _videoFileUrl   string
    // 视频落地页
    _videoMobileUrl   string
    // 视频简介
    _videoIntroduction   string
    // 视频长度
    _videoLength   string
    // 视频所述疾病
    _disease   string
    // 预览图url
    _priviewUrl   string
    // 视频标题
    _videoTitle   string
    // 视频id
    _videoId   string
}

// 初始化AlibabaAlihealthMdeerVideoSyncRequest对象
func NewAlibabaAlihealthMdeerVideoSyncRequest() *AlibabaAlihealthMdeerVideoSyncRequest{
    return &AlibabaAlihealthMdeerVideoSyncRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthMdeerVideoSyncRequest) GetApiMethodName() string {
    return "alibaba.alihealth.mdeer.video.sync"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthMdeerVideoSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PartnerPortraitUrl Setter
// 合作方头像url
func (r *AlibabaAlihealthMdeerVideoSyncRequest) SetPartnerPortraitUrl(_partnerPortraitUrl string) error {
    r._partnerPortraitUrl = _partnerPortraitUrl
    r.Set("partner_portrait_url", _partnerPortraitUrl)
    return nil
}

// PartnerPortraitUrl Getter
func (r AlibabaAlihealthMdeerVideoSyncRequest) GetPartnerPortraitUrl() string {
    return r._partnerPortraitUrl
}
// PhoneNumber Setter
// 作者电话
func (r *AlibabaAlihealthMdeerVideoSyncRequest) SetPhoneNumber(_phoneNumber string) error {
    r._phoneNumber = _phoneNumber
    r.Set("phone_number", _phoneNumber)
    return nil
}

// PhoneNumber Getter
func (r AlibabaAlihealthMdeerVideoSyncRequest) GetPhoneNumber() string {
    return r._phoneNumber
}
// AuthorIntroduction Setter
// 作者简介
func (r *AlibabaAlihealthMdeerVideoSyncRequest) SetAuthorIntroduction(_authorIntroduction string) error {
    r._authorIntroduction = _authorIntroduction
    r.Set("author_introduction", _authorIntroduction)
    return nil
}

// AuthorIntroduction Getter
func (r AlibabaAlihealthMdeerVideoSyncRequest) GetAuthorIntroduction() string {
    return r._authorIntroduction
}
// AuthorDepartment Setter
// 作者科室
func (r *AlibabaAlihealthMdeerVideoSyncRequest) SetAuthorDepartment(_authorDepartment string) error {
    r._authorDepartment = _authorDepartment
    r.Set("author_department", _authorDepartment)
    return nil
}

// AuthorDepartment Getter
func (r AlibabaAlihealthMdeerVideoSyncRequest) GetAuthorDepartment() string {
    return r._authorDepartment
}
// AuthorLevel Setter
// 作者级别
func (r *AlibabaAlihealthMdeerVideoSyncRequest) SetAuthorLevel(_authorLevel string) error {
    r._authorLevel = _authorLevel
    r.Set("author_level", _authorLevel)
    return nil
}

// AuthorLevel Getter
func (r AlibabaAlihealthMdeerVideoSyncRequest) GetAuthorLevel() string {
    return r._authorLevel
}
// HospitalLevel Setter
// 医院级别
func (r *AlibabaAlihealthMdeerVideoSyncRequest) SetHospitalLevel(_hospitalLevel string) error {
    r._hospitalLevel = _hospitalLevel
    r.Set("hospital_level", _hospitalLevel)
    return nil
}

// HospitalLevel Getter
func (r AlibabaAlihealthMdeerVideoSyncRequest) GetHospitalLevel() string {
    return r._hospitalLevel
}
// HospitalName Setter
// 医院名称
func (r *AlibabaAlihealthMdeerVideoSyncRequest) SetHospitalName(_hospitalName string) error {
    r._hospitalName = _hospitalName
    r.Set("hospital_name", _hospitalName)
    return nil
}

// HospitalName Getter
func (r AlibabaAlihealthMdeerVideoSyncRequest) GetHospitalName() string {
    return r._hospitalName
}
// PortraitUrl Setter
// 作者头像
func (r *AlibabaAlihealthMdeerVideoSyncRequest) SetPortraitUrl(_portraitUrl string) error {
    r._portraitUrl = _portraitUrl
    r.Set("portrait_url", _portraitUrl)
    return nil
}

// PortraitUrl Getter
func (r AlibabaAlihealthMdeerVideoSyncRequest) GetPortraitUrl() string {
    return r._portraitUrl
}
// AuthorName Setter
// 作者名称
func (r *AlibabaAlihealthMdeerVideoSyncRequest) SetAuthorName(_authorName string) error {
    r._authorName = _authorName
    r.Set("author_name", _authorName)
    return nil
}

// AuthorName Getter
func (r AlibabaAlihealthMdeerVideoSyncRequest) GetAuthorName() string {
    return r._authorName
}
// AuthorId Setter
// 作者id
func (r *AlibabaAlihealthMdeerVideoSyncRequest) SetAuthorId(_authorId string) error {
    r._authorId = _authorId
    r.Set("author_id", _authorId)
    return nil
}

// AuthorId Getter
func (r AlibabaAlihealthMdeerVideoSyncRequest) GetAuthorId() string {
    return r._authorId
}
// PartnerHomepage Setter
// 合作方主页
func (r *AlibabaAlihealthMdeerVideoSyncRequest) SetPartnerHomepage(_partnerHomepage string) error {
    r._partnerHomepage = _partnerHomepage
    r.Set("partner_homepage", _partnerHomepage)
    return nil
}

// PartnerHomepage Getter
func (r AlibabaAlihealthMdeerVideoSyncRequest) GetPartnerHomepage() string {
    return r._partnerHomepage
}
// PartnerName Setter
// 合作方名称
func (r *AlibabaAlihealthMdeerVideoSyncRequest) SetPartnerName(_partnerName string) error {
    r._partnerName = _partnerName
    r.Set("partner_name", _partnerName)
    return nil
}

// PartnerName Getter
func (r AlibabaAlihealthMdeerVideoSyncRequest) GetPartnerName() string {
    return r._partnerName
}
// ReleaseDate Setter
// 发布日期
func (r *AlibabaAlihealthMdeerVideoSyncRequest) SetReleaseDate(_releaseDate string) error {
    r._releaseDate = _releaseDate
    r.Set("release_date", _releaseDate)
    return nil
}

// ReleaseDate Getter
func (r AlibabaAlihealthMdeerVideoSyncRequest) GetReleaseDate() string {
    return r._releaseDate
}
// VideoFileUrl Setter
// 视频文件url
func (r *AlibabaAlihealthMdeerVideoSyncRequest) SetVideoFileUrl(_videoFileUrl string) error {
    r._videoFileUrl = _videoFileUrl
    r.Set("video_file_url", _videoFileUrl)
    return nil
}

// VideoFileUrl Getter
func (r AlibabaAlihealthMdeerVideoSyncRequest) GetVideoFileUrl() string {
    return r._videoFileUrl
}
// VideoMobileUrl Setter
// 视频落地页
func (r *AlibabaAlihealthMdeerVideoSyncRequest) SetVideoMobileUrl(_videoMobileUrl string) error {
    r._videoMobileUrl = _videoMobileUrl
    r.Set("video_mobile_url", _videoMobileUrl)
    return nil
}

// VideoMobileUrl Getter
func (r AlibabaAlihealthMdeerVideoSyncRequest) GetVideoMobileUrl() string {
    return r._videoMobileUrl
}
// VideoIntroduction Setter
// 视频简介
func (r *AlibabaAlihealthMdeerVideoSyncRequest) SetVideoIntroduction(_videoIntroduction string) error {
    r._videoIntroduction = _videoIntroduction
    r.Set("video_introduction", _videoIntroduction)
    return nil
}

// VideoIntroduction Getter
func (r AlibabaAlihealthMdeerVideoSyncRequest) GetVideoIntroduction() string {
    return r._videoIntroduction
}
// VideoLength Setter
// 视频长度
func (r *AlibabaAlihealthMdeerVideoSyncRequest) SetVideoLength(_videoLength string) error {
    r._videoLength = _videoLength
    r.Set("video_length", _videoLength)
    return nil
}

// VideoLength Getter
func (r AlibabaAlihealthMdeerVideoSyncRequest) GetVideoLength() string {
    return r._videoLength
}
// Disease Setter
// 视频所述疾病
func (r *AlibabaAlihealthMdeerVideoSyncRequest) SetDisease(_disease string) error {
    r._disease = _disease
    r.Set("disease", _disease)
    return nil
}

// Disease Getter
func (r AlibabaAlihealthMdeerVideoSyncRequest) GetDisease() string {
    return r._disease
}
// PriviewUrl Setter
// 预览图url
func (r *AlibabaAlihealthMdeerVideoSyncRequest) SetPriviewUrl(_priviewUrl string) error {
    r._priviewUrl = _priviewUrl
    r.Set("priview_url", _priviewUrl)
    return nil
}

// PriviewUrl Getter
func (r AlibabaAlihealthMdeerVideoSyncRequest) GetPriviewUrl() string {
    return r._priviewUrl
}
// VideoTitle Setter
// 视频标题
func (r *AlibabaAlihealthMdeerVideoSyncRequest) SetVideoTitle(_videoTitle string) error {
    r._videoTitle = _videoTitle
    r.Set("video_title", _videoTitle)
    return nil
}

// VideoTitle Getter
func (r AlibabaAlihealthMdeerVideoSyncRequest) GetVideoTitle() string {
    return r._videoTitle
}
// VideoId Setter
// 视频id
func (r *AlibabaAlihealthMdeerVideoSyncRequest) SetVideoId(_videoId string) error {
    r._videoId = _videoId
    r.Set("video_id", _videoId)
    return nil
}

// VideoId Getter
func (r AlibabaAlihealthMdeerVideoSyncRequest) GetVideoId() string {
    return r._videoId
}
