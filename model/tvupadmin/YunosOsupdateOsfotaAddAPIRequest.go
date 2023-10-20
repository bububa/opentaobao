package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunososupdateosfotaaddAPIRequest 添加系统升级任务 API请求
// yunos.osupdate.osfota.add
//
// 添加osupdate系统升级任务
type YunososupdateosfotaaddAPIRequest struct {
	model.Params
	// 系统升级任务json格式
	_osFotaJson string
}

// NewYunososupdateosfotaaddRequest 初始化YunososupdateosfotaaddAPIRequest对象
func NewYunososupdateosfotaaddRequest() *YunososupdateosfotaaddAPIRequest {
	return &YunososupdateosfotaaddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunososupdateosfotaaddAPIRequest) GetApiMethodName() string {
	return "yunos.osupdate.osfota.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunososupdateosfotaaddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunososupdateosfotaaddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOsFotaJson is OsFotaJson Setter
// 系统升级任务json格式
func (r *YunososupdateosfotaaddAPIRequest) SetOsFotaJson(_osFotaJson string) error {
	r._osFotaJson = _osFotaJson
	r.Set("os_fota_json", _osFotaJson)
	return nil
}

// GetOsFotaJson OsFotaJson Getter
func (r YunososupdateosfotaaddAPIRequest) GetOsFotaJson() string {
	return r._osFotaJson
}
