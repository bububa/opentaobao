package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosTvpubadminContentSingleVideoGetlistAPIRequest
单视频审核获取视频列表 API请求
yunos.tvpubadmin.content.single.video.getlist

牌照方在审核单视频时获取单视频列表接口 */
type YunosTvpubadminContentSingleVideoGetlistAPIRequest struct {
	model.Params
	// 视频外部来源类型: 1:YOUKU, 2:MONGO_TV, 3:TAOTVMEDIA, 4:GOLIVE
	_extType int64
	// 审核状态：1未提审，2审核中，3通过，4不通过，5已下线
	_licenseState int64
	// 单页数量
	_pageSize int64
	// 查询时间范围，结束时间
	_gmtEnd string
	// 视频id
	_extVideoStrId string
	// 查询多个审核状态
	_licenseStateList []int64
	// 时间类型：1-licenseSubmitTime, 2-licenseAuditTime, 3-youkuPublishTime
	_dateType int64
	// 主分类
	_category int64
	// 页码
	_pageNo int64
	// 查询时间范围，开始时间
	_gmtStart string
	// 牌照方
	_license int64
	// 视屏名称
	_videoTitleLike string
	// 审核优先级，紧急4，高3，中2，低1
	_priority int64
}

// New
