package pentraprism

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaopentaprismtasktriggerAPIRequest 推进单条任务进度 API请求
// taobao.pentaprism.task.trigger
//
// 外网用户推进单条五棱镜任务进度
type TaobaopentaprismtasktriggerAPIRequest struct {
	model.Params
	// TOP接口标准入参
	_openPo *OpenTaskPo
}

// NewTaobaopentaprismtasktriggerRequest 初始化TaobaopentaprismtasktriggerAPIRequest对象
func NewTaobaopentaprismtasktriggerRequest() *TaobaopentaprismtasktriggerAPIRequest {
	return &TaobaopentaprismtasktriggerAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaopentaprismtasktriggerAPIRequest) GetApiMethodName() string {
	return "taobao.pentaprism.task.trigger"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaopentaprismtasktriggerAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaopentaprismtasktriggerAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOpenPo is OpenPo Setter
// TOP接口标准入参
func (r *TaobaopentaprismtasktriggerAPIRequest) SetOpenPo(_openPo *OpenTaskPo) error {
	r._openPo = _openPo
	r.Set("open_po", _openPo)
	return nil
}

// GetOpenPo OpenPo Getter
func (r TaobaopentaprismtasktriggerAPIRequest) GetOpenPo() *OpenTaskPo {
	return r._openPo
}
