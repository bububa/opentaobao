package tmallservice

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterWorkcardExtrachargeCreateAPIRequest 创建工单额外收费项 API请求
// tmall.servicecenter.workcard.extracharge.create
//
// 创建额外收费项
type TmallServicecenterWorkcardExtrachargeCreateAPIRequest struct {
	model.Params
	// 工单id
	_workcardId int64
	// 额外收费项列表
	_extraChargeItemList *WorkcardExtraChargeCreateTuple
}

// NewTmallServicecenterWorkcardExtrachargeCreateRequest 初始化TmallServicecenterWorkcardExtrachargeCreateAPIRequest对象
func NewTmallServicecenterWorkcardExtrachargeCreateRequest() *TmallServicecenterWorkcardExtrachargeCreateAPIRequest {
	return &TmallServicecenterWorkcardExtrachargeCreateAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallServicecenterWorkcardExtrachargeCreateAPIRequest) Reset() {
	r._workcardId = 0
	r._extraChargeItemList = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServicecenterWorkcardExtrachargeCreateAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.workcard.extracharge.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServicecenterWorkcardExtrachargeCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallServicecenterWorkcardExtrachargeCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWorkcardId is WorkcardId Setter
// 工单id
func (r *TmallServicecenterWorkcardExtrachargeCreateAPIRequest) SetWorkcardId(_workcardId int64) error {
	r._workcardId = _workcardId
	r.Set("workcard_id", _workcardId)
	return nil
}

// GetWorkcardId WorkcardId Getter
func (r TmallServicecenterWorkcardExtrachargeCreateAPIRequest) GetWorkcardId() int64 {
	return r._workcardId
}

// SetExtraChargeItemList is ExtraChargeItemList Setter
// 额外收费项列表
func (r *TmallServicecenterWorkcardExtrachargeCreateAPIRequest) SetExtraChargeItemList(_extraChargeItemList *WorkcardExtraChargeCreateTuple) error {
	r._extraChargeItemList = _extraChargeItemList
	r.Set("extra_charge_item_list", _extraChargeItemList)
	return nil
}

// GetExtraChargeItemList ExtraChargeItemList Getter
func (r TmallServicecenterWorkcardExtrachargeCreateAPIRequest) GetExtraChargeItemList() *WorkcardExtraChargeCreateTuple {
	return r._extraChargeItemList
}

var poolTmallServicecenterWorkcardExtrachargeCreateAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallServicecenterWorkcardExtrachargeCreateRequest()
	},
}

// GetTmallServicecenterWorkcardExtrachargeCreateRequest 从 sync.Pool 获取 TmallServicecenterWorkcardExtrachargeCreateAPIRequest
func GetTmallServicecenterWorkcardExtrachargeCreateAPIRequest() *TmallServicecenterWorkcardExtrachargeCreateAPIRequest {
	return poolTmallServicecenterWorkcardExtrachargeCreateAPIRequest.Get().(*TmallServicecenterWorkcardExtrachargeCreateAPIRequest)
}

// ReleaseTmallServicecenterWorkcardExtrachargeCreateAPIRequest 将 TmallServicecenterWorkcardExtrachargeCreateAPIRequest 放入 sync.Pool
func ReleaseTmallServicecenterWorkcardExtrachargeCreateAPIRequest(v *TmallServicecenterWorkcardExtrachargeCreateAPIRequest) {
	v.Reset()
	poolTmallServicecenterWorkcardExtrachargeCreateAPIRequest.Put(v)
}
