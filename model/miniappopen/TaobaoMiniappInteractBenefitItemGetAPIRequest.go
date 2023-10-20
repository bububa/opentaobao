package miniappopen

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMiniappInteractBenefitItemGetAPIRequest 读取实物权益奖池对应绑定的专属下单商品 API请求
// taobao.miniapp.interact.benefit.item.get
//
// 读取实物权益奖池对应绑定的专属下单商品
type TaobaoMiniappInteractBenefitItemGetAPIRequest struct {
	model.Params
	// 查询条件
	_miniAppSellerStrategyBenefitItemBindRequest *SellerStrategyBenefitItemBindOpenRequest
}

// NewTaobaoMiniappInteractBenefitItemGetRequest 初始化TaobaoMiniappInteractBenefitItemGetAPIRequest对象
func NewTaobaoMiniappInteractBenefitItemGetRequest() *TaobaoMiniappInteractBenefitItemGetAPIRequest {
	return &TaobaoMiniappInteractBenefitItemGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoMiniappInteractBenefitItemGetAPIRequest) Reset() {
	r._miniAppSellerStrategyBenefitItemBindRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoMiniappInteractBenefitItemGetAPIRequest) GetApiMethodName() string {
	return "taobao.miniapp.interact.benefit.item.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoMiniappInteractBenefitItemGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoMiniappInteractBenefitItemGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMiniAppSellerStrategyBenefitItemBindRequest is MiniAppSellerStrategyBenefitItemBindRequest Setter
// 查询条件
func (r *TaobaoMiniappInteractBenefitItemGetAPIRequest) SetMiniAppSellerStrategyBenefitItemBindRequest(_miniAppSellerStrategyBenefitItemBindRequest *SellerStrategyBenefitItemBindOpenRequest) error {
	r._miniAppSellerStrategyBenefitItemBindRequest = _miniAppSellerStrategyBenefitItemBindRequest
	r.Set("mini_app_seller_strategy_benefit_item_bind_request", _miniAppSellerStrategyBenefitItemBindRequest)
	return nil
}

// GetMiniAppSellerStrategyBenefitItemBindRequest MiniAppSellerStrategyBenefitItemBindRequest Getter
func (r TaobaoMiniappInteractBenefitItemGetAPIRequest) GetMiniAppSellerStrategyBenefitItemBindRequest() *SellerStrategyBenefitItemBindOpenRequest {
	return r._miniAppSellerStrategyBenefitItemBindRequest
}

var poolTaobaoMiniappInteractBenefitItemGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoMiniappInteractBenefitItemGetRequest()
	},
}

// GetTaobaoMiniappInteractBenefitItemGetRequest 从 sync.Pool 获取 TaobaoMiniappInteractBenefitItemGetAPIRequest
func GetTaobaoMiniappInteractBenefitItemGetAPIRequest() *TaobaoMiniappInteractBenefitItemGetAPIRequest {
	return poolTaobaoMiniappInteractBenefitItemGetAPIRequest.Get().(*TaobaoMiniappInteractBenefitItemGetAPIRequest)
}

// ReleaseTaobaoMiniappInteractBenefitItemGetAPIRequest 将 TaobaoMiniappInteractBenefitItemGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoMiniappInteractBenefitItemGetAPIRequest(v *TaobaoMiniappInteractBenefitItemGetAPIRequest) {
	v.Reset()
	poolTaobaoMiniappInteractBenefitItemGetAPIRequest.Put(v)
}
