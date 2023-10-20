package tmallgenie

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaaicontentbusinesssendplanqueryAPIRequest 内容商业化发放权益查询 API请求
// alibaba.ai.content.business.send.plan.query
//
// 天猫精灵内容生态的权益查询
type AlibabaaicontentbusinesssendplanqueryAPIRequest struct {
	model.Params
	// 入参对象
	_param0 *PurchaseForOuterDto
}

// NewAlibabaaicontentbusinesssendplanqueryRequest 初始化AlibabaaicontentbusinesssendplanqueryAPIRequest对象
func NewAlibabaaicontentbusinesssendplanqueryRequest() *AlibabaaicontentbusinesssendplanqueryAPIRequest {
	return &AlibabaaicontentbusinesssendplanqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaaicontentbusinesssendplanqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.ai.content.business.send.plan.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaaicontentbusinesssendplanqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaaicontentbusinesssendplanqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 入参对象
func (r *AlibabaaicontentbusinesssendplanqueryAPIRequest) SetParam0(_param0 *PurchaseForOuterDto) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaaicontentbusinesssendplanqueryAPIRequest) GetParam0() *PurchaseForOuterDto {
	return r._param0
}
