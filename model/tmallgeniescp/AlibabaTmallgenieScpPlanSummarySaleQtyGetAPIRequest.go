package tmallgeniescp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabatmallgeniescpplansummarysaleqtygetAPIRequest 同步销售数据按照渠道类型汇总 API请求
// alibaba.tmallgenie.scp.plan.summary.sale.qty.get
//
// 同步销售数据按照渠道类型汇总
type AlibabatmallgeniescpplansummarysaleqtygetAPIRequest struct {
	model.Params
	// 扩展参数
	_requestExtendJson string
}

// NewAlibabatmallgeniescpplansummarysaleqtygetRequest 初始化AlibabatmallgeniescpplansummarysaleqtygetAPIRequest对象
func NewAlibabatmallgeniescpplansummarysaleqtygetRequest() *AlibabatmallgeniescpplansummarysaleqtygetAPIRequest {
	return &AlibabatmallgeniescpplansummarysaleqtygetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabatmallgeniescpplansummarysaleqtygetAPIRequest) GetApiMethodName() string {
	return "alibaba.tmallgenie.scp.plan.summary.sale.qty.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabatmallgeniescpplansummarysaleqtygetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabatmallgeniescpplansummarysaleqtygetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequestExtendJson is RequestExtendJson Setter
// 扩展参数
func (r *AlibabatmallgeniescpplansummarysaleqtygetAPIRequest) SetRequestExtendJson(_requestExtendJson string) error {
	r._requestExtendJson = _requestExtendJson
	r.Set("request_extend_json", _requestExtendJson)
	return nil
}

// GetRequestExtendJson RequestExtendJson Getter
func (r AlibabatmallgeniescpplansummarysaleqtygetAPIRequest) GetRequestExtendJson() string {
	return r._requestExtendJson
}
