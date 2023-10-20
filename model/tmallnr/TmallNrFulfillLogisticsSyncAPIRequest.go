package tmallnr

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallnrfulfilllogisticssyncAPIRequest 同城配物流信息回传 API请求
// tmall.nr.fulfill.logistics.sync
//
// 同城配业务物流信息回传，通过接口将物流信息同步给天猫
type TmallnrfulfilllogisticssyncAPIRequest struct {
	model.Params
	// 物流回传参数
	_param0 *NrLogisticsInfoSynReqDto
}

// NewTmallnrfulfilllogisticssyncRequest 初始化TmallnrfulfilllogisticssyncAPIRequest对象
func NewTmallnrfulfilllogisticssyncRequest() *TmallnrfulfilllogisticssyncAPIRequest {
	return &TmallnrfulfilllogisticssyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallnrfulfilllogisticssyncAPIRequest) GetApiMethodName() string {
	return "tmall.nr.fulfill.logistics.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallnrfulfilllogisticssyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallnrfulfilllogisticssyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 物流回传参数
func (r *TmallnrfulfilllogisticssyncAPIRequest) SetParam0(_param0 *NrLogisticsInfoSynReqDto) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r TmallnrfulfilllogisticssyncAPIRequest) GetParam0() *NrLogisticsInfoSynReqDto {
	return r._param0
}
