package train

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTrainAgentHandleticketConfirmVtwoAPIRequest
代理商出票中v2--增加鉴权校验 API请求
taobao.train.agent.handleticket.confirm.vtwo

代理商出票中 */
type TaobaoTrainAgentHandleticketConfirmVtwoAPIRequest struct {
	model.Params
	// 扩展参数
	_extendParams string
	// 主站id
	_mainOrderId int64
	// 代理商id
	_sellerId int64
}

// NewTaobaoTrainAgentHandleticketConfirmVtwoRequest 初始化TaobaoTrainAgentHandleticketConfirmVtwoAPIRequest对象
func NewTaobaoTrainAgentHandleticketConfirmVtwoRequest() *TaobaoTrainAgentHandleticketConfirmVtwoAPIRequest {
	return &TaobaoTrainAgentHandleticketConfirmVtwoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTrainAgentHandleticketConfirmVtwoAPIRequest) GetApiMethodName() string {
	return "taobao.train.agent.handleticket.confirm.vtwo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTrainAgentHandleticketConfirmVtwoAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ExtendParams Setter
// 扩展参数
func (r *TaobaoTrainAgentHandleticketConfirmVtwoAPIRequest) SetExtendParams(_extendParams string) error {
	r._extendParams = _extendParams
	r.Set("extend_params", _extendParams)
	return nil
}

// Get ExtendParams Getter
func (r TaobaoTrainAgentHandleticketConfirmVtwoAPIRequest) GetExtendParams() string {
	return r._extendParams
}

// Set is MainOrderId Setter
// 主站id
func (r *TaobaoTrainAgentHandleticketConfirmVtwoAPIRequest) SetMainOrderId(_mainOrderId int64) error {
	r._mainOrderId = _mainOrderId
	r.Set("main_order_id", _mainOrderId)
	return nil
}

// Get MainOrderId Getter
func (r TaobaoTrainAgentHandleticketConfirmVtwoAPIRequest) GetMainOrderId() int64 {
	return r._mainOrderId
}

// Set is SellerId Setter
// 代理商id
func (r *TaobaoTrainAgentHandleticketConfirmVtwoAPIRequest) SetSellerId(_sellerId int64) error {
	r._sellerId = _sellerId
	r.Set("seller_id", _sellerId)
	return nil
}

// Get SellerId Getter
func (r TaobaoTrainAgentHandleticketConfirmVtwoAPIRequest) GetSellerId() int64 {
	return r._sellerId
}
