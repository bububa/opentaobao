package train

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTrainAgentHandleticketConfirmAPIRequest 代理商出票中 API请求
// taobao.train.agent.handleticket.confirm
//
// 代理商出票中
type TaobaoTrainAgentHandleticketConfirmAPIRequest struct {
	model.Params
	// 扩展参数
	_extendParams string
	// 主站id
	_mainOrderId int64
	// 代理商id
	_sellerId int64
}

// NewTaobaoTrainAgentHandleticketConfirmRequest 初始化TaobaoTrainAgentHandleticketConfirmAPIRequest对象
func NewTaobaoTrainAgentHandleticketConfirmRequest() *TaobaoTrainAgentHandleticketConfirmAPIRequest {
	return &TaobaoTrainAgentHandleticketConfirmAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTrainAgentHandleticketConfirmAPIRequest) GetApiMethodName() string {
	return "taobao.train.agent.handleticket.confirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTrainAgentHandleticketConfirmAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ExtendParams Setter
// 扩展参数
func (r *TaobaoTrainAgentHandleticketConfirmAPIRequest) SetExtendParams(_extendParams string) error {
	r._extendParams = _extendParams
	r.Set("extend_params", _extendParams)
	return nil
}

// Get ExtendParams Getter
func (r TaobaoTrainAgentHandleticketConfirmAPIRequest) GetExtendParams() string {
	return r._extendParams
}

// Set is MainOrderId Setter
// 主站id
func (r *TaobaoTrainAgentHandleticketConfirmAPIRequest) SetMainOrderId(_mainOrderId int64) error {
	r._mainOrderId = _mainOrderId
	r.Set("main_order_id", _mainOrderId)
	return nil
}

// Get MainOrderId Getter
func (r TaobaoTrainAgentHandleticketConfirmAPIRequest) GetMainOrderId() int64 {
	return r._mainOrderId
}

// Set is SellerId Setter
// 代理商id
func (r *TaobaoTrainAgentHandleticketConfirmAPIRequest) SetSellerId(_sellerId int64) error {
	r._sellerId = _sellerId
	r.Set("seller_id", _sellerId)
	return nil
}

// Get SellerId Getter
func (r TaobaoTrainAgentHandleticketConfirmAPIRequest) GetSellerId() int64 {
	return r._sellerId
}
