package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosimbaadgroupdeleteAPIRequest 删除一个推广组 API请求
// taobao.simba.adgroup.delete
//
// 删除一个推广组
type TaobaosimbaadgroupdeleteAPIRequest struct {
	model.Params
	// 主人昵称
	_nick string
	// 推广组Id
	_adgroupId int64
}

// NewTaobaosimbaadgroupdeleteRequest 初始化TaobaosimbaadgroupdeleteAPIRequest对象
func NewTaobaosimbaadgroupdeleteRequest() *TaobaosimbaadgroupdeleteAPIRequest {
	return &TaobaosimbaadgroupdeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosimbaadgroupdeleteAPIRequest) GetApiMethodName() string {
	return "taobao.simba.adgroup.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosimbaadgroupdeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosimbaadgroupdeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 主人昵称
func (r *TaobaosimbaadgroupdeleteAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaosimbaadgroupdeleteAPIRequest) GetNick() string {
	return r._nick
}

// SetAdgroupId is AdgroupId Setter
// 推广组Id
func (r *TaobaosimbaadgroupdeleteAPIRequest) SetAdgroupId(_adgroupId int64) error {
	r._adgroupId = _adgroupId
	r.Set("adgroup_id", _adgroupId)
	return nil
}

// GetAdgroupId AdgroupId Getter
func (r TaobaosimbaadgroupdeleteAPIRequest) GetAdgroupId() int64 {
	return r._adgroupId
}
