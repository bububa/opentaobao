package tvupadmin

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminDiccontroltaskGetinfoAPIRequest 获取停开服任务详情 API请求
// yunos.tvpubadmin.diccontroltask.getinfo
//
// 获取停开服任务详情
type YunosTvpubadminDiccontroltaskGetinfoAPIRequest struct {
	model.Params
	// 任务ID
	_id int64
	// 牌照方
	_license int64
}

// NewYunosTvpubadminDiccontroltaskGetinfoRequest 初始化YunosTvpubadminDiccontroltaskGetinfoAPIRequest对象
func NewYunosTvpubadminDiccontroltaskGetinfoRequest() *YunosTvpubadminDiccontroltaskGetinfoAPIRequest {
	return &YunosTvpubadminDiccontroltaskGetinfoAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *YunosTvpubadminDiccontroltaskGetinfoAPIRequest) Reset() {
	r._id = 0
	r._license = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosTvpubadminDiccontroltaskGetinfoAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.diccontroltask.getinfo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosTvpubadminDiccontroltaskGetinfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosTvpubadminDiccontroltaskGetinfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetId is Id Setter
// 任务ID
func (r *YunosTvpubadminDiccontroltaskGetinfoAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r YunosTvpubadminDiccontroltaskGetinfoAPIRequest) GetId() int64 {
	return r._id
}

// SetLicense is License Setter
// 牌照方
func (r *YunosTvpubadminDiccontroltaskGetinfoAPIRequest) SetLicense(_license int64) error {
	r._license = _license
	r.Set("license", _license)
	return nil
}

// GetLicense License Getter
func (r YunosTvpubadminDiccontroltaskGetinfoAPIRequest) GetLicense() int64 {
	return r._license
}

var poolYunosTvpubadminDiccontroltaskGetinfoAPIRequest = sync.Pool{
	New: func() any {
		return NewYunosTvpubadminDiccontroltaskGetinfoRequest()
	},
}

// GetYunosTvpubadminDiccontroltaskGetinfoRequest 从 sync.Pool 获取 YunosTvpubadminDiccontroltaskGetinfoAPIRequest
func GetYunosTvpubadminDiccontroltaskGetinfoAPIRequest() *YunosTvpubadminDiccontroltaskGetinfoAPIRequest {
	return poolYunosTvpubadminDiccontroltaskGetinfoAPIRequest.Get().(*YunosTvpubadminDiccontroltaskGetinfoAPIRequest)
}

// ReleaseYunosTvpubadminDiccontroltaskGetinfoAPIRequest 将 YunosTvpubadminDiccontroltaskGetinfoAPIRequest 放入 sync.Pool
func ReleaseYunosTvpubadminDiccontroltaskGetinfoAPIRequest(v *YunosTvpubadminDiccontroltaskGetinfoAPIRequest) {
	v.Reset()
	poolYunosTvpubadminDiccontroltaskGetinfoAPIRequest.Put(v)
}
