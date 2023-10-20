package tmallnr

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallNrFulfillLogisticsConsignAPIRequest 同城配门店备货通知 API请求
// tmall.nr.fulfill.logistics.consign
//
// 同城配业务备货通知，商家告诉平台门店的货已经准备好，可以发货了；
type TmallNrFulfillLogisticsConsignAPIRequest struct {
	model.Params
	// 入参对象
	_param0 *NrStoreGoodsReadyReqDto
}

// NewTmallNrFulfillLogisticsConsignRequest 初始化TmallNrFulfillLogisticsConsignAPIRequest对象
func NewTmallNrFulfillLogisticsConsignRequest() *TmallNrFulfillLogisticsConsignAPIRequest {
	return &TmallNrFulfillLogisticsConsignAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallNrFulfillLogisticsConsignAPIRequest) GetApiMethodName() string {
	return "tmall.nr.fulfill.logistics.consign"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallNrFulfillLogisticsConsignAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallNrFulfillLogisticsConsignAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 入参对象
func (r *TmallNrFulfillLogisticsConsignAPIRequest) SetParam0(_param0 *NrStoreGoodsReadyReqDto) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r TmallNrFulfillLogisticsConsignAPIRequest) GetParam0() *NrStoreGoodsReadyReqDto {
	return r._param0
}
