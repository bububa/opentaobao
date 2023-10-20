package user

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaominiappmesssagereplyAPIRequest 轻店铺下行回复接口 API请求
// taobao.miniapp.messsage.reply
//
// 外部 isv 调用该进口来进行轻店铺消息的回复
type TaobaominiappmesssagereplyAPIRequest struct {
	model.Params
	// 入参
	_param *ReplyMessageDto
}

// NewTaobaominiappmesssagereplyRequest 初始化TaobaominiappmesssagereplyAPIRequest对象
func NewTaobaominiappmesssagereplyRequest() *TaobaominiappmesssagereplyAPIRequest {
	return &TaobaominiappmesssagereplyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaominiappmesssagereplyAPIRequest) GetApiMethodName() string {
	return "taobao.miniapp.messsage.reply"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaominiappmesssagereplyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaominiappmesssagereplyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 入参
func (r *TaobaominiappmesssagereplyAPIRequest) SetParam(_param *ReplyMessageDto) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r TaobaominiappmesssagereplyAPIRequest) GetParam() *ReplyMessageDto {
	return r._param
}
