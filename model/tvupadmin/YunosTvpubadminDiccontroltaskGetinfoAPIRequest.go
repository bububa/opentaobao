package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunostvpubadmindiccontroltaskgetinfoAPIRequest 获取停开服任务详情 API请求
// yunos.tvpubadmin.diccontroltask.getinfo
//
// 获取停开服任务详情
type YunostvpubadmindiccontroltaskgetinfoAPIRequest struct {
	model.Params
	// 任务ID
	_id int64
	// 牌照方
	_license int64
}

// NewYunostvpubadmindiccontroltaskgetinfoRequest 初始化YunostvpubadmindiccontroltaskgetinfoAPIRequest对象
func NewYunostvpubadmindiccontroltaskgetinfoRequest() *YunostvpubadmindiccontroltaskgetinfoAPIRequest {
	return &YunostvpubadmindiccontroltaskgetinfoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunostvpubadmindiccontroltaskgetinfoAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.diccontroltask.getinfo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunostvpubadmindiccontroltaskgetinfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunostvpubadmindiccontroltaskgetinfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetId is Id Setter
// 任务ID
func (r *YunostvpubadmindiccontroltaskgetinfoAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r YunostvpubadmindiccontroltaskgetinfoAPIRequest) GetId() int64 {
	return r._id
}

// SetLicense is License Setter
// 牌照方
func (r *YunostvpubadmindiccontroltaskgetinfoAPIRequest) SetLicense(_license int64) error {
	r._license = _license
	r.Set("license", _license)
	return nil
}

// GetLicense License Getter
func (r YunostvpubadmindiccontroltaskgetinfoAPIRequest) GetLicense() int64 {
	return r._license
}
