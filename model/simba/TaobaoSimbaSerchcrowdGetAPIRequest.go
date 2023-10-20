package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosimbaserchcrowdgetAPIRequest 根据推广单元id获取搜索溢价人群 API请求
// taobao.simba.serchcrowd.get
//
// 根据推广单元id获取搜索溢价人群
type TaobaosimbaserchcrowdgetAPIRequest struct {
	model.Params
	// 被操作者的淘宝昵称
	_nick string
	// 推广单元id
	_adgroupId int64
}

// NewTaobaosimbaserchcrowdgetRequest 初始化TaobaosimbaserchcrowdgetAPIRequest对象
func NewTaobaosimbaserchcrowdgetRequest() *TaobaosimbaserchcrowdgetAPIRequest {
	return &TaobaosimbaserchcrowdgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosimbaserchcrowdgetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.serchcrowd.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosimbaserchcrowdgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosimbaserchcrowdgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 被操作者的淘宝昵称
func (r *TaobaosimbaserchcrowdgetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaosimbaserchcrowdgetAPIRequest) GetNick() string {
	return r._nick
}

// SetAdgroupId is AdgroupId Setter
// 推广单元id
func (r *TaobaosimbaserchcrowdgetAPIRequest) SetAdgroupId(_adgroupId int64) error {
	r._adgroupId = _adgroupId
	r.Set("adgroup_id", _adgroupId)
	return nil
}

// GetAdgroupId AdgroupId Getter
func (r TaobaosimbaserchcrowdgetAPIRequest) GetAdgroupId() int64 {
	return r._adgroupId
}
