package normalvisa

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripTravelNormalvisaUpdatepersonstautsAPIRequest
更新签证办理进度 API请求
taobao.alitrip.travel.normalvisa.updatepersonstauts

更新签证办理进度 */
type TaobaoAlitripTravelNormalvisaUpdatepersonstautsAPIRequest struct {
	model.Params
	// 更新信息
	_normalVisaUpdateUnitList []NormalVisaUpdateUnit
	// 订单号
	_bizOrderId int64
}

// New
