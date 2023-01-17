package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAsrDataservicePromotionruleDeleteAPIRequest 优惠规则删除 API请求
// alibaba.asr.dataservice.promotionrule.delete
//
// 删除优惠规则，例如星巴克删除优惠规则
type AlibabaAsrDataservicePromotionruleDeleteAPIRequest struct {
	model.Params
	// poskey
	_posKey int64
}

// NewAlibabaAsrDataservicePromotionruleDeleteRequest 初始化AlibabaAsrDataservicePromotionruleDeleteAPIRequest对象
func NewAlibabaAsrDataservicePromotionruleDeleteRequest() *AlibabaAsrDataservicePromotionruleDeleteAPIRequest {
	return &AlibabaAsrDataservicePromotionruleDeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAsrDataservicePromotionruleDeleteAPIRequest) GetApiMethodName() string {
	return "alibaba.asr.dataservice.promotionrule.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAsrDataservicePromotionruleDeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAsrDataservicePromotionruleDeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPosKey is PosKey Setter
// poskey
func (r *AlibabaAsrDataservicePromotionruleDeleteAPIRequest) SetPosKey(_posKey int64) error {
	r._posKey = _posKey
	r.Set("pos_key", _posKey)
	return nil
}

// GetPosKey PosKey Getter
func (r AlibabaAsrDataservicePromotionruleDeleteAPIRequest) GetPosKey() int64 {
	return r._posKey
}
