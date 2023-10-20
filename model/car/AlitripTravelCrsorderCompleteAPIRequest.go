package car

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitriptravelcrsordercompleteAPIRequest CRS接送机商家服务完成接口 API请求
// alitrip.travel.crsorder.complete
//
// 提供给CRS接送机商家的服务完成回调接口
type AlitriptravelcrsordercompleteAPIRequest struct {
	model.Params
	// 请求对象
	_crsOrderCompleteParam *CrsOrderCompleteParam
}

// NewAlitriptravelcrsordercompleteRequest 初始化AlitriptravelcrsordercompleteAPIRequest对象
func NewAlitriptravelcrsordercompleteRequest() *AlitriptravelcrsordercompleteAPIRequest {
	return &AlitriptravelcrsordercompleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitriptravelcrsordercompleteAPIRequest) GetApiMethodName() string {
	return "alitrip.travel.crsorder.complete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitriptravelcrsordercompleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitriptravelcrsordercompleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCrsOrderCompleteParam is CrsOrderCompleteParam Setter
// 请求对象
func (r *AlitriptravelcrsordercompleteAPIRequest) SetCrsOrderCompleteParam(_crsOrderCompleteParam *CrsOrderCompleteParam) error {
	r._crsOrderCompleteParam = _crsOrderCompleteParam
	r.Set("crs_order_complete_param", _crsOrderCompleteParam)
	return nil
}

// GetCrsOrderCompleteParam CrsOrderCompleteParam Getter
func (r AlitriptravelcrsordercompleteAPIRequest) GetCrsOrderCompleteParam() *CrsOrderCompleteParam {
	return r._crsOrderCompleteParam
}
