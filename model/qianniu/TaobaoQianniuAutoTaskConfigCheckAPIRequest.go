package qianniu

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQianniuAutoTaskConfigCheckAPIRequest 自动化任务设置校验 API请求
// taobao.qianniu.auto.task.config.check
//
// 校验自动化任务配置
type TaobaoQianniuAutoTaskConfigCheckAPIRequest struct {
	model.Params
	// 场景，例如：催拍= RemindBuyBot
	_autoTask string
}

// NewTaobaoQianniuAutoTaskConfigCheckRequest 初始化TaobaoQianniuAutoTaskConfigCheckAPIRequest对象
func NewTaobaoQianniuAutoTaskConfigCheckRequest() *TaobaoQianniuAutoTaskConfigCheckAPIRequest {
	return &TaobaoQianniuAutoTaskConfigCheckAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoQianniuAutoTaskConfigCheckAPIRequest) Reset() {
	r._autoTask = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQianniuAutoTaskConfigCheckAPIRequest) GetApiMethodName() string {
	return "taobao.qianniu.auto.task.config.check"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQianniuAutoTaskConfigCheckAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoQianniuAutoTaskConfigCheckAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAutoTask is AutoTask Setter
// 场景，例如：催拍= RemindBuyBot
func (r *TaobaoQianniuAutoTaskConfigCheckAPIRequest) SetAutoTask(_autoTask string) error {
	r._autoTask = _autoTask
	r.Set("auto_task", _autoTask)
	return nil
}

// GetAutoTask AutoTask Getter
func (r TaobaoQianniuAutoTaskConfigCheckAPIRequest) GetAutoTask() string {
	return r._autoTask
}

var poolTaobaoQianniuAutoTaskConfigCheckAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoQianniuAutoTaskConfigCheckRequest()
	},
}

// GetTaobaoQianniuAutoTaskConfigCheckRequest 从 sync.Pool 获取 TaobaoQianniuAutoTaskConfigCheckAPIRequest
func GetTaobaoQianniuAutoTaskConfigCheckAPIRequest() *TaobaoQianniuAutoTaskConfigCheckAPIRequest {
	return poolTaobaoQianniuAutoTaskConfigCheckAPIRequest.Get().(*TaobaoQianniuAutoTaskConfigCheckAPIRequest)
}

// ReleaseTaobaoQianniuAutoTaskConfigCheckAPIRequest 将 TaobaoQianniuAutoTaskConfigCheckAPIRequest 放入 sync.Pool
func ReleaseTaobaoQianniuAutoTaskConfigCheckAPIRequest(v *TaobaoQianniuAutoTaskConfigCheckAPIRequest) {
	v.Reset()
	poolTaobaoQianniuAutoTaskConfigCheckAPIRequest.Put(v)
}
