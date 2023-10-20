package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosubwaywordpackagegetAPIRequest 获取词包列表 API请求
// taobao.subway.wordpackage.get
//
// 获取流量智选、捡漏词包等词包列表
type TaobaosubwaywordpackagegetAPIRequest struct {
	model.Params
	// 主人昵称
	_nick string
	// 推广组id
	_adgroupId int64
}

// NewTaobaosubwaywordpackagegetRequest 初始化TaobaosubwaywordpackagegetAPIRequest对象
func NewTaobaosubwaywordpackagegetRequest() *TaobaosubwaywordpackagegetAPIRequest {
	return &TaobaosubwaywordpackagegetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosubwaywordpackagegetAPIRequest) GetApiMethodName() string {
	return "taobao.subway.wordpackage.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosubwaywordpackagegetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosubwaywordpackagegetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 主人昵称
func (r *TaobaosubwaywordpackagegetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaosubwaywordpackagegetAPIRequest) GetNick() string {
	return r._nick
}

// SetAdgroupId is AdgroupId Setter
// 推广组id
func (r *TaobaosubwaywordpackagegetAPIRequest) SetAdgroupId(_adgroupId int64) error {
	r._adgroupId = _adgroupId
	r.Set("adgroup_id", _adgroupId)
	return nil
}

// GetAdgroupId AdgroupId Getter
func (r TaobaosubwaywordpackagegetAPIRequest) GetAdgroupId() int64 {
	return r._adgroupId
}
