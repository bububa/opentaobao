package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIdleRecycleSpuTemplateModifyAPIRequest
闲鱼接收回收商spu模板挂载信息 API请求
alibaba.idle.recycle.spu.template.modify

闲鱼接收回收商spu模板挂载信息 */
type AlibabaIdleRecycleSpuTemplateModifyAPIRequest struct {
	model.Params
	// 回收商挂载模版信息
	_recycleSpuTemplate *RecycleSpuTemplate
}

// New
