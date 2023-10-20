package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminDiccontroltaskAddAPIRequest 新增停开服任务 API请求
// yunos.tvpubadmin.diccontroltask.add
//
// 新增停开服任务
type YunosTvpubadminDiccontroltaskAddAPIRequest struct {
	model.Params
	// 任务信息
	_task *DicControlTaskDo
}

// NewYunosTvpubadminDiccontroltaskAddRequest 初始化YunosTvpubadminDiccontroltaskAddAPIRequest对象
func NewYunosTvpubadminDiccontroltaskAddRequest() *YunosTvpubadminDiccontroltaskAddAPIRequest {
	return &YunosTvpubadminDiccontroltaskAddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosTvpubadminDiccontroltaskAddAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.diccontroltask.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosTvpubadminDiccontroltaskAddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosTvpubadminDiccontroltaskAddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTask is Task Setter
// 任务信息
func (r *YunosTvpubadminDiccontroltaskAddAPIRequest) SetTask(_task *DicControlTaskDo) error {
	r._task = _task
	r.Set("task", _task)
	return nil
}

// GetTask Task Getter
func (r YunosTvpubadminDiccontroltaskAddAPIRequest) GetTask() *DicControlTaskDo {
	return r._task
}
