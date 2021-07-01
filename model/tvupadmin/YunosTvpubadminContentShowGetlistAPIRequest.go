package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosTvpubadminContentShowGetlistAPIRequest
节目审核获取节目列表 API请求
yunos.tvpubadmin.content.show.getlist

节目审核获取节目列表 */
type YunosTvpubadminContentShowGetlistAPIRequest struct {
	model.Params
	// 视频外部来源类型: 1:YOUKU, 2:MONGO_TV, 3:TAOTVMEDIA, 4:GOLIVE
	_extType int64
	// 时间查询范围，结束时间
	_gmtEnd string
	// 审核状态：1未提审，2审核中，3通过，4不通过，5已下线
	_licenseState int64
	// 分页
	_pageSize int64
	// 节目ID
	_showId string
	// 视频ID
	_extVideoStrId string
	// 时间类型：1-licenseSubmitTime, 2-licenseAuditTime, 3-youkuPublishTime
	_dateType int64
	// 主分类
	_category int64
	// 节目名称
	_showName string
	// 分页，页码
	_pageNo int64
	// 时间查询范围，开始时间
	_gmtStart string
	// 牌照方
	_license int64
	// 视频名称
	_videoTitleLike string
	// 老媒资ID(program_id)
	_vmacLongId int64
}

// New
