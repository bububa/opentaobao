package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunostvpubadmindiccontroltaskaddAPIRequest 新增停开服任务 API请求
// yunos.tvpubadmin.diccontroltask.add
//
// 新增停开服任务
type YunostvpubadmindiccontroltaskaddAPIRequest struct {
	model.Params
	// 任务信息
	_task *DicControlTaskDo
}

// NewYunostvpubadmindiccontroltaskaddRequest 初始化YunostvpubadmindiccontroltaskaddAPIRequest对象
func NewYunostvpubadmindiccontroltaskaddRequest() *YunostvpubadmindiccontroltaskaddAPIRequest {
	return &YunostvpubadmindiccontroltaskaddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunostvpubadmindiccontroltaskaddAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.diccontroltask.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunostvpubadmindiccontroltaskaddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunostvpubadmindiccontroltaskaddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTask is Task Setter
// 任务信息
func (r *YunostvpubadmindiccontroltaskaddAPIRequest) SetTask(_task *DicControlTaskDo) error {
	r._task = _task
	r.Set("task", _task)
	return nil
}

// GetTask Task Getter
func (r YunostvpubadmindiccontroltaskaddAPIRequest) GetTask() *DicControlTaskDo {
	return r._task
}
