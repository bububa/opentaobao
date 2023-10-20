package tvupadmin

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosOsupdateOsfotaAddAPIRequest 添加系统升级任务 API请求
// yunos.osupdate.osfota.add
//
// 添加osupdate系统升级任务
type YunosOsupdateOsfotaAddAPIRequest struct {
	model.Params
	// 系统升级任务json格式
	_osFotaJson string
}

// NewYunosOsupdateOsfotaAddRequest 初始化YunosOsupdateOsfotaAddAPIRequest对象
func NewYunosOsupdateOsfotaAddRequest() *YunosOsupdateOsfotaAddAPIRequest {
	return &YunosOsupdateOsfotaAddAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *YunosOsupdateOsfotaAddAPIRequest) Reset() {
	r._osFotaJson = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosOsupdateOsfotaAddAPIRequest) GetApiMethodName() string {
	return "yunos.osupdate.osfota.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosOsupdateOsfotaAddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosOsupdateOsfotaAddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOsFotaJson is OsFotaJson Setter
// 系统升级任务json格式
func (r *YunosOsupdateOsfotaAddAPIRequest) SetOsFotaJson(_osFotaJson string) error {
	r._osFotaJson = _osFotaJson
	r.Set("os_fota_json", _osFotaJson)
	return nil
}

// GetOsFotaJson OsFotaJson Getter
func (r YunosOsupdateOsfotaAddAPIRequest) GetOsFotaJson() string {
	return r._osFotaJson
}

var poolYunosOsupdateOsfotaAddAPIRequest = sync.Pool{
	New: func() any {
		return NewYunosOsupdateOsfotaAddRequest()
	},
}

// GetYunosOsupdateOsfotaAddRequest 从 sync.Pool 获取 YunosOsupdateOsfotaAddAPIRequest
func GetYunosOsupdateOsfotaAddAPIRequest() *YunosOsupdateOsfotaAddAPIRequest {
	return poolYunosOsupdateOsfotaAddAPIRequest.Get().(*YunosOsupdateOsfotaAddAPIRequest)
}

// ReleaseYunosOsupdateOsfotaAddAPIRequest 将 YunosOsupdateOsfotaAddAPIRequest 放入 sync.Pool
func ReleaseYunosOsupdateOsfotaAddAPIRequest(v *YunosOsupdateOsfotaAddAPIRequest) {
	v.Reset()
	poolYunosOsupdateOsfotaAddAPIRequest.Put(v)
}
