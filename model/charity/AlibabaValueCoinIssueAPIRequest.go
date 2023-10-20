package charity

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabavaluecoinissueAPIRequest 爱豆发放 API请求
// alibaba.value.coin.issue
//
// 爱豆发放
type AlibabavaluecoinissueAPIRequest struct {
	model.Params
	// 爱豆发放参数
	_exCoinIssueParam *ExCoinIssueParam
}

// NewAlibabavaluecoinissueRequest 初始化AlibabavaluecoinissueAPIRequest对象
func NewAlibabavaluecoinissueRequest() *AlibabavaluecoinissueAPIRequest {
	return &AlibabavaluecoinissueAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabavaluecoinissueAPIRequest) GetApiMethodName() string {
	return "alibaba.value.coin.issue"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabavaluecoinissueAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabavaluecoinissueAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetExCoinIssueParam is ExCoinIssueParam Setter
// 爱豆发放参数
func (r *AlibabavaluecoinissueAPIRequest) SetExCoinIssueParam(_exCoinIssueParam *ExCoinIssueParam) error {
	r._exCoinIssueParam = _exCoinIssueParam
	r.Set("ex_coin_issue_param", _exCoinIssueParam)
	return nil
}

// GetExCoinIssueParam ExCoinIssueParam Getter
func (r AlibabavaluecoinissueAPIRequest) GetExCoinIssueParam() *ExCoinIssueParam {
	return r._exCoinIssueParam
}
