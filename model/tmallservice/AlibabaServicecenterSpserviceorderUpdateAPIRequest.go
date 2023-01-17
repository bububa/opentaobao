package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaServicecenterSpserviceorderUpdateAPIRequest 服务供应链服务单更新 API请求
// alibaba.servicecenter.spserviceorder.update
//
// 服务供应链服务单更新，服务商通过此接口将商品的sn等信息推送到服务单中
type AlibabaServicecenterSpserviceorderUpdateAPIRequest struct {
	model.Params
	// 新设备sn.当action填写addSn、changeSn时必填
	_action string
	// 新设备sn.当action填写addSn、changeSn时必填
	_newSn string
	// 旧设备sn，当action填写changeSn时必填
	_oldSn string
	// 服务有效期开始时间
	_gmtEffect string
	// 服务过期时间
	_gmtExpire string
	// 服务单id
	_spServiceOrderId int64
}

// NewAlibabaServicecenterSpserviceorderUpdateRequest 初始化AlibabaServicecenterSpserviceorderUpdateAPIRequest对象
func NewAlibabaServicecenterSpserviceorderUpdateRequest() *AlibabaServicecenterSpserviceorderUpdateAPIRequest {
	return &AlibabaServicecenterSpserviceorderUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaServicecenterSpserviceorderUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.servicecenter.spserviceorder.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaServicecenterSpserviceorderUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaServicecenterSpserviceorderUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAction is Action Setter
// 新设备sn.当action填写addSn、changeSn时必填
func (r *AlibabaServicecenterSpserviceorderUpdateAPIRequest) SetAction(_action string) error {
	r._action = _action
	r.Set("action", _action)
	return nil
}

// GetAction Action Getter
func (r AlibabaServicecenterSpserviceorderUpdateAPIRequest) GetAction() string {
	return r._action
}

// SetNewSn is NewSn Setter
// 新设备sn.当action填写addSn、changeSn时必填
func (r *AlibabaServicecenterSpserviceorderUpdateAPIRequest) SetNewSn(_newSn string) error {
	r._newSn = _newSn
	r.Set("new_sn", _newSn)
	return nil
}

// GetNewSn NewSn Getter
func (r AlibabaServicecenterSpserviceorderUpdateAPIRequest) GetNewSn() string {
	return r._newSn
}

// SetOldSn is OldSn Setter
// 旧设备sn，当action填写changeSn时必填
func (r *AlibabaServicecenterSpserviceorderUpdateAPIRequest) SetOldSn(_oldSn string) error {
	r._oldSn = _oldSn
	r.Set("old_sn", _oldSn)
	return nil
}

// GetOldSn OldSn Getter
func (r AlibabaServicecenterSpserviceorderUpdateAPIRequest) GetOldSn() string {
	return r._oldSn
}

// SetGmtEffect is GmtEffect Setter
// 服务有效期开始时间
func (r *AlibabaServicecenterSpserviceorderUpdateAPIRequest) SetGmtEffect(_gmtEffect string) error {
	r._gmtEffect = _gmtEffect
	r.Set("gmt_effect", _gmtEffect)
	return nil
}

// GetGmtEffect GmtEffect Getter
func (r AlibabaServicecenterSpserviceorderUpdateAPIRequest) GetGmtEffect() string {
	return r._gmtEffect
}

// SetGmtExpire is GmtExpire Setter
// 服务过期时间
func (r *AlibabaServicecenterSpserviceorderUpdateAPIRequest) SetGmtExpire(_gmtExpire string) error {
	r._gmtExpire = _gmtExpire
	r.Set("gmt_expire", _gmtExpire)
	return nil
}

// GetGmtExpire GmtExpire Getter
func (r AlibabaServicecenterSpserviceorderUpdateAPIRequest) GetGmtExpire() string {
	return r._gmtExpire
}

// SetSpServiceOrderId is SpServiceOrderId Setter
// 服务单id
func (r *AlibabaServicecenterSpserviceorderUpdateAPIRequest) SetSpServiceOrderId(_spServiceOrderId int64) error {
	r._spServiceOrderId = _spServiceOrderId
	r.Set("sp_service_order_id", _spServiceOrderId)
	return nil
}

// GetSpServiceOrderId SpServiceOrderId Getter
func (r AlibabaServicecenterSpserviceorderUpdateAPIRequest) GetSpServiceOrderId() int64 {
	return r._spServiceOrderId
}
