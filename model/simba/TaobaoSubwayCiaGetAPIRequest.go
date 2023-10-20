package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosubwayciagetAPIRequest 查询单元智能出价信息 API请求
// taobao.subway.cia.get
//
// 查询单元智能出价信息
type TaobaosubwayciagetAPIRequest struct {
	model.Params
	// 主人昵称
	_nick string
	// 推广组Id
	_adgroupId int64
}

// NewTaobaosubwayciagetRequest 初始化TaobaosubwayciagetAPIRequest对象
func NewTaobaosubwayciagetRequest() *TaobaosubwayciagetAPIRequest {
	return &TaobaosubwayciagetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosubwayciagetAPIRequest) GetApiMethodName() string {
	return "taobao.subway.cia.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosubwayciagetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosubwayciagetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 主人昵称
func (r *TaobaosubwayciagetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaosubwayciagetAPIRequest) GetNick() string {
	return r._nick
}

// SetAdgroupId is AdgroupId Setter
// 推广组Id
func (r *TaobaosubwayciagetAPIRequest) SetAdgroupId(_adgroupId int64) error {
	r._adgroupId = _adgroupId
	r.Set("adgroup_id", _adgroupId)
	return nil
}

// GetAdgroupId AdgroupId Getter
func (r TaobaosubwayciagetAPIRequest) GetAdgroupId() int64 {
	return r._adgroupId
}
