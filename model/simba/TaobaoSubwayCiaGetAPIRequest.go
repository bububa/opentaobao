package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSubwayCiaGetAPIRequest 查询单元智能出价信息 API请求
// taobao.subway.cia.get
//
// 查询单元智能出价信息
type TaobaoSubwayCiaGetAPIRequest struct {
	model.Params
	// 主人昵称
	_nick string
	// 推广组Id
	_adgroupId int64
}

// NewTaobaoSubwayCiaGetRequest 初始化TaobaoSubwayCiaGetAPIRequest对象
func NewTaobaoSubwayCiaGetRequest() *TaobaoSubwayCiaGetAPIRequest {
	return &TaobaoSubwayCiaGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSubwayCiaGetAPIRequest) GetApiMethodName() string {
	return "taobao.subway.cia.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSubwayCiaGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Nick Setter
// 主人昵称
func (r *TaobaoSubwayCiaGetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// Get Nick Getter
func (r TaobaoSubwayCiaGetAPIRequest) GetNick() string {
	return r._nick
}

// Set is AdgroupId Setter
// 推广组Id
func (r *TaobaoSubwayCiaGetAPIRequest) SetAdgroupId(_adgroupId int64) error {
	r._adgroupId = _adgroupId
	r.Set("adgroup_id", _adgroupId)
	return nil
}

// Get AdgroupId Getter
func (r TaobaoSubwayCiaGetAPIRequest) GetAdgroupId() int64 {
	return r._adgroupId
}
