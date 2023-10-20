package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosimbaserchcrowdbatchdeleteAPIRequest 单品搜索人群批量取消溢价 API请求
// taobao.simba.serchcrowd.batch.delete
//
// 删除单品搜索人群溢价功能
type TaobaosimbaserchcrowdbatchdeleteAPIRequest struct {
	model.Params
	// 需要删除的人群id
	_adgroupCrowdIds []string
	// 子帐号nick
	_subNick string
	// 被操作者的淘宝昵称
	_nick string
}

// NewTaobaosimbaserchcrowdbatchdeleteRequest 初始化TaobaosimbaserchcrowdbatchdeleteAPIRequest对象
func NewTaobaosimbaserchcrowdbatchdeleteRequest() *TaobaosimbaserchcrowdbatchdeleteAPIRequest {
	return &TaobaosimbaserchcrowdbatchdeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosimbaserchcrowdbatchdeleteAPIRequest) GetApiMethodName() string {
	return "taobao.simba.serchcrowd.batch.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosimbaserchcrowdbatchdeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosimbaserchcrowdbatchdeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAdgroupCrowdIds is AdgroupCrowdIds Setter
// 需要删除的人群id
func (r *TaobaosimbaserchcrowdbatchdeleteAPIRequest) SetAdgroupCrowdIds(_adgroupCrowdIds []string) error {
	r._adgroupCrowdIds = _adgroupCrowdIds
	r.Set("adgroup_crowd_ids", _adgroupCrowdIds)
	return nil
}

// GetAdgroupCrowdIds AdgroupCrowdIds Getter
func (r TaobaosimbaserchcrowdbatchdeleteAPIRequest) GetAdgroupCrowdIds() []string {
	return r._adgroupCrowdIds
}

// SetSubNick is SubNick Setter
// 子帐号nick
func (r *TaobaosimbaserchcrowdbatchdeleteAPIRequest) SetSubNick(_subNick string) error {
	r._subNick = _subNick
	r.Set("sub_nick", _subNick)
	return nil
}

// GetSubNick SubNick Getter
func (r TaobaosimbaserchcrowdbatchdeleteAPIRequest) GetSubNick() string {
	return r._subNick
}

// SetNick is Nick Setter
// 被操作者的淘宝昵称
func (r *TaobaosimbaserchcrowdbatchdeleteAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaosimbaserchcrowdbatchdeleteAPIRequest) GetNick() string {
	return r._nick
}
