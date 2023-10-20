package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihouseexistinghomethreedimensionsyncAPIRequest 二手房3D户型信息同步 API请求
// alibaba.alihouse.existinghome.threedimension.sync
//
// 二手房3D户型信息同步
type AlibabaalihouseexistinghomethreedimensionsyncAPIRequest struct {
	model.Params
	// 系统自动生成
	_threeDimensionalDto *SyncExistingHouseThreeDimensionalDto
}

// NewAlibabaalihouseexistinghomethreedimensionsyncRequest 初始化AlibabaalihouseexistinghomethreedimensionsyncAPIRequest对象
func NewAlibabaalihouseexistinghomethreedimensionsyncRequest() *AlibabaalihouseexistinghomethreedimensionsyncAPIRequest {
	return &AlibabaalihouseexistinghomethreedimensionsyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihouseexistinghomethreedimensionsyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.existinghome.threedimension.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihouseexistinghomethreedimensionsyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihouseexistinghomethreedimensionsyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetThreeDimensionalDto is ThreeDimensionalDto Setter
// 系统自动生成
func (r *AlibabaalihouseexistinghomethreedimensionsyncAPIRequest) SetThreeDimensionalDto(_threeDimensionalDto *SyncExistingHouseThreeDimensionalDto) error {
	r._threeDimensionalDto = _threeDimensionalDto
	r.Set("three_dimensional_dto", _threeDimensionalDto)
	return nil
}

// GetThreeDimensionalDto ThreeDimensionalDto Getter
func (r AlibabaalihouseexistinghomethreedimensionsyncAPIRequest) GetThreeDimensionalDto() *SyncExistingHouseThreeDimensionalDto {
	return r._threeDimensionalDto
}
