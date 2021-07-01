package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOmniorderDtdConsumeAPIRequest
门店自送对码进行核销 API请求
taobao.omniorder.dtd.consume

该接口根据传入的码及订单信息，如果码与订单一致，则对门店自送服务进行核销。 */
type TaobaoOmniorderDtdConsumeAPIRequest struct {
	model.Params
	// 核销信息
	_paramDoor2doorConsumeRequest *Door2doorConsumeRequest
}

// New
