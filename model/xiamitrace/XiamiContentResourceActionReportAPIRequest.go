package xiamitrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* XiamiContentResourceActionReportAPIRequest
曲库开放平台内容行为上报接口 API请求
xiami.content.resource.action.report

合作方对接入的曲库开放内容上报行为日志 */
type XiamiContentResourceActionReportAPIRequest struct {
	model.Params
	// 资源ID
	_resourceId string
	// 行为数量
	_num int64
	// 资源类型（可枚举）: song(歌曲)
	_resourceType string
	// 行为类型（可枚举）：LISTEN（主动试听）、PASSIVE_LISTEN（被动试听）
	_action string
	// 来源id，如歌单id
	_fromId string
	// 1推荐2歌单3标签
	_fromType int64
	// 用户id
	_openId string
	// 用户设备id
	_utdid string
}

// New
