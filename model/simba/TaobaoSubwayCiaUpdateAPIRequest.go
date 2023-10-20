package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosubwayciaupdateAPIRequest 批量修改单元智能出价 API请求
// taobao.subway.cia.update
//
// 批量修改直通车推广单元的智能出价配置
type TaobaosubwayciaupdateAPIRequest struct {
	model.Params
	// 系统自动生成
	_ciaConfigs []CiaUpdateDto
	// 主人昵称
	_nick string
}

// NewTaobaosubwayciaupdateRequest 初始化TaobaosubwayciaupdateAPIRequest对象
func NewTaobaosubwayciaupdateRequest() *TaobaosubwayciaupdateAPIRequest {
	return &TaobaosubwayciaupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosubwayciaupdateAPIRequest) GetApiMethodName() string {
	return "taobao.subway.cia.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosubwayciaupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosubwayciaupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCiaConfigs is CiaConfigs Setter
// 系统自动生成
func (r *TaobaosubwayciaupdateAPIRequest) SetCiaConfigs(_ciaConfigs []CiaUpdateDto) error {
	r._ciaConfigs = _ciaConfigs
	r.Set("cia_configs", _ciaConfigs)
	return nil
}

// GetCiaConfigs CiaConfigs Getter
func (r TaobaosubwayciaupdateAPIRequest) GetCiaConfigs() []CiaUpdateDto {
	return r._ciaConfigs
}

// SetNick is Nick Setter
// 主人昵称
func (r *TaobaosubwayciaupdateAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaosubwayciaupdateAPIRequest) GetNick() string {
	return r._nick
}
