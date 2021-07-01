package normalvisa

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripTravelNormalvisaStoreuserAPIRequest
代填办理人信息 API请求
taobao.alitrip.travel.normalvisa.storeuser

卖家代填买家填写办理人信息 */
type TaobaoAlitripTravelNormalvisaStoreuserAPIRequest struct {
	model.Params
	// 订单id
	_bizOrderId int64
	// 列表：签证人信息列表
	_normalVisaUserUnitList []NormalVisaUserUnit
}

// New
