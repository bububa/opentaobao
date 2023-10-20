package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunososupdateversionstatusupdateAPIRequest 更新应用升级状态 API请求
// yunos.osupdate.versionstatus.update
//
// 更新应用升级状态
type YunososupdateversionstatusupdateAPIRequest struct {
	model.Params
	// 状态值
	_status string
	// 升级任务ID
	_id int64
}

// NewYunososupdateversionstatusupdateRequest 初始化YunososupdateversionstatusupdateAPIRequest对象
func NewYunososupdateversionstatusupdateRequest() *YunososupdateversionstatusupdateAPIRequest {
	return &YunososupdateversionstatusupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunososupdateversionstatusupdateAPIRequest) GetApiMethodName() string {
	return "yunos.osupdate.versionstatus.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunososupdateversionstatusupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunososupdateversionstatusupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStatus is Status Setter
// 状态值
func (r *YunososupdateversionstatusupdateAPIRequest) SetStatus(_status string) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r YunososupdateversionstatusupdateAPIRequest) GetStatus() string {
	return r._status
}

// SetId is Id Setter
// 升级任务ID
func (r *YunososupdateversionstatusupdateAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r YunososupdateversionstatusupdateAPIRequest) GetId() int64 {
	return r._id
}
