package homeai

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaohomeaialgpredictAPIRequest 硬装预测接口 API请求
// taobao.homeai.alg.predict
//
// 居然之家硬装预测服务
type TaobaohomeaialgpredictAPIRequest struct {
	model.Params
	// 来源房间json
	_fromCase string
	// 我的家房间json
	_toCase string
}

// NewTaobaohomeaialgpredictRequest 初始化TaobaohomeaialgpredictAPIRequest对象
func NewTaobaohomeaialgpredictRequest() *TaobaohomeaialgpredictAPIRequest {
	return &TaobaohomeaialgpredictAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaohomeaialgpredictAPIRequest) GetApiMethodName() string {
	return "taobao.homeai.alg.predict"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaohomeaialgpredictAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaohomeaialgpredictAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFromCase is FromCase Setter
// 来源房间json
func (r *TaobaohomeaialgpredictAPIRequest) SetFromCase(_fromCase string) error {
	r._fromCase = _fromCase
	r.Set("from_case", _fromCase)
	return nil
}

// GetFromCase FromCase Getter
func (r TaobaohomeaialgpredictAPIRequest) GetFromCase() string {
	return r._fromCase
}

// SetToCase is ToCase Setter
// 我的家房间json
func (r *TaobaohomeaialgpredictAPIRequest) SetToCase(_toCase string) error {
	r._toCase = _toCase
	r.Set("to_case", _toCase)
	return nil
}

// GetToCase ToCase Getter
func (r TaobaohomeaialgpredictAPIRequest) GetToCase() string {
	return r._toCase
}
