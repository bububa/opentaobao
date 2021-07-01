package lstlogistics

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaLstShiporderCreateAPIRequest
零售通发货单创建 API请求
alibaba.lst.shiporder.create

通过该接口可以创建零售通运保保发货单，并处理相关业务流程。 */
type AlibabaLstShiporderCreateAPIRequest struct {
	model.Params
	// 创建发货单入参
	_shipOrder *LstThirdPartMainShipOrderCreateDto
}

// New
