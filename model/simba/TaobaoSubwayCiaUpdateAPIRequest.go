package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSubwayCiaUpdateAPIRequest 批量修改单元智能出价 API请求
// taobao.subway.cia.update
//
// 批量修改直通车推广单元的智能出价配置
type TaobaoSubwayCiaUpdateAPIRequest struct {
	model.Params
	// 系统自动生成
	_ciaConfigs []CiaUpdateDto
	// 主人昵称
	_nick string
}

// NewTaobaoSubwayCiaUpdateRequest 初始化TaobaoSubwayCiaUpdateAPIRequest对象
func NewTaobaoSubwayCiaUpdateRequest() *TaobaoSubwayCiaUpdateAPIRequest {
	return &TaobaoSubwayCiaUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSubwayCiaUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.subway.cia.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSubwayCiaUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetCiaConfigs is CiaConfigs Setter
// 系统自动生成
func (r *TaobaoSubwayCiaUpdateAPIRequest) SetCiaConfigs(_ciaConfigs []CiaUpdateDto) error {
	r._ciaConfigs = _ciaConfigs
	r.Set("cia_configs", _ciaConfigs)
	return nil
}

// GetCiaConfigs CiaConfigs Getter
func (r TaobaoSubwayCiaUpdateAPIRequest) GetCiaConfigs() []CiaUpdateDto {
	return r._ciaConfigs
}

// SetNick is Nick Setter
// 主人昵称
func (r *TaobaoSubwayCiaUpdateAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaoSubwayCiaUpdateAPIRequest) GetNick() string {
	return r._nick
}
