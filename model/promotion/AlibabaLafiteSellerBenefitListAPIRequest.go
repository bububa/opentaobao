package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabalafitesellerbenefitlistAPIRequest 商家自运营权益列表 API请求
// alibaba.lafite.seller.benefit.list
//
// 小程序isv可使用该接口获取权益列表
type AlibabalafitesellerbenefitlistAPIRequest struct {
	model.Params
	// 查询参数
	_benefitReadTopQuery *BenefitReadTopQuery
}

// NewAlibabalafitesellerbenefitlistRequest 初始化AlibabalafitesellerbenefitlistAPIRequest对象
func NewAlibabalafitesellerbenefitlistRequest() *AlibabalafitesellerbenefitlistAPIRequest {
	return &AlibabalafitesellerbenefitlistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabalafitesellerbenefitlistAPIRequest) GetApiMethodName() string {
	return "alibaba.lafite.seller.benefit.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabalafitesellerbenefitlistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabalafitesellerbenefitlistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBenefitReadTopQuery is BenefitReadTopQuery Setter
// 查询参数
func (r *AlibabalafitesellerbenefitlistAPIRequest) SetBenefitReadTopQuery(_benefitReadTopQuery *BenefitReadTopQuery) error {
	r._benefitReadTopQuery = _benefitReadTopQuery
	r.Set("benefit_read_top_query", _benefitReadTopQuery)
	return nil
}

// GetBenefitReadTopQuery BenefitReadTopQuery Getter
func (r AlibabalafitesellerbenefitlistAPIRequest) GetBenefitReadTopQuery() *BenefitReadTopQuery {
	return r._benefitReadTopQuery
}
