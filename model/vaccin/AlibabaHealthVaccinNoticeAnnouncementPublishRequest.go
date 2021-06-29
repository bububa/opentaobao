package vaccin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
支付宝疫苗POV公告通知 API请求
alibaba.health.vaccin.notice.announcement.publish

支付宝疫苗POV发布公告提醒信息
*/
type AlibabaHealthVaccinNoticeAnnouncementPublishRequest struct {
    model.Params
    // 接种点编码（通知方）
    _povNo   string
    // 公告标题
    _noticeTitle   string
    // 公告内容
    _noticeContent   string
    // 0:所有类型人群,1：宝宝、2：成人
    _noticeType   string
    // 公告发布时间
    _noticeTime   string
    // 需要接受公告的用户ID，默认是所有用户都可以看到
    _alipayUserIds   []string
    // 接种点名称
    _povName   string
}

// 初始化AlibabaHealthVaccinNoticeAnnouncementPublishRequest对象
func NewAlibabaHealthVaccinNoticeAnnouncementPublishRequest() *AlibabaHealthVaccinNoticeAnnouncementPublishRequest{
    return &AlibabaHealthVaccinNoticeAnnouncementPublishRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaHealthVaccinNoticeAnnouncementPublishRequest) GetApiMethodName() string {
    return "alibaba.health.vaccin.notice.announcement.publish"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaHealthVaccinNoticeAnnouncementPublishRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PovNo Setter
// 接种点编码（通知方）
func (r *AlibabaHealthVaccinNoticeAnnouncementPublishRequest) SetPovNo(_povNo string) error {
    r._povNo = _povNo
    r.Set("pov_no", _povNo)
    return nil
}

// PovNo Getter
func (r AlibabaHealthVaccinNoticeAnnouncementPublishRequest) GetPovNo() string {
    return r._povNo
}
// NoticeTitle Setter
// 公告标题
func (r *AlibabaHealthVaccinNoticeAnnouncementPublishRequest) SetNoticeTitle(_noticeTitle string) error {
    r._noticeTitle = _noticeTitle
    r.Set("notice_title", _noticeTitle)
    return nil
}

// NoticeTitle Getter
func (r AlibabaHealthVaccinNoticeAnnouncementPublishRequest) GetNoticeTitle() string {
    return r._noticeTitle
}
// NoticeContent Setter
// 公告内容
func (r *AlibabaHealthVaccinNoticeAnnouncementPublishRequest) SetNoticeContent(_noticeContent string) error {
    r._noticeContent = _noticeContent
    r.Set("notice_content", _noticeContent)
    return nil
}

// NoticeContent Getter
func (r AlibabaHealthVaccinNoticeAnnouncementPublishRequest) GetNoticeContent() string {
    return r._noticeContent
}
// NoticeType Setter
// 0:所有类型人群,1：宝宝、2：成人
func (r *AlibabaHealthVaccinNoticeAnnouncementPublishRequest) SetNoticeType(_noticeType string) error {
    r._noticeType = _noticeType
    r.Set("notice_type", _noticeType)
    return nil
}

// NoticeType Getter
func (r AlibabaHealthVaccinNoticeAnnouncementPublishRequest) GetNoticeType() string {
    return r._noticeType
}
// NoticeTime Setter
// 公告发布时间
func (r *AlibabaHealthVaccinNoticeAnnouncementPublishRequest) SetNoticeTime(_noticeTime string) error {
    r._noticeTime = _noticeTime
    r.Set("notice_time", _noticeTime)
    return nil
}

// NoticeTime Getter
func (r AlibabaHealthVaccinNoticeAnnouncementPublishRequest) GetNoticeTime() string {
    return r._noticeTime
}
// AlipayUserIds Setter
// 需要接受公告的用户ID，默认是所有用户都可以看到
func (r *AlibabaHealthVaccinNoticeAnnouncementPublishRequest) SetAlipayUserIds(_alipayUserIds []string) error {
    r._alipayUserIds = _alipayUserIds
    r.Set("alipay_user_ids", _alipayUserIds)
    return nil
}

// AlipayUserIds Getter
func (r AlibabaHealthVaccinNoticeAnnouncementPublishRequest) GetAlipayUserIds() []string {
    return r._alipayUserIds
}
// PovName Setter
// 接种点名称
func (r *AlibabaHealthVaccinNoticeAnnouncementPublishRequest) SetPovName(_povName string) error {
    r._povName = _povName
    r.Set("pov_name", _povName)
    return nil
}

// PovName Getter
func (r AlibabaHealthVaccinNoticeAnnouncementPublishRequest) GetPovName() string {
    return r._povName
}
