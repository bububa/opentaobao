package user

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaominiappmesssagenormalsendAPIRequest 轻店铺下行普通消息给用户 API请求
// taobao.miniapp.messsage.normal.send
//
// 小程序下行单个普通消息
type TaobaominiappmesssagenormalsendAPIRequest struct {
	model.Params
	// 普通消息结构
	_param *DownNormalMessageDto
}

// NewTaobaominiappmesssagenormalsendRequest 初始化TaobaominiappmesssagenormalsendAPIRequest对象
func NewTaobaominiappmesssagenormalsendRequest() *TaobaominiappmesssagenormalsendAPIRequest {
	return &TaobaominiappmesssagenormalsendAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaominiappmesssagenormalsendAPIRequest) GetApiMethodName() string {
	return "taobao.miniapp.messsage.normal.send"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaominiappmesssagenormalsendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaominiappmesssagenormalsendAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 普通消息结构
func (r *TaobaominiappmesssagenormalsendAPIRequest) SetParam(_param *DownNormalMessageDto) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r TaobaominiappmesssagenormalsendAPIRequest) GetParam() *DownNormalMessageDto {
	return r._param
}
