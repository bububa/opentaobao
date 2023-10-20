package pentraprism

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPentaprismTaskTriggerFromAPIRequest 任务进度推进（根据fromtoken） API请求
// taobao.pentaprism.task.trigger.from
//
// 外网用户推进单条五棱镜任务进度
type TaobaoPentaprismTaskTriggerFromAPIRequest struct {
	model.Params
	// TOP接口标准入参
	_openPo *OpenTaskPo
}

// NewTaobaoPentaprismTaskTriggerFromRequest 初始化TaobaoPentaprismTaskTriggerFromAPIRequest对象
func NewTaobaoPentaprismTaskTriggerFromRequest() *TaobaoPentaprismTaskTriggerFromAPIRequest {
	return &TaobaoPentaprismTaskTriggerFromAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoPentaprismTaskTriggerFromAPIRequest) GetApiMethodName() string {
	return "taobao.pentaprism.task.trigger.from"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoPentaprismTaskTriggerFromAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoPentaprismTaskTriggerFromAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOpenPo is OpenPo Setter
// TOP接口标准入参
func (r *TaobaoPentaprismTaskTriggerFromAPIRequest) SetOpenPo(_openPo *OpenTaskPo) error {
	r._openPo = _openPo
	r.Set("open_po", _openPo)
	return nil
}

// GetOpenPo OpenPo Getter
func (r TaobaoPentaprismTaskTriggerFromAPIRequest) GetOpenPo() *OpenTaskPo {
	return r._openPo
}
