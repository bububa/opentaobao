package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosOsupdateOsfotaAddAPIRequest
添加系统升级任务 API请求
yunos.osupdate.osfota.add

添加osupdate系统升级任务 */
type YunosOsupdateOsfotaAddAPIRequest struct {
	model.Params
	// 系统升级任务json格式
	_osFotaJson string
}

// NewYunosOsupdateOsfotaAddRequest 初始化YunosOsupdateOsfotaAddAPIRequest对象
func NewYunosOsupdateOsfotaAddRequest() *YunosOsupdateOsfotaAddAPIRequest {
	return &YunosOsupdateOsfotaAddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosOsupdateOsfotaAddAPIRequest) GetApiMethodName() string {
	return "yunos.osupdate.osfota.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosOsupdateOsfotaAddAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is OsFotaJson Setter
// 系统升级任务json格式
func (r *YunosOsupdateOsfotaAddAPIRequest) SetOsFotaJson(_osFotaJson string) error {
	r._osFotaJson = _osFotaJson
	r.Set("os_fota_json", _osFotaJson)
	return nil
}

// Get OsFotaJson Getter
func (r YunosOsupdateOsfotaAddAPIRequest) GetOsFotaJson() string {
	return r._osFotaJson
}
