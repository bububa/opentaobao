package homeai

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoHomeaiAlgPredictAPIRequest 硬装预测接口 API请求
// taobao.homeai.alg.predict
//
// 居然之家硬装预测服务
type TaobaoHomeaiAlgPredictAPIRequest struct {
	model.Params
	// 来源房间json
	_fromCase string
	// 我的家房间json
	_toCase string
}

// NewTaobaoHomeaiAlgPredictRequest 初始化TaobaoHomeaiAlgPredictAPIRequest对象
func NewTaobaoHomeaiAlgPredictRequest() *TaobaoHomeaiAlgPredictAPIRequest {
	return &TaobaoHomeaiAlgPredictAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoHomeaiAlgPredictAPIRequest) Reset() {
	r._fromCase = ""
	r._toCase = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoHomeaiAlgPredictAPIRequest) GetApiMethodName() string {
	return "taobao.homeai.alg.predict"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoHomeaiAlgPredictAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoHomeaiAlgPredictAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFromCase is FromCase Setter
// 来源房间json
func (r *TaobaoHomeaiAlgPredictAPIRequest) SetFromCase(_fromCase string) error {
	r._fromCase = _fromCase
	r.Set("from_case", _fromCase)
	return nil
}

// GetFromCase FromCase Getter
func (r TaobaoHomeaiAlgPredictAPIRequest) GetFromCase() string {
	return r._fromCase
}

// SetToCase is ToCase Setter
// 我的家房间json
func (r *TaobaoHomeaiAlgPredictAPIRequest) SetToCase(_toCase string) error {
	r._toCase = _toCase
	r.Set("to_case", _toCase)
	return nil
}

// GetToCase ToCase Getter
func (r TaobaoHomeaiAlgPredictAPIRequest) GetToCase() string {
	return r._toCase
}

var poolTaobaoHomeaiAlgPredictAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoHomeaiAlgPredictRequest()
	},
}

// GetTaobaoHomeaiAlgPredictRequest 从 sync.Pool 获取 TaobaoHomeaiAlgPredictAPIRequest
func GetTaobaoHomeaiAlgPredictAPIRequest() *TaobaoHomeaiAlgPredictAPIRequest {
	return poolTaobaoHomeaiAlgPredictAPIRequest.Get().(*TaobaoHomeaiAlgPredictAPIRequest)
}

// ReleaseTaobaoHomeaiAlgPredictAPIRequest 将 TaobaoHomeaiAlgPredictAPIRequest 放入 sync.Pool
func ReleaseTaobaoHomeaiAlgPredictAPIRequest(v *TaobaoHomeaiAlgPredictAPIRequest) {
	v.Reset()
	poolTaobaoHomeaiAlgPredictAPIRequest.Put(v)
}
