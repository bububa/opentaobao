package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunostvpubadmindevicetvidAPIRequest 查询终端信息 API请求
// yunos.tvpubadmin.device.tvid
//
// 通过UUID查询终端信息
type YunostvpubadmindevicetvidAPIRequest struct {
	model.Params
	// 设备的UUID
	_uuid string
}

// NewYunostvpubadmindevicetvidRequest 初始化YunostvpubadmindevicetvidAPIRequest对象
func NewYunostvpubadmindevicetvidRequest() *YunostvpubadmindevicetvidAPIRequest {
	return &YunostvpubadmindevicetvidAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunostvpubadmindevicetvidAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.device.tvid"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunostvpubadmindevicetvidAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunostvpubadmindevicetvidAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUuid is Uuid Setter
// 设备的UUID
func (r *YunostvpubadmindevicetvidAPIRequest) SetUuid(_uuid string) error {
	r._uuid = _uuid
	r.Set("uuid", _uuid)
	return nil
}

// GetUuid Uuid Getter
func (r YunostvpubadmindevicetvidAPIRequest) GetUuid() string {
	return r._uuid
}
