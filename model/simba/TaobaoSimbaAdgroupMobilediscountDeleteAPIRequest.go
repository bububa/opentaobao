package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosimbaadgroupmobilediscountdeleteAPIRequest 批量删除adgroup的移动溢价 API请求
// taobao.simba.adgroup.mobilediscount.delete
//
// 批量删除adgroup的移动溢价
type TaobaosimbaadgroupmobilediscountdeleteAPIRequest struct {
	model.Params
	// adgroup主键数组（批量最多支持200个）
	_adgroupIds []string
	// 昵称
	_nick string
}

// NewTaobaosimbaadgroupmobilediscountdeleteRequest 初始化TaobaosimbaadgroupmobilediscountdeleteAPIRequest对象
func NewTaobaosimbaadgroupmobilediscountdeleteRequest() *TaobaosimbaadgroupmobilediscountdeleteAPIRequest {
	return &TaobaosimbaadgroupmobilediscountdeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosimbaadgroupmobilediscountdeleteAPIRequest) GetApiMethodName() string {
	return "taobao.simba.adgroup.mobilediscount.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosimbaadgroupmobilediscountdeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosimbaadgroupmobilediscountdeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAdgroupIds is AdgroupIds Setter
// adgroup主键数组（批量最多支持200个）
func (r *TaobaosimbaadgroupmobilediscountdeleteAPIRequest) SetAdgroupIds(_adgroupIds []string) error {
	r._adgroupIds = _adgroupIds
	r.Set("adgroup_ids", _adgroupIds)
	return nil
}

// GetAdgroupIds AdgroupIds Getter
func (r TaobaosimbaadgroupmobilediscountdeleteAPIRequest) GetAdgroupIds() []string {
	return r._adgroupIds
}

// SetNick is Nick Setter
// 昵称
func (r *TaobaosimbaadgroupmobilediscountdeleteAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaosimbaadgroupmobilediscountdeleteAPIRequest) GetNick() string {
	return r._nick
}
