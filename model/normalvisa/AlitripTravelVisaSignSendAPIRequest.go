package normalvisa

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripTravelVisaSignSendAPIRequest
签证批量申请人送签接口 API请求
alitrip.travel.visa.sign.send

签证批量申请人送签接口，用于商家批量送签。 */
type AlitripTravelVisaSignSendAPIRequest struct {
	model.Params
	// 国家id。目前只支持越南，越南国家id:27027
	_nationId int64
	// 送签类型：1-非加急，2-加急，默认非加急
	_signType int64
	// 申请人ids
	_applyIds []string
}

// New
