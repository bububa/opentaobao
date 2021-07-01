package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSubwayMarshLandRptGetAPIRequest
获取捡漏词包分时报表数据 API请求
taobao.subway.marsh.land.rpt.get

获取捡漏词包分时报表数据 */
type TaobaoSubwayMarshLandRptGetAPIRequest struct {
	model.Params
	// 主人昵称
	_nick string
	// 2021-05-11
	_endDate string
	// 推广组id
	_adgroupIdEqual string
	// 词包类型（捡漏词包填19）
	_isAutoMatchEqual string
	// 计划id
	_campaignIdEqual string
	// 2021-05-05
	_startDate string
}

// New
