package tmallcar

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallAliautoMetaReceiveAPIRequest) GetApiMethodName() string {
	return "tmall.aliauto.meta.receive"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallAliautoMetaReceiveAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Command Setter
// 元数据参数集
func (r *TmallAliautoMetaReceiveAPIRequest) SetCommand(_command *ResourceMetaCommand) error {
	r._command = _command
	r.Set("command", _command)
	return nil
}

// Get Command Getter
func (r TmallAliautoMetaReceiveAPIRequest) GetCommand() *ResourceMetaCommand {
	return r._command
}
