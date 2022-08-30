package crm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoCrmExchangeCrowdinstanceAddPrivyAPIRequest 向活动人群实例中增加买家（隐私号版） API请求
// taobao.crm.exchange.crowdinstance.add.privy
//
// 向活动人群实例中增加买家
type TaobaoCrmExchangeCrowdinstanceAddPrivyAPIRequest struct {
	model.Params
	// 操作原因
	_reason string
	// ouid
	_ouid string
	// 人群实例ID
	_crowdInstanceId int64
}

// NewTaobaoCrmExchangeCrowdinstanceAddPrivyRequest 初始化TaobaoCrmExchangeCrowdinstanceAddPrivyAPIRequest对象
func NewTaobaoCrmExchangeCrowdinstanceAddPrivyRequest() *TaobaoCrmExchangeCrowdinstanceAddPrivyAPIRequest {
	return &TaobaoCrmExchangeCrowdinstanceAddPrivyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoCrmExchangeCrowdinstanceAddPrivyAPIRequest) GetApiMethodName() string {
	return "taobao.crm.exchange.crowdinstance.add.privy"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoCrmExchangeCrowdinstanceAddPrivyAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetReason is Reason Setter
// 操作原因
func (r *TaobaoCrmExchangeCrowdinstanceAddPrivyAPIRequest) SetReason(_reason string) error {
	r._reason = _reason
	r.Set("reason", _reason)
	return nil
}

// GetReason Reason Getter
func (r TaobaoCrmExchangeCrowdinstanceAddPrivyAPIRequest) GetReason() string {
	return r._reason
}

// SetOuid is Ouid Setter
// ouid
func (r *TaobaoCrmExchangeCrowdinstanceAddPrivyAPIRequest) SetOuid(_ouid string) error {
	r._ouid = _ouid
	r.Set("ouid", _ouid)
	return nil
}

// GetOuid Ouid Getter
func (r TaobaoCrmExchangeCrowdinstanceAddPrivyAPIRequest) GetOuid() string {
	return r._ouid
}

// SetCrowdInstanceId is CrowdInstanceId Setter
// 人群实例ID
func (r *TaobaoCrmExchangeCrowdinstanceAddPrivyAPIRequest) SetCrowdInstanceId(_crowdInstanceId int64) error {
	r._crowdInstanceId = _crowdInstanceId
	r.Set("crowd_instance_id", _crowdInstanceId)
	return nil
}

// GetCrowdInstanceId CrowdInstanceId Getter
func (r TaobaoCrmExchangeCrowdinstanceAddPrivyAPIRequest) GetCrowdInstanceId() int64 {
	return r._crowdInstanceId
}
