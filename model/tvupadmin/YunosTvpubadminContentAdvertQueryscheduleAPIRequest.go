package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosTvpubadminContentAdvertQueryscheduleAPIRequest
广告牌照管控查询 API请求
yunos.tvpubadmin.content.advert.queryschedule

广告牌照管控查询 */
type YunosTvpubadminContentAdvertQueryscheduleAPIRequest struct {
	model.Params
	// 查询范围: 1-牌照，4-uuid
	_range int64
	// 分页，页码
	_pageNo int64
	// 分页，单页数量
	_pageSize int64
	// 日期
	_gmtStart string
	// uuid
	_uuid string
	// 牌照方，1-华数，7-CIBN
	_license int64
	// 广告类型
	_sityType int64
}

// New
