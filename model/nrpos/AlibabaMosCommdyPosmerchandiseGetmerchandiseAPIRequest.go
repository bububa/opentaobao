package nrpos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMosCommdyPosmerchandiseGetmerchandiseAPIRequest 去前置机商品在线查询 API请求
// alibaba.mos.commdy.posmerchandise.getmerchandise
//
// 去前置机商品在线查询接口
type AlibabaMosCommdyPosmerchandiseGetmerchandiseAPIRequest struct {
	model.Params
	// 查询参数列表
	_posMerchandiseList []QueryMerchandiseDto
}

// NewAlibabaMosCommdyPosmerchandiseGetmerchandiseRequest 初始化AlibabaMosCommdyPosmerchandiseGetmerchandiseAPIRequest对象
func NewAlibabaMosCommdyPosmerchandiseGetmerchandiseRequest() *AlibabaMosCommdyPosmerchandiseGetmerchandiseAPIRequest {
	return &AlibabaMosCommdyPosmerchandiseGetmerchandiseAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMosCommdyPosmerchandiseGetmerchandiseAPIRequest) GetApiMethodName() string {
	return "alibaba.mos.commdy.posmerchandise.getmerchandise"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMosCommdyPosmerchandiseGetmerchandiseAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaMosCommdyPosmerchandiseGetmerchandiseAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPosMerchandiseList is PosMerchandiseList Setter
// 查询参数列表
func (r *AlibabaMosCommdyPosmerchandiseGetmerchandiseAPIRequest) SetPosMerchandiseList(_posMerchandiseList []QueryMerchandiseDto) error {
	r._posMerchandiseList = _posMerchandiseList
	r.Set("pos_merchandise_list", _posMerchandiseList)
	return nil
}

// GetPosMerchandiseList PosMerchandiseList Getter
func (r AlibabaMosCommdyPosmerchandiseGetmerchandiseAPIRequest) GetPosMerchandiseList() []QueryMerchandiseDto {
	return r._posMerchandiseList
}
