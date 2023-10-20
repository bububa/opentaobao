package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidlerecyclesputemplatemodifyAPIRequest 闲鱼接收回收商spu模板挂载信息 API请求
// alibaba.idle.recycle.spu.template.modify
//
// 闲鱼接收回收商spu模板挂载信息
type AlibabaidlerecyclesputemplatemodifyAPIRequest struct {
	model.Params
	// 服务商支持的SPU挂载
	_recycleSpuTemplate *RecycleSpuTemplate
}

// NewAlibabaidlerecyclesputemplatemodifyRequest 初始化AlibabaidlerecyclesputemplatemodifyAPIRequest对象
func NewAlibabaidlerecyclesputemplatemodifyRequest() *AlibabaidlerecyclesputemplatemodifyAPIRequest {
	return &AlibabaidlerecyclesputemplatemodifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaidlerecyclesputemplatemodifyAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.recycle.spu.template.modify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaidlerecyclesputemplatemodifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaidlerecyclesputemplatemodifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRecycleSpuTemplate is RecycleSpuTemplate Setter
// 服务商支持的SPU挂载
func (r *AlibabaidlerecyclesputemplatemodifyAPIRequest) SetRecycleSpuTemplate(_recycleSpuTemplate *RecycleSpuTemplate) error {
	r._recycleSpuTemplate = _recycleSpuTemplate
	r.Set("recycle_spu_template", _recycleSpuTemplate)
	return nil
}

// GetRecycleSpuTemplate RecycleSpuTemplate Getter
func (r AlibabaidlerecyclesputemplatemodifyAPIRequest) GetRecycleSpuTemplate() *RecycleSpuTemplate {
	return r._recycleSpuTemplate
}
