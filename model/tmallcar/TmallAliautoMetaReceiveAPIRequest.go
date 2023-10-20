package tmallcar

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallAliautoMetaReceiveAPIRequest 汽车说明书元数据上传 API请求
// tmall.aliauto.meta.receive
//
// 天猫汽车对外提供的汽车资源元数据上传接口
type TmallAliautoMetaReceiveAPIRequest struct {
	model.Params
	// 元数据参数集
	_command *ResourceMetaCommand
}

// NewTmallAliautoMetaReceiveRequest 初始化TmallAliautoMetaReceiveAPIRequest对象
func NewTmallAliautoMetaReceiveRequest() *TmallAliautoMetaReceiveAPIRequest {
	return &TmallAliautoMetaReceiveAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallAliautoMetaReceiveAPIRequest) Reset() {
	r._command = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallAliautoMetaReceiveAPIRequest) GetApiMethodName() string {
	return "tmall.aliauto.meta.receive"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallAliautoMetaReceiveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallAliautoMetaReceiveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCommand is Command Setter
// 元数据参数集
func (r *TmallAliautoMetaReceiveAPIRequest) SetCommand(_command *ResourceMetaCommand) error {
	r._command = _command
	r.Set("command", _command)
	return nil
}

// GetCommand Command Getter
func (r TmallAliautoMetaReceiveAPIRequest) GetCommand() *ResourceMetaCommand {
	return r._command
}

var poolTmallAliautoMetaReceiveAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallAliautoMetaReceiveRequest()
	},
}

// GetTmallAliautoMetaReceiveRequest 从 sync.Pool 获取 TmallAliautoMetaReceiveAPIRequest
func GetTmallAliautoMetaReceiveAPIRequest() *TmallAliautoMetaReceiveAPIRequest {
	return poolTmallAliautoMetaReceiveAPIRequest.Get().(*TmallAliautoMetaReceiveAPIRequest)
}

// ReleaseTmallAliautoMetaReceiveAPIRequest 将 TmallAliautoMetaReceiveAPIRequest 放入 sync.Pool
func ReleaseTmallAliautoMetaReceiveAPIRequest(v *TmallAliautoMetaReceiveAPIRequest) {
	v.Reset()
	poolTmallAliautoMetaReceiveAPIRequest.Put(v)
}
