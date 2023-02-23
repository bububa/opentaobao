package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminDeviceApkinfoAPIRequest 获取停开服apk信息 API请求
// yunos.tvpubadmin.device.apkinfo
//
// 获取停开服apk信息
type YunosTvpubadminDeviceApkinfoAPIRequest struct {
	model.Params
	// apkid
	_id int64
}

// NewYunosTvpubadminDeviceApkinfoRequest 初始化YunosTvpubadminDeviceApkinfoAPIRequest对象
func NewYunosTvpubadminDeviceApkinfoRequest() *YunosTvpubadminDeviceApkinfoAPIRequest {
	return &YunosTvpubadminDeviceApkinfoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosTvpubadminDeviceApkinfoAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.device.apkinfo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosTvpubadminDeviceApkinfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosTvpubadminDeviceApkinfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetId is Id Setter
// apkid
func (r *YunosTvpubadminDeviceApkinfoAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r YunosTvpubadminDeviceApkinfoAPIRequest) GetId() int64 {
	return r._id
}
