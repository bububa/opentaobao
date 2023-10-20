package tmallcar

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallcarcarefreedetailgetAPIRequest 查询业务单信息 API请求
// tmall.car.carefree.detail.get
//
// 查询业务单信息
type TmallcarcarefreedetailgetAPIRequest struct {
	model.Params
	// 查询参数对象
	_param0 *CarefreeDetailQueryReq
}

// NewTmallcarcarefreedetailgetRequest 初始化TmallcarcarefreedetailgetAPIRequest对象
func NewTmallcarcarefreedetailgetRequest() *TmallcarcarefreedetailgetAPIRequest {
	return &TmallcarcarefreedetailgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallcarcarefreedetailgetAPIRequest) GetApiMethodName() string {
	return "tmall.car.carefree.detail.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallcarcarefreedetailgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallcarcarefreedetailgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 查询参数对象
func (r *TmallcarcarefreedetailgetAPIRequest) SetParam0(_param0 *CarefreeDetailQueryReq) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r TmallcarcarefreedetailgetAPIRequest) GetParam0() *CarefreeDetailQueryReq {
	return r._param0
}
