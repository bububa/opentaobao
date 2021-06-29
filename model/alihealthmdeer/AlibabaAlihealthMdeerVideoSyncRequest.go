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
    partnerPortraitUrl   string
    // 作者电话
    phoneNumber   string
    // 作者简介
    authorIntroduction   string
    // 作者科室
    authorDepartment   string
    // 作者级别
    authorLevel   string
    // 医院级别
    hospitalLevel   string
    // 医院名称
    hospitalName   string
    // 作者头像
    portraitUrl   string
    // 作者名称
    authorName   string
    // 作者id
    authorId   string
    // 合作方主页
    partnerHomepage   string
    // 合作方名称
    partnerName   string
    // 发布日期
    releaseDate   string
    // 视频文件url
    videoFileUrl   string
    // 视频落地页
    videoMobileUrl   string
    // 视频简介
    videoIntroduction   string
    // 视频长度
    videoLength   string
    // 视频所述疾病
    disease   string
    // 预览图url
    priviewUrl   string
    // 视频标题
    videoTitle   string
    // 视频id
    videoId   string
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
func (r *AlibabaAlihealthMdeerVideoSyncRequest) SetPartnerPortraitUrl(partnerPortraitUrl string) error {
    r.partnerPortraitUrl = partnerPortraitUrl
    r.Set("partner_portrait_url", partnerPortraitUrl)
    return nil
}

// PartnerPortraitUrl Getter
func (r AlibabaAlihealthMdeerVideoSyncRequest) GetPartnerPortraitUrl() string {
    return r.partnerPortraitUrl
}
// PhoneNumber Setter
// 作者电话
func (r *AlibabaAlihealthMdeerVideoSyncRequest) SetPhoneNumber(phoneNumber string) error {
    r.phoneNumber = phoneNumber
    r.Set("phone_number", phoneNumber)
    return nil
}

// PhoneNumber Getter
func (r AlibabaAlihealthMdeerVideoSyncRequest) GetPhoneNumber() string {
    return r.phoneNumber
}
// AuthorIntroduction Setter
// 作者简介
func (r *AlibabaAlihealthMdeerVideoSyncRequest) SetAuthorIntroduction(authorIntroduction string) error {
    r.authorIntroduction = authorIntroduction
    r.Set("author_introduction", authorIntroduction)
    return nil
}

// AuthorIntroduction Getter
func (r AlibabaAlihealthMdeerVideoSyncRequest) GetAuthorIntroduction() string {
    return r.authorIntroduction
}
// AuthorDepartment Setter
// 作者科室
func (r *AlibabaAlihealthMdeerVideoSyncRequest) SetAuthorDepartment(authorDepartment string) error {
    r.authorDepartment = authorDepartment
    r.Set("author_department", authorDepartment)
    return nil
}

// AuthorDepartment Getter
func (r AlibabaAlihealthMdeerVideoSyncRequest) GetAuthorDepartment() string {
    return r.authorDepartment
}
// AuthorLevel Setter
// 作者级别
func (r *AlibabaAlihealthMdeerVideoSyncRequest) SetAuthorLevel(authorLevel string) error {
    r.authorLevel = authorLevel
    r.Set("author_level", authorLevel)
    return nil
}

// AuthorLevel Getter
func (r AlibabaAlihealthMdeerVideoSyncRequest) GetAuthorLevel() string {
    return r.authorLevel
}
// HospitalLevel Setter
// 医院级别
func (r *AlibabaAlihealthMdeerVideoSyncRequest) SetHospitalLevel(hospitalLevel string) error {
    r.hospitalLevel = hospitalLevel
    r.Set("hospital_level", hospitalLevel)
    return nil
}

// HospitalLevel Getter
func (r AlibabaAlihealthMdeerVideoSyncRequest) GetHospitalLevel() string {
    return r.hospitalLevel
}
// HospitalName Setter
// 医院名称
func (r *AlibabaAlihealthMdeerVideoSyncRequest) SetHospitalName(hospitalName string) error {
    r.hospitalName = hospitalName
    r.Set("hospital_name", hospitalName)
    return nil
}

// HospitalName Getter
func (r AlibabaAlihealthMdeerVideoSyncRequest) GetHospitalName() string {
    return r.hospitalName
}
// PortraitUrl Setter
// 作者头像
func (r *AlibabaAlihealthMdeerVideoSyncRequest) SetPortraitUrl(portraitUrl string) error {
    r.portraitUrl = portraitUrl
    r.Set("portrait_url", portraitUrl)
    return nil
}

// PortraitUrl Getter
func (r AlibabaAlihealthMdeerVideoSyncRequest) GetPortraitUrl() string {
    return r.portraitUrl
}
// AuthorName Setter
// 作者名称
func (r *AlibabaAlihealthMdeerVideoSyncRequest) SetAuthorName(authorName string) error {
    r.authorName = authorName
    r.Set("author_name", authorName)
    return nil
}

// AuthorName Getter
func (r AlibabaAlihealthMdeerVideoSyncRequest) GetAuthorName() string {
    return r.authorName
}
// AuthorId Setter
// 作者id
func (r *AlibabaAlihealthMdeerVideoSyncRequest) SetAuthorId(authorId string) error {
    r.authorId = authorId
    r.Set("author_id", authorId)
    return nil
}

// AuthorId Getter
func (r AlibabaAlihealthMdeerVideoSyncRequest) GetAuthorId() string {
    return r.authorId
}
// PartnerHomepage Setter
// 合作方主页
func (r *AlibabaAlihealthMdeerVideoSyncRequest) SetPartnerHomepage(partnerHomepage string) error {
    r.partnerHomepage = partnerHomepage
    r.Set("partner_homepage", partnerHomepage)
    return nil
}

// PartnerHomepage Getter
func (r AlibabaAlihealthMdeerVideoSyncRequest) GetPartnerHomepage() string {
    return r.partnerHomepage
}
// PartnerName Setter
// 合作方名称
func (r *AlibabaAlihealthMdeerVideoSyncRequest) SetPartnerName(partnerName string) error {
    r.partnerName = partnerName
    r.Set("partner_name", partnerName)
    return nil
}

// PartnerName Getter
func (r AlibabaAlihealthMdeerVideoSyncRequest) GetPartnerName() string {
    return r.partnerName
}
// ReleaseDate Setter
// 发布日期
func (r *AlibabaAlihealthMdeerVideoSyncRequest) SetReleaseDate(releaseDate string) error {
    r.releaseDate = releaseDate
    r.Set("release_date", releaseDate)
    return nil
}

// ReleaseDate Getter
func (r AlibabaAlihealthMdeerVideoSyncRequest) GetReleaseDate() string {
    return r.releaseDate
}
// VideoFileUrl Setter
// 视频文件url
func (r *AlibabaAlihealthMdeerVideoSyncRequest) SetVideoFileUrl(videoFileUrl string) error {
    r.videoFileUrl = videoFileUrl
    r.Set("video_file_url", videoFileUrl)
    return nil
}

// VideoFileUrl Getter
func (r AlibabaAlihealthMdeerVideoSyncRequest) GetVideoFileUrl() string {
    return r.videoFileUrl
}
// VideoMobileUrl Setter
// 视频落地页
func (r *AlibabaAlihealthMdeerVideoSyncRequest) SetVideoMobileUrl(videoMobileUrl string) error {
    r.videoMobileUrl = videoMobileUrl
    r.Set("video_mobile_url", videoMobileUrl)
    return nil
}

// VideoMobileUrl Getter
func (r AlibabaAlihealthMdeerVideoSyncRequest) GetVideoMobileUrl() string {
    return r.videoMobileUrl
}
// VideoIntroduction Setter
// 视频简介
func (r *AlibabaAlihealthMdeerVideoSyncRequest) SetVideoIntroduction(videoIntroduction string) error {
    r.videoIntroduction = videoIntroduction
    r.Set("video_introduction", videoIntroduction)
    return nil
}

// VideoIntroduction Getter
func (r AlibabaAlihealthMdeerVideoSyncRequest) GetVideoIntroduction() string {
    return r.videoIntroduction
}
// VideoLength Setter
// 视频长度
func (r *AlibabaAlihealthMdeerVideoSyncRequest) SetVideoLength(videoLength string) error {
    r.videoLength = videoLength
    r.Set("video_length", videoLength)
    return nil
}

// VideoLength Getter
func (r AlibabaAlihealthMdeerVideoSyncRequest) GetVideoLength() string {
    return r.videoLength
}
// Disease Setter
// 视频所述疾病
func (r *AlibabaAlihealthMdeerVideoSyncRequest) SetDisease(disease string) error {
    r.disease = disease
    r.Set("disease", disease)
    return nil
}

// Disease Getter
func (r AlibabaAlihealthMdeerVideoSyncRequest) GetDisease() string {
    return r.disease
}
// PriviewUrl Setter
// 预览图url
func (r *AlibabaAlihealthMdeerVideoSyncRequest) SetPriviewUrl(priviewUrl string) error {
    r.priviewUrl = priviewUrl
    r.Set("priview_url", priviewUrl)
    return nil
}

// PriviewUrl Getter
func (r AlibabaAlihealthMdeerVideoSyncRequest) GetPriviewUrl() string {
    return r.priviewUrl
}
// VideoTitle Setter
// 视频标题
func (r *AlibabaAlihealthMdeerVideoSyncRequest) SetVideoTitle(videoTitle string) error {
    r.videoTitle = videoTitle
    r.Set("video_title", videoTitle)
    return nil
}

// VideoTitle Getter
func (r AlibabaAlihealthMdeerVideoSyncRequest) GetVideoTitle() string {
    return r.videoTitle
}
// VideoId Setter
// 视频id
func (r *AlibabaAlihealthMdeerVideoSyncRequest) SetVideoId(videoId string) error {
    r.videoId = videoId
    r.Set("video_id", videoId)
    return nil
}

// VideoId Getter
func (r AlibabaAlihealthMdeerVideoSyncRequest) GetVideoId() string {
    return r.videoId
}
