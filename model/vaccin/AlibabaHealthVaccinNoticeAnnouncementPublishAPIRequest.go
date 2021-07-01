package vaccin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaHealthVaccinNoticeAnnouncementPublishAPIRequest
支付宝疫苗POV公告通知 API请求
alibaba.health.vaccin.notice.announcement.publish

支付宝疫苗POV发布公告提醒信息 */
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

// New
