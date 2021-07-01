package alihealthmdeer

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthMdeerVideoSyncAPIRequest
合作伙伴视频同步给医知鹿接口 API请求
alibaba.alihealth.mdeer.video.sync

合伙做伴内容同步接口，用来视频同步 */
type AlibabaAlihealthMdeerVideoSyncAPIRequest struct {
	model.Params
	// 合作方头像url
	_partnerPortraitUrl string
	// 作者电话
	_phoneNumber string
	// 作者简介
	_authorIntroduction string
	// 作者科室
	_authorDepartment string
	// 作者级别
	_authorLevel string
	// 医院级别
	_hospitalLevel string
	// 医院名称
	_hospitalName string
	// 作者头像
	_portraitUrl string
	// 作者名称
	_authorName string
	// 作者id
	_authorId string
	// 合作方主页
	_partnerHomepage string
	// 合作方名称
	_partnerName string
	// 发布日期
	_releaseDate string
	// 视频文件url
	_videoFileUrl string
	// 视频落地页
	_videoMobileUrl string
	// 视频简介
	_videoIntroduction string
	// 视频长度
	_videoLength string
	// 视频所述疾病
	_disease string
	// 预览图url
	_priviewUrl string
	// 视频标题
	_videoTitle string
	// 视频id
	_videoId string
}

// New
