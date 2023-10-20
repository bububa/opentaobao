package dt

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoadsdataimportAPIRequest 数据导入 API请求
// taobao.ads.data.import
//
// 数据导入
type TaobaoadsdataimportAPIRequest struct {
	model.Params
	// 请求入参
	_param0 *RequesterDataJobSaveCmd
}

// NewTaobaoadsdataimportRequest 初始化TaobaoadsdataimportAPIRequest对象
func NewTaobaoadsdataimportRequest() *TaobaoadsdataimportAPIRequest {
	return &TaobaoadsdataimportAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoadsdataimportAPIRequest) GetApiMethodName() string {
	return "taobao.ads.data.import"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoadsdataimportAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoadsdataimportAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 请求入参
func (r *TaobaoadsdataimportAPIRequest) SetParam0(_param0 *RequesterDataJobSaveCmd) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r TaobaoadsdataimportAPIRequest) GetParam0() *RequesterDataJobSaveCmd {
	return r._param0
}
