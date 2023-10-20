package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosimbasalestarcreativedeleteAPIRequest (新)销量明星删除创意相关接口 API请求
// taobao.simba.salestar.creative.delete
//
// 删除一个创意
type TaobaosimbasalestarcreativedeleteAPIRequest struct {
	model.Params
	// 创意Id
	_creativeId int64
}

// NewTaobaosimbasalestarcreativedeleteRequest 初始化TaobaosimbasalestarcreativedeleteAPIRequest对象
func NewTaobaosimbasalestarcreativedeleteRequest() *TaobaosimbasalestarcreativedeleteAPIRequest {
	return &TaobaosimbasalestarcreativedeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosimbasalestarcreativedeleteAPIRequest) GetApiMethodName() string {
	return "taobao.simba.salestar.creative.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosimbasalestarcreativedeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosimbasalestarcreativedeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCreativeId is CreativeId Setter
// 创意Id
func (r *TaobaosimbasalestarcreativedeleteAPIRequest) SetCreativeId(_creativeId int64) error {
	r._creativeId = _creativeId
	r.Set("creative_id", _creativeId)
	return nil
}

// GetCreativeId CreativeId Getter
func (r TaobaosimbasalestarcreativedeleteAPIRequest) GetCreativeId() int64 {
	return r._creativeId
}
