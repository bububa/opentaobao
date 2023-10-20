package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosimbaserchcrowdstatebatchupdateAPIRequest 单品搜索人群修改状态 API请求
// taobao.simba.serchcrowd.state.batch.update
//
// 暂停或启用单品推广搜索人群溢价
type TaobaosimbaserchcrowdstatebatchupdateAPIRequest struct {
	model.Params
	// 需要修改出价的人群包id,批量传入时用,分割
	_adgroupCrowdIds []string
	// 被操作者的淘宝昵称
	_nick string
	// 推广单元id
	_adgroupId int64
	// 人群状态,0:暂停;1:启用
	_state int64
}

// NewTaobaosimbaserchcrowdstatebatchupdateRequest 初始化TaobaosimbaserchcrowdstatebatchupdateAPIRequest对象
func NewTaobaosimbaserchcrowdstatebatchupdateRequest() *TaobaosimbaserchcrowdstatebatchupdateAPIRequest {
	return &TaobaosimbaserchcrowdstatebatchupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosimbaserchcrowdstatebatchupdateAPIRequest) GetApiMethodName() string {
	return "taobao.simba.serchcrowd.state.batch.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosimbaserchcrowdstatebatchupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosimbaserchcrowdstatebatchupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAdgroupCrowdIds is AdgroupCrowdIds Setter
// 需要修改出价的人群包id,批量传入时用,分割
func (r *TaobaosimbaserchcrowdstatebatchupdateAPIRequest) SetAdgroupCrowdIds(_adgroupCrowdIds []string) error {
	r._adgroupCrowdIds = _adgroupCrowdIds
	r.Set("adgroup_crowd_ids", _adgroupCrowdIds)
	return nil
}

// GetAdgroupCrowdIds AdgroupCrowdIds Getter
func (r TaobaosimbaserchcrowdstatebatchupdateAPIRequest) GetAdgroupCrowdIds() []string {
	return r._adgroupCrowdIds
}

// SetNick is Nick Setter
// 被操作者的淘宝昵称
func (r *TaobaosimbaserchcrowdstatebatchupdateAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaosimbaserchcrowdstatebatchupdateAPIRequest) GetNick() string {
	return r._nick
}

// SetAdgroupId is AdgroupId Setter
// 推广单元id
func (r *TaobaosimbaserchcrowdstatebatchupdateAPIRequest) SetAdgroupId(_adgroupId int64) error {
	r._adgroupId = _adgroupId
	r.Set("adgroup_id", _adgroupId)
	return nil
}

// GetAdgroupId AdgroupId Getter
func (r TaobaosimbaserchcrowdstatebatchupdateAPIRequest) GetAdgroupId() int64 {
	return r._adgroupId
}

// SetState is State Setter
// 人群状态,0:暂停;1:启用
func (r *TaobaosimbaserchcrowdstatebatchupdateAPIRequest) SetState(_state int64) error {
	r._state = _state
	r.Set("state", _state)
	return nil
}

// GetState State Getter
func (r TaobaosimbaserchcrowdstatebatchupdateAPIRequest) GetState() int64 {
	return r._state
}
