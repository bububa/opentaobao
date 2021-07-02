package tvupadmin

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosTvpubadminDiccontroltaskGetinfoAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.diccontroltask.getinfo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosTvpubadminDiccontroltaskGetinfoAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Id Setter
// 任务ID
func (r *YunosTvpubadminDiccontroltaskGetinfoAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// Get Id Getter
func (r YunosTvpubadminDiccontroltaskGetinfoAPIRequest) GetId() int64 {
	return r._id
}

// Set is License Setter
// 牌照方
func (r *YunosTvpubadminDiccontroltaskGetinfoAPIRequest) SetLicense(_license int64) error {
	r._license = _license
	r.Set("license", _license)
	return nil
}

// Get License Getter
func (r YunosTvpubadminDiccontroltaskGetinfoAPIRequest) GetLicense() int64 {
	return r._license
}
