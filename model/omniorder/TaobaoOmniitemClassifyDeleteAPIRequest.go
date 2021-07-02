package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOmniitemClassifyDeleteAPIRequest 删除一个分类 API请求
// taobao.omniitem.classify.delete
//
// 删除一个分类
type TaobaoOmniitemClassifyDeleteAPIRequest struct {
	model.Params
	// 分类ID
	_classifyId int64
	// 操作人信息（暂时不填）
	_operator string
}

// NewTaobaoOmniitemClassifyDeleteRequest 初始化TaobaoOmniitemClassifyDeleteAPIRequest对象
func NewTaobaoOmniitemClassifyDeleteRequest() *TaobaoOmniitemClassifyDeleteAPIRequest {
	return &TaobaoOmniitemClassifyDeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOmniitemClassifyDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.omniitem.classify.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOmniitemClassifyDeleteAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ClassifyId Setter
// 分类ID
func (r *TaobaoOmniitemClassifyDeleteAPIRequest) SetClassifyId(_classifyId int64) error {
	r._classifyId = _classifyId
	r.Set("classify_id", _classifyId)
	return nil
}

// Get ClassifyId Getter
func (r TaobaoOmniitemClassifyDeleteAPIRequest) GetClassifyId() int64 {
	return r._classifyId
}

// Set is Operator Setter
// 操作人信息（暂时不填）
func (r *TaobaoOmniitemClassifyDeleteAPIRequest) SetOperator(_operator string) error {
	r._operator = _operator
	r.Set("operator", _operator)
	return nil
}

// Get Operator Getter
func (r TaobaoOmniitemClassifyDeleteAPIRequest) GetOperator() string {
	return r._operator
}
