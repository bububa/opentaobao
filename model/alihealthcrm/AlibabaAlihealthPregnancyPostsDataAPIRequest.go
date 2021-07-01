package alihealthcrm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthPregnancyPostsDataAPIRequest
发回帖子信息同步 API请求
alibaba.alihealth.pregnancy.posts.data

发回帖子信息同步 */
type AlibabaAlihealthPregnancyPostsDataAPIRequest struct {
	model.Params
	// 用户id
	_userId int64
	// 事件类型 0发帖 1回帖
	_eventType int64
	// 主贴id
	_mainId int64
	// 回帖id
	_replyId int64
	// 标题
	_title string
	// 内容
	_content string
	// 图片url
	_picUrl string
	// 发帖时间
	_date int64
}

// New
