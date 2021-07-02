package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunosOsupdateVersionstatusUpdateAPIRequest 更新应用升级状态 API请求
// yunos.osupdate.versionstatus.update
//
// 更新应用升级状态
type YunosOsupdateVersionstatusUpdateAPIRequest struct {
	model.Params
	// 升级任务ID
	_id int64
	// 状态值
	_status string
}

// NewYunosOsupdateVersionstatusUpdateRequest 初始化YunosOsupdateVersionstatusUpdateAPIRequest对象
func NewYunosOsupdateVersionstatusUpdateRequest() *YunosOsupdateVersionstatusUpdateAPIRequest {
	return &YunosOsupdateVersionstatusUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosOsupdateVersionstatusUpdateAPIRequest) GetApiMethodName() string {
	return "yunos.osupdate.versionstatus.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosOsupdateVersionstatusUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Id Setter
// 升级任务ID
func (r *YunosOsupdateVersionstatusUpdateAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// Get Id Getter
func (r YunosOsupdateVersionstatusUpdateAPIRequest) GetId() int64 {
	return r._id
}

// Set is Status Setter
// 状态值
func (r *YunosOsupdateVersionstatusUpdateAPIRequest) SetStatus(_status string) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// Get Status Getter
func (r YunosOsupdateVersionstatusUpdateAPIRequest) GetStatus() string {
	return r._status
}
