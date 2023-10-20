package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSubwayWordpackageGetAPIRequest 获取词包列表 API请求
// taobao.subway.wordpackage.get
//
// 获取流量智选、捡漏词包等词包列表
type TaobaoSubwayWordpackageGetAPIRequest struct {
	model.Params
	// 主人昵称
	_nick string
	// 推广组id
	_adgroupId int64
}

// NewTaobaoSubwayWordpackageGetRequest 初始化TaobaoSubwayWordpackageGetAPIRequest对象
func NewTaobaoSubwayWordpackageGetRequest() *TaobaoSubwayWordpackageGetAPIRequest {
	return &TaobaoSubwayWordpackageGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSubwayWordpackageGetAPIRequest) GetApiMethodName() string {
	return "taobao.subway.wordpackage.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSubwayWordpackageGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSubwayWordpackageGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 主人昵称
func (r *TaobaoSubwayWordpackageGetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaoSubwayWordpackageGetAPIRequest) GetNick() string {
	return r._nick
}

// SetAdgroupId is AdgroupId Setter
// 推广组id
func (r *TaobaoSubwayWordpackageGetAPIRequest) SetAdgroupId(_adgroupId int64) error {
	r._adgroupId = _adgroupId
	r.Set("adgroup_id", _adgroupId)
	return nil
}

// GetAdgroupId AdgroupId Getter
func (r TaobaoSubwayWordpackageGetAPIRequest) GetAdgroupId() int64 {
	return r._adgroupId
}
