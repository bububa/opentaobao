package train

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTrainAgentHandleticketConfirmVtwoAPIRequest 代理商出票中v2--增加鉴权校验 API请求
// taobao.train.agent.handleticket.confirm.vtwo
//
// 代理商出票中
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
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoTrainAgentHandleticketConfirmVtwoAPIRequest) Reset() {
	r._extendParams = ""
	r._mainOrderId = 0
	r._sellerId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTrainAgentHandleticketConfirmVtwoAPIRequest) GetApiMethodName() string {
	return "taobao.train.agent.handleticket.confirm.vtwo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTrainAgentHandleticketConfirmVtwoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTrainAgentHandleticketConfirmVtwoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetExtendParams is ExtendParams Setter
// 扩展参数
func (r *TaobaoTrainAgentHandleticketConfirmVtwoAPIRequest) SetExtendParams(_extendParams string) error {
	r._extendParams = _extendParams
	r.Set("extend_params", _extendParams)
	return nil
}

// GetExtendParams ExtendParams Getter
func (r TaobaoTrainAgentHandleticketConfirmVtwoAPIRequest) GetExtendParams() string {
	return r._extendParams
}

// SetMainOrderId is MainOrderId Setter
// 主站id
func (r *TaobaoTrainAgentHandleticketConfirmVtwoAPIRequest) SetMainOrderId(_mainOrderId int64) error {
	r._mainOrderId = _mainOrderId
	r.Set("main_order_id", _mainOrderId)
	return nil
}

// GetMainOrderId MainOrderId Getter
func (r TaobaoTrainAgentHandleticketConfirmVtwoAPIRequest) GetMainOrderId() int64 {
	return r._mainOrderId
}

// SetSellerId is SellerId Setter
// 代理商id
func (r *TaobaoTrainAgentHandleticketConfirmVtwoAPIRequest) SetSellerId(_sellerId int64) error {
	r._sellerId = _sellerId
	r.Set("seller_id", _sellerId)
	return nil
}

// GetSellerId SellerId Getter
func (r TaobaoTrainAgentHandleticketConfirmVtwoAPIRequest) GetSellerId() int64 {
	return r._sellerId
}

var poolTaobaoTrainAgentHandleticketConfirmVtwoAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoTrainAgentHandleticketConfirmVtwoRequest()
	},
}

// GetTaobaoTrainAgentHandleticketConfirmVtwoRequest 从 sync.Pool 获取 TaobaoTrainAgentHandleticketConfirmVtwoAPIRequest
func GetTaobaoTrainAgentHandleticketConfirmVtwoAPIRequest() *TaobaoTrainAgentHandleticketConfirmVtwoAPIRequest {
	return poolTaobaoTrainAgentHandleticketConfirmVtwoAPIRequest.Get().(*TaobaoTrainAgentHandleticketConfirmVtwoAPIRequest)
}

// ReleaseTaobaoTrainAgentHandleticketConfirmVtwoAPIRequest 将 TaobaoTrainAgentHandleticketConfirmVtwoAPIRequest 放入 sync.Pool
func ReleaseTaobaoTrainAgentHandleticketConfirmVtwoAPIRequest(v *TaobaoTrainAgentHandleticketConfirmVtwoAPIRequest) {
	v.Reset()
	poolTaobaoTrainAgentHandleticketConfirmVtwoAPIRequest.Put(v)
}
