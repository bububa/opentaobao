package crm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaocrmexchangecrowdinstanceaddprivyAPIRequest 向活动人群实例中增加买家（隐私号版） API请求
// taobao.crm.exchange.crowdinstance.add.privy
//
// 向活动人群实例中增加买家
type TaobaocrmexchangecrowdinstanceaddprivyAPIRequest struct {
	model.Params
	// ouid
	_ouid string
	// 操作原因
	_reason string
	// 人群实例ID
	_crowdInstanceId int64
}

// NewTaobaocrmexchangecrowdinstanceaddprivyRequest 初始化TaobaocrmexchangecrowdinstanceaddprivyAPIRequest对象
func NewTaobaocrmexchangecrowdinstanceaddprivyRequest() *TaobaocrmexchangecrowdinstanceaddprivyAPIRequest {
	return &TaobaocrmexchangecrowdinstanceaddprivyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaocrmexchangecrowdinstanceaddprivyAPIRequest) GetApiMethodName() string {
	return "taobao.crm.exchange.crowdinstance.add.privy"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaocrmexchangecrowdinstanceaddprivyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaocrmexchangecrowdinstanceaddprivyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOuid is Ouid Setter
// ouid
func (r *TaobaocrmexchangecrowdinstanceaddprivyAPIRequest) SetOuid(_ouid string) error {
	r._ouid = _ouid
	r.Set("ouid", _ouid)
	return nil
}

// GetOuid Ouid Getter
func (r TaobaocrmexchangecrowdinstanceaddprivyAPIRequest) GetOuid() string {
	return r._ouid
}

// SetReason is Reason Setter
// 操作原因
func (r *TaobaocrmexchangecrowdinstanceaddprivyAPIRequest) SetReason(_reason string) error {
	r._reason = _reason
	r.Set("reason", _reason)
	return nil
}

// GetReason Reason Getter
func (r TaobaocrmexchangecrowdinstanceaddprivyAPIRequest) GetReason() string {
	return r._reason
}

// SetCrowdInstanceId is CrowdInstanceId Setter
// 人群实例ID
func (r *TaobaocrmexchangecrowdinstanceaddprivyAPIRequest) SetCrowdInstanceId(_crowdInstanceId int64) error {
	r._crowdInstanceId = _crowdInstanceId
	r.Set("crowd_instance_id", _crowdInstanceId)
	return nil
}

// GetCrowdInstanceId CrowdInstanceId Getter
func (r TaobaocrmexchangecrowdinstanceaddprivyAPIRequest) GetCrowdInstanceId() int64 {
	return r._crowdInstanceId
}
