package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallservicesettlementbillinfoqueryAPIRequest 查询工单结算信息 API请求
// tmall.service.settlement.billinfo.query
//
// 提供给服务商查询工单结算信息，包含结算的分成金额以及结算的收款明细，平台抽佣比例。用于服务商进行对账
type TmallservicesettlementbillinfoqueryAPIRequest struct {
	model.Params
	// 已经结算的工单ID
	_workcardId int64
}

// NewTmallservicesettlementbillinfoqueryRequest 初始化TmallservicesettlementbillinfoqueryAPIRequest对象
func NewTmallservicesettlementbillinfoqueryRequest() *TmallservicesettlementbillinfoqueryAPIRequest {
	return &TmallservicesettlementbillinfoqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallservicesettlementbillinfoqueryAPIRequest) GetApiMethodName() string {
	return "tmall.service.settlement.billinfo.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallservicesettlementbillinfoqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallservicesettlementbillinfoqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWorkcardId is WorkcardId Setter
// 已经结算的工单ID
func (r *TmallservicesettlementbillinfoqueryAPIRequest) SetWorkcardId(_workcardId int64) error {
	r._workcardId = _workcardId
	r.Set("workcard_id", _workcardId)
	return nil
}

// GetWorkcardId WorkcardId Getter
func (r TmallservicesettlementbillinfoqueryAPIRequest) GetWorkcardId() int64 {
	return r._workcardId
}
