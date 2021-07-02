package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleRecycleSpuTemplateModifyAPIRequest 闲鱼接收回收商spu模板挂载信息 API请求
// alibaba.idle.recycle.spu.template.modify
//
// 闲鱼接收回收商spu模板挂载信息
type AlibabaIdleRecycleSpuTemplateModifyAPIRequest struct {
	model.Params
	// 回收商挂载模版信息
	_recycleSpuTemplate *RecycleSpuTemplate
}

// NewAlibabaIdleRecycleSpuTemplateModifyRequest 初始化AlibabaIdleRecycleSpuTemplateModifyAPIRequest对象
func NewAlibabaIdleRecycleSpuTemplateModifyRequest() *AlibabaIdleRecycleSpuTemplateModifyAPIRequest {
	return &AlibabaIdleRecycleSpuTemplateModifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleRecycleSpuTemplateModifyAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.recycle.spu.template.modify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleRecycleSpuTemplateModifyAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is RecycleSpuTemplate Setter
// 回收商挂载模版信息
func (r *AlibabaIdleRecycleSpuTemplateModifyAPIRequest) SetRecycleSpuTemplate(_recycleSpuTemplate *RecycleSpuTemplate) error {
	r._recycleSpuTemplate = _recycleSpuTemplate
	r.Set("recycle_spu_template", _recycleSpuTemplate)
	return nil
}

// Get RecycleSpuTemplate Getter
func (r AlibabaIdleRecycleSpuTemplateModifyAPIRequest) GetRecycleSpuTemplate() *RecycleSpuTemplate {
	return r._recycleSpuTemplate
}
