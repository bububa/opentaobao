package moscm

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMosGoodsBulkinputcspuAPIRequest 批量录入商品信息 API请求
// alibaba.mos.goods.bulkinputcspu
//
// 用于商品信息的批量导入到银泰商品中台
type AlibabaMosGoodsBulkinputcspuAPIRequest struct {
	model.Params
	// 录入商品信息列表（最大列表长度：20）
	_cspuInputDtoList []CspuInputDto
}

// NewAlibabaMosGoodsBulkinputcspuRequest 初始化AlibabaMosGoodsBulkinputcspuAPIRequest对象
func NewAlibabaMosGoodsBulkinputcspuRequest() *AlibabaMosGoodsBulkinputcspuAPIRequest {
	return &AlibabaMosGoodsBulkinputcspuAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaMosGoodsBulkinputcspuAPIRequest) Reset() {
	r._cspuInputDtoList = r._cspuInputDtoList[:0]
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMosGoodsBulkinputcspuAPIRequest) GetApiMethodName() string {
	return "alibaba.mos.goods.bulkinputcspu"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMosGoodsBulkinputcspuAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaMosGoodsBulkinputcspuAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCspuInputDtoList is CspuInputDtoList Setter
// 录入商品信息列表（最大列表长度：20）
func (r *AlibabaMosGoodsBulkinputcspuAPIRequest) SetCspuInputDtoList(_cspuInputDtoList []CspuInputDto) error {
	r._cspuInputDtoList = _cspuInputDtoList
	r.Set("cspu_input_dto_list", _cspuInputDtoList)
	return nil
}

// GetCspuInputDtoList CspuInputDtoList Getter
func (r AlibabaMosGoodsBulkinputcspuAPIRequest) GetCspuInputDtoList() []CspuInputDto {
	return r._cspuInputDtoList
}

var poolAlibabaMosGoodsBulkinputcspuAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaMosGoodsBulkinputcspuRequest()
	},
}

// GetAlibabaMosGoodsBulkinputcspuRequest 从 sync.Pool 获取 AlibabaMosGoodsBulkinputcspuAPIRequest
func GetAlibabaMosGoodsBulkinputcspuAPIRequest() *AlibabaMosGoodsBulkinputcspuAPIRequest {
	return poolAlibabaMosGoodsBulkinputcspuAPIRequest.Get().(*AlibabaMosGoodsBulkinputcspuAPIRequest)
}

// ReleaseAlibabaMosGoodsBulkinputcspuAPIRequest 将 AlibabaMosGoodsBulkinputcspuAPIRequest 放入 sync.Pool
func ReleaseAlibabaMosGoodsBulkinputcspuAPIRequest(v *AlibabaMosGoodsBulkinputcspuAPIRequest) {
	v.Reset()
	poolAlibabaMosGoodsBulkinputcspuAPIRequest.Put(v)
}
