package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosimbacreativedeleteAPIRequest 删除创意 API请求
// taobao.simba.creative.delete
//
// 删除一个创意
type TaobaosimbacreativedeleteAPIRequest struct {
	model.Params
	// 主人昵称
	_nick string
	// 创意Id
	_creativeId int64
}

// NewTaobaosimbacreativedeleteRequest 初始化TaobaosimbacreativedeleteAPIRequest对象
func NewTaobaosimbacreativedeleteRequest() *TaobaosimbacreativedeleteAPIRequest {
	return &TaobaosimbacreativedeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosimbacreativedeleteAPIRequest) GetApiMethodName() string {
	return "taobao.simba.creative.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosimbacreativedeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosimbacreativedeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 主人昵称
func (r *TaobaosimbacreativedeleteAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaosimbacreativedeleteAPIRequest) GetNick() string {
	return r._nick
}

// SetCreativeId is CreativeId Setter
// 创意Id
func (r *TaobaosimbacreativedeleteAPIRequest) SetCreativeId(_creativeId int64) error {
	r._creativeId = _creativeId
	r.Set("creative_id", _creativeId)
	return nil
}

// GetCreativeId CreativeId Getter
func (r TaobaosimbacreativedeleteAPIRequest) GetCreativeId() int64 {
	return r._creativeId
}
