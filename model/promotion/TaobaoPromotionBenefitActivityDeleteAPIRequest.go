package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaopromotionbenefitactivitydeleteAPIRequest 删除关联的活动权益 API请求
// taobao.promotion.benefit.activity.delete
//
// 删除关联的活动权益
type TaobaopromotionbenefitactivitydeleteAPIRequest struct {
	model.Params
	// ISV活动关联权益后获得的关联ID
	_relationId int64
}

// NewTaobaopromotionbenefitactivitydeleteRequest 初始化TaobaopromotionbenefitactivitydeleteAPIRequest对象
func NewTaobaopromotionbenefitactivitydeleteRequest() *TaobaopromotionbenefitactivitydeleteAPIRequest {
	return &TaobaopromotionbenefitactivitydeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaopromotionbenefitactivitydeleteAPIRequest) GetApiMethodName() string {
	return "taobao.promotion.benefit.activity.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaopromotionbenefitactivitydeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaopromotionbenefitactivitydeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRelationId is RelationId Setter
// ISV活动关联权益后获得的关联ID
func (r *TaobaopromotionbenefitactivitydeleteAPIRequest) SetRelationId(_relationId int64) error {
	r._relationId = _relationId
	r.Set("relation_id", _relationId)
	return nil
}

// GetRelationId RelationId Getter
func (r TaobaopromotionbenefitactivitydeleteAPIRequest) GetRelationId() int64 {
	return r._relationId
}
