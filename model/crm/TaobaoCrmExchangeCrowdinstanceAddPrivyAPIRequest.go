package crm

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoCrmExchangeCrowdinstanceAddPrivyAPIRequest 向活动人群实例中增加买家（隐私号版） API请求
// taobao.crm.exchange.crowdinstance.add.privy
//
// 向活动人群实例中增加买家
type TaobaoCrmExchangeCrowdinstanceAddPrivyAPIRequest struct {
	model.Params
	// ouid
	_ouid string
	// 操作原因
	_reason string
	// 人群实例ID
	_crowdInstanceId int64
}

// NewTaobaoCrmExchangeCrowdinstanceAddPrivyRequest 初始化TaobaoCrmExchangeCrowdinstanceAddPrivyAPIRequest对象
func NewTaobaoCrmExchangeCrowdinstanceAddPrivyRequest() *TaobaoCrmExchangeCrowdinstanceAddPrivyAPIRequest {
	return &TaobaoCrmExchangeCrowdinstanceAddPrivyAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoCrmExchangeCrowdinstanceAddPrivyAPIRequest) Reset() {
	r._ouid = ""
	r._reason = ""
	r._crowdInstanceId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoCrmExchangeCrowdinstanceAddPrivyAPIRequest) GetApiMethodName() string {
	return "taobao.crm.exchange.crowdinstance.add.privy"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoCrmExchangeCrowdinstanceAddPrivyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoCrmExchangeCrowdinstanceAddPrivyAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolTaobaoCrmExchangeCrowdinstanceAddPrivyAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoCrmExchangeCrowdinstanceAddPrivyRequest()
	},
}

// GetTaobaoCrmExchangeCrowdinstanceAddPrivyRequest 从 sync.Pool 获取 TaobaoCrmExchangeCrowdinstanceAddPrivyAPIRequest
func GetTaobaoCrmExchangeCrowdinstanceAddPrivyAPIRequest() *TaobaoCrmExchangeCrowdinstanceAddPrivyAPIRequest {
	return poolTaobaoCrmExchangeCrowdinstanceAddPrivyAPIRequest.Get().(*TaobaoCrmExchangeCrowdinstanceAddPrivyAPIRequest)
}

// ReleaseTaobaoCrmExchangeCrowdinstanceAddPrivyAPIRequest 将 TaobaoCrmExchangeCrowdinstanceAddPrivyAPIRequest 放入 sync.Pool
func ReleaseTaobaoCrmExchangeCrowdinstanceAddPrivyAPIRequest(v *TaobaoCrmExchangeCrowdinstanceAddPrivyAPIRequest) {
	v.Reset()
	poolTaobaoCrmExchangeCrowdinstanceAddPrivyAPIRequest.Put(v)
}
