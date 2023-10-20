package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunostvpubadmindiccontroltaskupdateAPIRequest 停开服任务状态变更 API请求
// yunos.tvpubadmin.diccontroltask.update
//
// 停开服任务状态变更
type YunostvpubadmindiccontroltaskupdateAPIRequest struct {
	model.Params
	// 任务ID
	_id int64
	// 任务状态
	_status int64
	// 牌照方
	_license int64
}

// NewYunostvpubadmindiccontroltaskupdateRequest 初始化YunostvpubadmindiccontroltaskupdateAPIRequest对象
func NewYunostvpubadmindiccontroltaskupdateRequest() *YunostvpubadmindiccontroltaskupdateAPIRequest {
	return &YunostvpubadmindiccontroltaskupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunostvpubadmindiccontroltaskupdateAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.diccontroltask.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunostvpubadmindiccontroltaskupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunostvpubadmindiccontroltaskupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetId is Id Setter
// 任务ID
func (r *YunostvpubadmindiccontroltaskupdateAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r YunostvpubadmindiccontroltaskupdateAPIRequest) GetId() int64 {
	return r._id
}

// SetStatus is Status Setter
// 任务状态
func (r *YunostvpubadmindiccontroltaskupdateAPIRequest) SetStatus(_status int64) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r YunostvpubadmindiccontroltaskupdateAPIRequest) GetStatus() int64 {
	return r._status
}

// SetLicense is License Setter
// 牌照方
func (r *YunostvpubadmindiccontroltaskupdateAPIRequest) SetLicense(_license int64) error {
	r._license = _license
	r.Set("license", _license)
	return nil
}

// GetLicense License Getter
func (r YunostvpubadmindiccontroltaskupdateAPIRequest) GetLicense() int64 {
	return r._license
}
