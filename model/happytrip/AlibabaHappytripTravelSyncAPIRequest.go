package happytrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaHappytripTravelSyncAPIRequest
差旅申请单同步接口 API请求
alibaba.happytrip.travel.sync

以外部差旅申请单id（outer_travel_head_id）为主键，保存或更新差旅单信息到欢行系统中 */
type AlibabaHappytripTravelSyncAPIRequest struct {
	model.Params
	// 差旅申请单对象
	_travelHeadDto *TravelHeadDto
}

// New
