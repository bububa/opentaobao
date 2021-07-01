package alitripbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripBpCouponinfoSyncAPIRequest
飞猪广告券信息同步接口 API请求
alitrip.bp.couponinfo.sync

飞猪商业化券信息同步 */
type AlitripBpCouponinfoSyncAPIRequest struct {
	model.Params
	// 商业化券同步接口请求
	_paramCouponDataRequest *CouponDataRequest
}

// New
