package vaccin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
支付宝疫苗POV公告通知 APIRequest
alibaba.health.vaccin.notice.announcement.publish

支付宝疫苗POV发布公告提醒信息
*/
type AlibabaHealthVaccinNoticeAnnouncementPublishRequest struct {
    model.Params

    // 接种点编码（通知方）
    povNo   string 

    // 公告标题
    noticeTitle   string 

    // 公告内容
    noticeContent   string 

    // 0:所有类型人群,1：宝宝、2：成人
    noticeType   string 

    // 公告发布时间
    noticeTime   string 

    // 需要接受公告的用户ID，默认是所有用户都可以看到
    alipayUserIds   []string 

    // 接种点名称
    povName   string 

}

func NewAlibabaHealthVaccinNoticeAnnouncementPublishRequest() *AlibabaHealthVaccinNoticeAnnouncementPublishRequest{
    return &AlibabaHealthVaccinNoticeAnnouncementPublishRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaHealthVaccinNoticeAnnouncementPublishRequest) GetApiMethodName() string {
    return "alibaba.health.vaccin.notice.announcement.publish"
}

func (r AlibabaHealthVaccinNoticeAnnouncementPublishRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaHealthVaccinNoticeAnnouncementPublishRequest) SetPovNo(povNo string) error {
    r.povNo = povNo
    r.Set("pov_no", povNo)
    return nil
}

func (r AlibabaHealthVaccinNoticeAnnouncementPublishRequest) GetPovNo() string {
    return r.povNo
}

func (r *AlibabaHealthVaccinNoticeAnnouncementPublishRequest) SetNoticeTitle(noticeTitle string) error {
    r.noticeTitle = noticeTitle
    r.Set("notice_title", noticeTitle)
    return nil
}

func (r AlibabaHealthVaccinNoticeAnnouncementPublishRequest) GetNoticeTitle() string {
    return r.noticeTitle
}

func (r *AlibabaHealthVaccinNoticeAnnouncementPublishRequest) SetNoticeContent(noticeContent string) error {
    r.noticeContent = noticeContent
    r.Set("notice_content", noticeContent)
    return nil
}

func (r AlibabaHealthVaccinNoticeAnnouncementPublishRequest) GetNoticeContent() string {
    return r.noticeContent
}

func (r *AlibabaHealthVaccinNoticeAnnouncementPublishRequest) SetNoticeType(noticeType string) error {
    r.noticeType = noticeType
    r.Set("notice_type", noticeType)
    return nil
}

func (r AlibabaHealthVaccinNoticeAnnouncementPublishRequest) GetNoticeType() string {
    return r.noticeType
}

func (r *AlibabaHealthVaccinNoticeAnnouncementPublishRequest) SetNoticeTime(noticeTime string) error {
    r.noticeTime = noticeTime
    r.Set("notice_time", noticeTime)
    return nil
}

func (r AlibabaHealthVaccinNoticeAnnouncementPublishRequest) GetNoticeTime() string {
    return r.noticeTime
}

func (r *AlibabaHealthVaccinNoticeAnnouncementPublishRequest) SetAlipayUserIds(alipayUserIds []string) error {
    r.alipayUserIds = alipayUserIds
    r.Set("alipay_user_ids", alipayUserIds)
    return nil
}

func (r AlibabaHealthVaccinNoticeAnnouncementPublishRequest) GetAlipayUserIds() []string {
    return r.alipayUserIds
}

func (r *AlibabaHealthVaccinNoticeAnnouncementPublishRequest) SetPovName(povName string) error {
    r.povName = povName
    r.Set("pov_name", povName)
    return nil
}

func (r AlibabaHealthVaccinNoticeAnnouncementPublishRequest) GetPovName() string {
    return r.povName
}

