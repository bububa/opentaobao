package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosimbasalestaradgroupdeleteAPIRequest (新)销量明星删除推广单元接口 API请求
// taobao.simba.salestar.adgroup.delete
//
// 删除一个推广组
type TaobaosimbasalestaradgroupdeleteAPIRequest struct {
	model.Params
	// 主人昵称
	_nick string
	// 推广组Id
	_adgroupId int64
}

// NewTaobaosimbasalestaradgroupdeleteRequest 初始化TaobaosimbasalestaradgroupdeleteAPIRequest对象
func NewTaobaosimbasalestaradgroupdeleteRequest() *TaobaosimbasalestaradgroupdeleteAPIRequest {
	return &TaobaosimbasalestaradgroupdeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosimbasalestaradgroupdeleteAPIRequest) GetApiMethodName() string {
	return "taobao.simba.salestar.adgroup.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosimbasalestaradgroupdeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosimbasalestaradgroupdeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 主人昵称
func (r *TaobaosimbasalestaradgroupdeleteAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaosimbasalestaradgroupdeleteAPIRequest) GetNick() string {
	return r._nick
}

// SetAdgroupId is AdgroupId Setter
// 推广组Id
func (r *TaobaosimbasalestaradgroupdeleteAPIRequest) SetAdgroupId(_adgroupId int64) error {
	r._adgroupId = _adgroupId
	r.Set("adgroup_id", _adgroupId)
	return nil
}

// GetAdgroupId AdgroupId Getter
func (r TaobaosimbasalestaradgroupdeleteAPIRequest) GetAdgroupId() int64 {
	return r._adgroupId
}
