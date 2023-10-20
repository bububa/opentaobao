package nrt

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallnrtstallsynchronizeAPIRequest 摊位信息同步 API请求
// tmall.nrt.stall.synchronize
//
// 摊位信息同步
type TmallnrtstallsynchronizeAPIRequest struct {
	model.Params
	// 参数对象
	_stall *NrtStoreDto
}

// NewTmallnrtstallsynchronizeRequest 初始化TmallnrtstallsynchronizeAPIRequest对象
func NewTmallnrtstallsynchronizeRequest() *TmallnrtstallsynchronizeAPIRequest {
	return &TmallnrtstallsynchronizeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallnrtstallsynchronizeAPIRequest) GetApiMethodName() string {
	return "tmall.nrt.stall.synchronize"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallnrtstallsynchronizeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallnrtstallsynchronizeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStall is Stall Setter
// 参数对象
func (r *TmallnrtstallsynchronizeAPIRequest) SetStall(_stall *NrtStoreDto) error {
	r._stall = _stall
	r.Set("stall", _stall)
	return nil
}

// GetStall Stall Getter
func (r TmallnrtstallsynchronizeAPIRequest) GetStall() *NrtStoreDto {
	return r._stall
}
