package tmallgenie

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaaicontentbusinesssendplanreceiveAPIRequest 天猫精灵商业化采销发放计划领取 API请求
// alibaba.ai.content.business.send.plan.receive
//
// 天猫精灵商业化采销发放计划领取
type AlibabaaicontentbusinesssendplanreceiveAPIRequest struct {
	model.Params
	// 入参对象
	_param0 *PurchaseForOuterDto
}

// NewAlibabaaicontentbusinesssendplanreceiveRequest 初始化AlibabaaicontentbusinesssendplanreceiveAPIRequest对象
func NewAlibabaaicontentbusinesssendplanreceiveRequest() *AlibabaaicontentbusinesssendplanreceiveAPIRequest {
	return &AlibabaaicontentbusinesssendplanreceiveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaaicontentbusinesssendplanreceiveAPIRequest) GetApiMethodName() string {
	return "alibaba.ai.content.business.send.plan.receive"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaaicontentbusinesssendplanreceiveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaaicontentbusinesssendplanreceiveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 入参对象
func (r *AlibabaaicontentbusinesssendplanreceiveAPIRequest) SetParam0(_param0 *PurchaseForOuterDto) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaaicontentbusinesssendplanreceiveAPIRequest) GetParam0() *PurchaseForOuterDto {
	return r._param0
}
