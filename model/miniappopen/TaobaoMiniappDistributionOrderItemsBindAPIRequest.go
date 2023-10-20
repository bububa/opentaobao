package miniappopen

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMiniappDistributionOrderItemsBindAPIRequest 小程序投放-基于投放计划绑定/解绑商品 API请求
// taobao.miniapp.distribution.order.items.bind
//
// 提供给使用了投放插件的服务商，可以调用该API实现帮助商家更新已创建的投放单中的绑定商品信息。
type TaobaoMiniappDistributionOrderItemsBindAPIRequest struct {
	model.Params
	// 商品id列表
	_targetEntityList []string
	// 投放计划标识id
	_distributeId int64
	// true表示新增绑定，false表示解绑
	_addBind bool
}

// NewTaobaoMiniappDistributionOrderItemsBindRequest 初始化TaobaoMiniappDistributionOrderItemsBindAPIRequest对象
func NewTaobaoMiniappDistributionOrderItemsBindRequest() *TaobaoMiniappDistributionOrderItemsBindAPIRequest {
	return &TaobaoMiniappDistributionOrderItemsBindAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoMiniappDistributionOrderItemsBindAPIRequest) Reset() {
	r._targetEntityList = r._targetEntityList[:0]
	r._distributeId = 0
	r._addBind = false
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoMiniappDistributionOrderItemsBindAPIRequest) GetApiMethodName() string {
	return "taobao.miniapp.distribution.order.items.bind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoMiniappDistributionOrderItemsBindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoMiniappDistributionOrderItemsBindAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTargetEntityList is TargetEntityList Setter
// 商品id列表
func (r *TaobaoMiniappDistributionOrderItemsBindAPIRequest) SetTargetEntityList(_targetEntityList []string) error {
	r._targetEntityList = _targetEntityList
	r.Set("target_entity_list", _targetEntityList)
	return nil
}

// GetTargetEntityList TargetEntityList Getter
func (r TaobaoMiniappDistributionOrderItemsBindAPIRequest) GetTargetEntityList() []string {
	return r._targetEntityList
}

// SetDistributeId is DistributeId Setter
// 投放计划标识id
func (r *TaobaoMiniappDistributionOrderItemsBindAPIRequest) SetDistributeId(_distributeId int64) error {
	r._distributeId = _distributeId
	r.Set("distribute_id", _distributeId)
	return nil
}

// GetDistributeId DistributeId Getter
func (r TaobaoMiniappDistributionOrderItemsBindAPIRequest) GetDistributeId() int64 {
	return r._distributeId
}

// SetAddBind is AddBind Setter
// true表示新增绑定，false表示解绑
func (r *TaobaoMiniappDistributionOrderItemsBindAPIRequest) SetAddBind(_addBind bool) error {
	r._addBind = _addBind
	r.Set("add_bind", _addBind)
	return nil
}

// GetAddBind AddBind Getter
func (r TaobaoMiniappDistributionOrderItemsBindAPIRequest) GetAddBind() bool {
	return r._addBind
}

var poolTaobaoMiniappDistributionOrderItemsBindAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoMiniappDistributionOrderItemsBindRequest()
	},
}

// GetTaobaoMiniappDistributionOrderItemsBindRequest 从 sync.Pool 获取 TaobaoMiniappDistributionOrderItemsBindAPIRequest
func GetTaobaoMiniappDistributionOrderItemsBindAPIRequest() *TaobaoMiniappDistributionOrderItemsBindAPIRequest {
	return poolTaobaoMiniappDistributionOrderItemsBindAPIRequest.Get().(*TaobaoMiniappDistributionOrderItemsBindAPIRequest)
}

// ReleaseTaobaoMiniappDistributionOrderItemsBindAPIRequest 将 TaobaoMiniappDistributionOrderItemsBindAPIRequest 放入 sync.Pool
func ReleaseTaobaoMiniappDistributionOrderItemsBindAPIRequest(v *TaobaoMiniappDistributionOrderItemsBindAPIRequest) {
	v.Reset()
	poolTaobaoMiniappDistributionOrderItemsBindAPIRequest.Put(v)
}
