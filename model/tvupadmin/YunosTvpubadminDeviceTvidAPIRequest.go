package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminDeviceTvidAPIRequest 查询终端信息 API请求
// yunos.tvpubadmin.device.tvid
//
// 通过UUID查询终端信息
type YunosTvpubadminDeviceTvidAPIRequest struct {
	model.Params
	// 设备的UUID
	_uuid string
}

// NewYunosTvpubadminDeviceTvidRequest 初始化YunosTvpubadminDeviceTvidAPIRequest对象
func NewYunosTvpubadminDeviceTvidRequest() *YunosTvpubadminDeviceTvidAPIRequest {
	return &YunosTvpubadminDeviceTvidAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosTvpubadminDeviceTvidAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.device.tvid"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosTvpubadminDeviceTvidAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetUuid is Uuid Setter
// 设备的UUID
func (r *YunosTvpubadminDeviceTvidAPIRequest) SetUuid(_uuid string) error {
	r._uuid = _uuid
	r.Set("uuid", _uuid)
	return nil
}

// GetUuid Uuid Getter
func (r YunosTvpubadminDeviceTvidAPIRequest) GetUuid() string {
	return r._uuid
}
