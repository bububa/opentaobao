package pentraprism

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPentaprismTaskTriggerAPIRequest 推进单条任务进度 API请求
// taobao.pentaprism.task.trigger
//
// 外网用户推进单条五棱镜任务进度
type TaobaoPentaprismTaskTriggerAPIRequest struct {
	model.Params
	// TOP接口标准入参
	_openPo *OpenTaskPo
}

// NewTaobaoPentaprismTaskTriggerRequest 初始化TaobaoPentaprismTaskTriggerAPIRequest对象
func NewTaobaoPentaprismTaskTriggerRequest() *TaobaoPentaprismTaskTriggerAPIRequest {
	return &TaobaoPentaprismTaskTriggerAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoPentaprismTaskTriggerAPIRequest) GetApiMethodName() string {
	return "taobao.pentaprism.task.trigger"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoPentaprismTaskTriggerAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is OpenPo Setter
// TOP接口标准入参
func (r *TaobaoPentaprismTaskTriggerAPIRequest) SetOpenPo(_openPo *OpenTaskPo) error {
	r._openPo = _openPo
	r.Set("open_po", _openPo)
	return nil
}

// Get OpenPo Getter
func (r TaobaoPentaprismTaskTriggerAPIRequest) GetOpenPo() *OpenTaskPo {
	return r._openPo
}
