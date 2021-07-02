package vaccin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHealthVaccinNoticeAnnouncementPublishAPIRequest 支付宝疫苗POV公告通知 API请求
// alibaba.health.vaccin.notice.announcement.publish
//
// 支付宝疫苗POV发布公告提醒信息
type AlibabaHealthVaccinNoticeAnnouncementPublishAPIRequest struct {
	model.Params
	// 接种点编码（通知方）
	_povNo string
	// 公告标题
	_noticeTitle string
	// 公告内容
	_noticeContent string
	// 0:所有类型人群,1：宝宝、2：成人
	_noticeType string
	// 公告发布时间
	_noticeTime string
	// 需要接受公告的用户ID，默认是所有用户都可以看到
	_alipayUserIds []string
	// 接种点名称
	_povName string
}

// NewAlibabaHealthVaccinNoticeAnnouncementPublishRequest 初始化AlibabaHealthVaccinNoticeAnnouncementPublishAPIRequest对象
func NewAlibabaHealthVaccinNoticeAnnouncementPublishRequest() *AlibabaHealthVaccinNoticeAnnouncementPublishAPIRequest {
	return &AlibabaHealthVaccinNoticeAnnouncementPublishAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaHealthVaccinNoticeAnnouncementPublishAPIRequest) GetApiMethodName() string {
	return "alibaba.health.vaccin.notice.announcement.publish"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaHealthVaccinNoticeAnnouncementPublishAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetPovNo is PovNo Setter
// 接种点编码（通知方）
func (r *AlibabaHealthVaccinNoticeAnnouncementPublishAPIRequest) SetPovNo(_povNo string) error {
	r._povNo = _povNo
	r.Set("pov_no", _povNo)
	return nil
}

// GetPovNo PovNo Getter
func (r AlibabaHealthVaccinNoticeAnnouncementPublishAPIRequest) GetPovNo() string {
	return r._povNo
}

// SetNoticeTitle is NoticeTitle Setter
// 公告标题
func (r *AlibabaHealthVaccinNoticeAnnouncementPublishAPIRequest) SetNoticeTitle(_noticeTitle string) error {
	r._noticeTitle = _noticeTitle
	r.Set("notice_title", _noticeTitle)
	return nil
}

// GetNoticeTitle NoticeTitle Getter
func (r AlibabaHealthVaccinNoticeAnnouncementPublishAPIRequest) GetNoticeTitle() string {
	return r._noticeTitle
}

// SetNoticeContent is NoticeContent Setter
// 公告内容
func (r *AlibabaHealthVaccinNoticeAnnouncementPublishAPIRequest) SetNoticeContent(_noticeContent string) error {
	r._noticeContent = _noticeContent
	r.Set("notice_content", _noticeContent)
	return nil
}

// GetNoticeContent NoticeContent Getter
func (r AlibabaHealthVaccinNoticeAnnouncementPublishAPIRequest) GetNoticeContent() string {
	return r._noticeContent
}

// SetNoticeType is NoticeType Setter
// 0:所有类型人群,1：宝宝、2：成人
func (r *AlibabaHealthVaccinNoticeAnnouncementPublishAPIRequest) SetNoticeType(_noticeType string) error {
	r._noticeType = _noticeType
	r.Set("notice_type", _noticeType)
	return nil
}

// GetNoticeType NoticeType Getter
func (r AlibabaHealthVaccinNoticeAnnouncementPublishAPIRequest) GetNoticeType() string {
	return r._noticeType
}

// SetNoticeTime is NoticeTime Setter
// 公告发布时间
func (r *AlibabaHealthVaccinNoticeAnnouncementPublishAPIRequest) SetNoticeTime(_noticeTime string) error {
	r._noticeTime = _noticeTime
	r.Set("notice_time", _noticeTime)
	return nil
}

// GetNoticeTime NoticeTime Getter
func (r AlibabaHealthVaccinNoticeAnnouncementPublishAPIRequest) GetNoticeTime() string {
	return r._noticeTime
}

// SetAlipayUserIds is AlipayUserIds Setter
// 需要接受公告的用户ID，默认是所有用户都可以看到
func (r *AlibabaHealthVaccinNoticeAnnouncementPublishAPIRequest) SetAlipayUserIds(_alipayUserIds []string) error {
	r._alipayUserIds = _alipayUserIds
	r.Set("alipay_user_ids", _alipayUserIds)
	return nil
}

// GetAlipayUserIds AlipayUserIds Getter
func (r AlibabaHealthVaccinNoticeAnnouncementPublishAPIRequest) GetAlipayUserIds() []string {
	return r._alipayUserIds
}

// SetPovName is PovName Setter
// 接种点名称
func (r *AlibabaHealthVaccinNoticeAnnouncementPublishAPIRequest) SetPovName(_povName string) error {
	r._povName = _povName
	r.Set("pov_name", _povName)
	return nil
}

// GetPovName PovName Getter
func (r AlibabaHealthVaccinNoticeAnnouncementPublishAPIRequest) GetPovName() string {
	return r._povName
}
