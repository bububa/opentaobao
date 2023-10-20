package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaopromotionbenefitactivitydetailgetAPIRequest 活动关联的权益详情获取 API请求
// taobao.promotion.benefit.activity.detail.get
//
// 活动关联的权益详情获取
type TaobaopromotionbenefitactivitydetailgetAPIRequest struct {
	model.Params
	// 查询活动关联权益详情的请求
	_queryRequest *ActivityRelationDetailRequest
}

// NewTaobaopromotionbenefitactivitydetailgetRequest 初始化TaobaopromotionbenefitactivitydetailgetAPIRequest对象
func NewTaobaopromotionbenefitactivitydetailgetRequest() *TaobaopromotionbenefitactivitydetailgetAPIRequest {
	return &TaobaopromotionbenefitactivitydetailgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaopromotionbenefitactivitydetailgetAPIRequest) GetApiMethodName() string {
	return "taobao.promotion.benefit.activity.detail.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaopromotionbenefitactivitydetailgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaopromotionbenefitactivitydetailgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQueryRequest is QueryRequest Setter
// 查询活动关联权益详情的请求
func (r *TaobaopromotionbenefitactivitydetailgetAPIRequest) SetQueryRequest(_queryRequest *ActivityRelationDetailRequest) error {
	r._queryRequest = _queryRequest
	r.Set("query_request", _queryRequest)
	return nil
}

// GetQueryRequest QueryRequest Getter
func (r TaobaopromotionbenefitactivitydetailgetAPIRequest) GetQueryRequest() *ActivityRelationDetailRequest {
	return r._queryRequest
}
