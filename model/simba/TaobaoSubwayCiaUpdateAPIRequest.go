package simba

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSubwayCiaUpdateAPIRequest) Reset() {
	r._ciaConfigs = r._ciaConfigs[:0]
	r._nick = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSubwayCiaUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.subway.cia.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSubwayCiaUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSubwayCiaUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolTaobaoSubwayCiaUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSubwayCiaUpdateRequest()
	},
}

// GetTaobaoSubwayCiaUpdateRequest 从 sync.Pool 获取 TaobaoSubwayCiaUpdateAPIRequest
func GetTaobaoSubwayCiaUpdateAPIRequest() *TaobaoSubwayCiaUpdateAPIRequest {
	return poolTaobaoSubwayCiaUpdateAPIRequest.Get().(*TaobaoSubwayCiaUpdateAPIRequest)
}

// ReleaseTaobaoSubwayCiaUpdateAPIRequest 将 TaobaoSubwayCiaUpdateAPIRequest 放入 sync.Pool
func ReleaseTaobaoSubwayCiaUpdateAPIRequest(v *TaobaoSubwayCiaUpdateAPIRequest) {
	v.Reset()
	poolTaobaoSubwayCiaUpdateAPIRequest.Put(v)
}
