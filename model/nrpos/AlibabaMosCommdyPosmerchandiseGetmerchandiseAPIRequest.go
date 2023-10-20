package nrpos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabamoscommdyposmerchandisegetmerchandiseAPIRequest 去前置机商品在线查询 API请求
// alibaba.mos.commdy.posmerchandise.getmerchandise
//
// 去前置机商品在线查询接口
type AlibabamoscommdyposmerchandisegetmerchandiseAPIRequest struct {
	model.Params
	// 查询参数列表
	_posMerchandiseList []QueryMerchandiseDto
}

// NewAlibabamoscommdyposmerchandisegetmerchandiseRequest 初始化AlibabamoscommdyposmerchandisegetmerchandiseAPIRequest对象
func NewAlibabamoscommdyposmerchandisegetmerchandiseRequest() *AlibabamoscommdyposmerchandisegetmerchandiseAPIRequest {
	return &AlibabamoscommdyposmerchandisegetmerchandiseAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabamoscommdyposmerchandisegetmerchandiseAPIRequest) GetApiMethodName() string {
	return "alibaba.mos.commdy.posmerchandise.getmerchandise"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabamoscommdyposmerchandisegetmerchandiseAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabamoscommdyposmerchandisegetmerchandiseAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPosMerchandiseList is PosMerchandiseList Setter
// 查询参数列表
func (r *AlibabamoscommdyposmerchandisegetmerchandiseAPIRequest) SetPosMerchandiseList(_posMerchandiseList []QueryMerchandiseDto) error {
	r._posMerchandiseList = _posMerchandiseList
	r.Set("pos_merchandise_list", _posMerchandiseList)
	return nil
}

// GetPosMerchandiseList PosMerchandiseList Getter
func (r AlibabamoscommdyposmerchandisegetmerchandiseAPIRequest) GetPosMerchandiseList() []QueryMerchandiseDto {
	return r._posMerchandiseList
}
