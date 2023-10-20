package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallservicecenterworkcardexpressordercreateAPIRequest 物流订单创建API API请求
// tmall.servicecenter.workcard.expressorder.create
//
// 天猫服务寄送类业务，服务商履约完成后寄回操作时，提供的物流寄件单创建API
type TmallservicecenterworkcardexpressordercreateAPIRequest struct {
	model.Params
	// 物流单关联的工单List
	_workcardIdList []int64
	// 真实履约服务商Nick（非ERP系统不要填写）
	_realTpNick string
	// erp外部物流订单号
	_externalLogisticsOrderId string
}

// NewTmallservicecenterworkcardexpressordercreateRequest 初始化TmallservicecenterworkcardexpressordercreateAPIRequest对象
func NewTmallservicecenterworkcardexpressordercreateRequest() *TmallservicecenterworkcardexpressordercreateAPIRequest {
	return &TmallservicecenterworkcardexpressordercreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallservicecenterworkcardexpressordercreateAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.workcard.expressorder.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallservicecenterworkcardexpressordercreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallservicecenterworkcardexpressordercreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWorkcardIdList is WorkcardIdList Setter
// 物流单关联的工单List
func (r *TmallservicecenterworkcardexpressordercreateAPIRequest) SetWorkcardIdList(_workcardIdList []int64) error {
	r._workcardIdList = _workcardIdList
	r.Set("workcard_id_list", _workcardIdList)
	return nil
}

// GetWorkcardIdList WorkcardIdList Getter
func (r TmallservicecenterworkcardexpressordercreateAPIRequest) GetWorkcardIdList() []int64 {
	return r._workcardIdList
}

// SetRealTpNick is RealTpNick Setter
// 真实履约服务商Nick（非ERP系统不要填写）
func (r *TmallservicecenterworkcardexpressordercreateAPIRequest) SetRealTpNick(_realTpNick string) error {
	r._realTpNick = _realTpNick
	r.Set("real_tp_nick", _realTpNick)
	return nil
}

// GetRealTpNick RealTpNick Getter
func (r TmallservicecenterworkcardexpressordercreateAPIRequest) GetRealTpNick() string {
	return r._realTpNick
}

// SetExternalLogisticsOrderId is ExternalLogisticsOrderId Setter
// erp外部物流订单号
func (r *TmallservicecenterworkcardexpressordercreateAPIRequest) SetExternalLogisticsOrderId(_externalLogisticsOrderId string) error {
	r._externalLogisticsOrderId = _externalLogisticsOrderId
	r.Set("external_logistics_order_id", _externalLogisticsOrderId)
	return nil
}

// GetExternalLogisticsOrderId ExternalLogisticsOrderId Getter
func (r TmallservicecenterworkcardexpressordercreateAPIRequest) GetExternalLogisticsOrderId() string {
	return r._externalLogisticsOrderId
}
