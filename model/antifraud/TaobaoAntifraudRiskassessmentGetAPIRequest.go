package antifraud

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAntifraudRiskassessmentGetAPIRequest 反欺诈风险识别 API请求
// taobao.antifraud.riskassessment.get
//
// 反欺诈服务是阿里大数据风控服务能力的对外输出，通过用户信誉、行为分析精准识别可信用户和风险用户并实时防御，解决交易、支付、活动等关键业务环节存在的欺诈威胁，保护企业品牌和数据，降低企业经济损失
type TaobaoAntifraudRiskassessmentGetAPIRequest struct {
	model.Params
	// 风控查询参数
	_collinadataContext *CollinadataContext
}

// NewTaobaoAntifraudRiskassessmentGetRequest 初始化TaobaoAntifraudRiskassessmentGetAPIRequest对象
func NewTaobaoAntifraudRiskassessmentGetRequest() *TaobaoAntifraudRiskassessmentGetAPIRequest {
	return &TaobaoAntifraudRiskassessmentGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAntifraudRiskassessmentGetAPIRequest) GetApiMethodName() string {
	return "taobao.antifraud.riskassessment.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAntifraudRiskassessmentGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetCollinadataContext is CollinadataContext Setter
// 风控查询参数
func (r *TaobaoAntifraudRiskassessmentGetAPIRequest) SetCollinadataContext(_collinadataContext *CollinadataContext) error {
	r._collinadataContext = _collinadataContext
	r.Set("collinadata_context", _collinadataContext)
	return nil
}

// GetCollinadataContext CollinadataContext Getter
func (r TaobaoAntifraudRiskassessmentGetAPIRequest) GetCollinadataContext() *CollinadataContext {
	return r._collinadataContext
}
