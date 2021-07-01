package traveltrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripTravelBookinfoQueryAPIRequest
飞猪度假-订单二次预约查询接口 API请求
alitrip.travel.bookinfo.query

飞猪度假订单二次预约详情查询接口 */
type AlitripTravelBookinfoQueryAPIRequest struct {
	model.Params
	// 预定信息id
	_bookinfoId int64
}

// New
