package tmallcar

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallaliautometareceiveAPIRequest 汽车说明书元数据上传 API请求
// tmall.aliauto.meta.receive
//
// 天猫汽车对外提供的汽车资源元数据上传接口
type TmallaliautometareceiveAPIRequest struct {
	model.Params
	// 元数据参数集
	_command *ResourceMetaCommand
}

// NewTmallaliautometareceiveRequest 初始化TmallaliautometareceiveAPIRequest对象
func NewTmallaliautometareceiveRequest() *TmallaliautometareceiveAPIRequest {
	return &TmallaliautometareceiveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallaliautometareceiveAPIRequest) GetApiMethodName() string {
	return "tmall.aliauto.meta.receive"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallaliautometareceiveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallaliautometareceiveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCommand is Command Setter
// 元数据参数集
func (r *TmallaliautometareceiveAPIRequest) SetCommand(_command *ResourceMetaCommand) error {
	r._command = _command
	r.Set("command", _command)
	return nil
}

// GetCommand Command Getter
func (r TmallaliautometareceiveAPIRequest) GetCommand() *ResourceMetaCommand {
	return r._command
}
