package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunostvpubadmindeviceapkinfoAPIRequest 获取停开服apk信息 API请求
// yunos.tvpubadmin.device.apkinfo
//
// 获取停开服apk信息
type YunostvpubadmindeviceapkinfoAPIRequest struct {
	model.Params
	// apkid
	_id int64
}

// NewYunostvpubadmindeviceapkinfoRequest 初始化YunostvpubadmindeviceapkinfoAPIRequest对象
func NewYunostvpubadmindeviceapkinfoRequest() *YunostvpubadmindeviceapkinfoAPIRequest {
	return &YunostvpubadmindeviceapkinfoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunostvpubadmindeviceapkinfoAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.device.apkinfo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunostvpubadmindeviceapkinfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunostvpubadmindeviceapkinfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetId is Id Setter
// apkid
func (r *YunostvpubadmindeviceapkinfoAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r YunostvpubadmindeviceapkinfoAPIRequest) GetId() int64 {
	return r._id
}
