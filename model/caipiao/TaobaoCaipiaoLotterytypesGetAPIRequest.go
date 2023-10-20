package caipiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaocaipiaolotterytypesgetAPIRequest 获取可用的彩种列表 API请求
// taobao.caipiao.lotterytypes.get
//
// 获取彩票系统支持的可用于赠送的彩种列表
type TaobaocaipiaolotterytypesgetAPIRequest struct {
	model.Params
}

// NewTaobaocaipiaolotterytypesgetRequest 初始化TaobaocaipiaolotterytypesgetAPIRequest对象
func NewTaobaocaipiaolotterytypesgetRequest() *TaobaocaipiaolotterytypesgetAPIRequest {
	return &TaobaocaipiaolotterytypesgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaocaipiaolotterytypesgetAPIRequest) GetApiMethodName() string {
	return "taobao.caipiao.lotterytypes.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaocaipiaolotterytypesgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaocaipiaolotterytypesgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}
