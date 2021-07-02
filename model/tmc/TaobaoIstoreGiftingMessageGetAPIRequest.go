package tmc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoIstoreGiftingMessageGetAPIRequest gifting消息获取 API请求
// taobao.istore.gifting.message.get
//
// 该api通过参数查询对应的gifting消息
type TaobaoIstoreGiftingMessageGetAPIRequest struct {
	model.Params
	// 消息查询条件
	_giftMessageBizCondition *GiftMessageBizCondition
}

// NewTaobaoIstoreGiftingMessageGetRequest 初始化TaobaoIstoreGiftingMessageGetAPIRequest对象
func NewTaobaoIstoreGiftingMessageGetRequest() *TaobaoIstoreGiftingMessageGetAPIRequest {
	return &TaobaoIstoreGiftingMessageGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoIstoreGiftingMessageGetAPIRequest) GetApiMethodName() string {
	return "taobao.istore.gifting.message.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoIstoreGiftingMessageGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is GiftMessageBizCondition Setter
// 消息查询条件
func (r *TaobaoIstoreGiftingMessageGetAPIRequest) SetGiftMessageBizCondition(_giftMessageBizCondition *GiftMessageBizCondition) error {
	r._giftMessageBizCondition = _giftMessageBizCondition
	r.Set("gift_message_biz_condition", _giftMessageBizCondition)
	return nil
}

// Get GiftMessageBizCondition Getter
func (r TaobaoIstoreGiftingMessageGetAPIRequest) GetGiftMessageBizCondition() *GiftMessageBizCondition {
	return r._giftMessageBizCondition
}
