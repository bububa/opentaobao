package alihouse

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseExistinghomeThreedimensionSyncAPIRequest 二手房3D户型信息同步 API请求
// alibaba.alihouse.existinghome.threedimension.sync
//
// 二手房3D户型信息同步
type AlibabaAlihouseExistinghomeThreedimensionSyncAPIRequest struct {
	model.Params
	// 系统自动生成
	_threeDimensionalDto *SyncExistingHouseThreeDimensionalDto
}

// NewAlibabaAlihouseExistinghomeThreedimensionSyncRequest 初始化AlibabaAlihouseExistinghomeThreedimensionSyncAPIRequest对象
func NewAlibabaAlihouseExistinghomeThreedimensionSyncRequest() *AlibabaAlihouseExistinghomeThreedimensionSyncAPIRequest {
	return &AlibabaAlihouseExistinghomeThreedimensionSyncAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihouseExistinghomeThreedimensionSyncAPIRequest) Reset() {
	r._threeDimensionalDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseExistinghomeThreedimensionSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.existinghome.threedimension.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseExistinghomeThreedimensionSyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseExistinghomeThreedimensionSyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetThreeDimensionalDto is ThreeDimensionalDto Setter
// 系统自动生成
func (r *AlibabaAlihouseExistinghomeThreedimensionSyncAPIRequest) SetThreeDimensionalDto(_threeDimensionalDto *SyncExistingHouseThreeDimensionalDto) error {
	r._threeDimensionalDto = _threeDimensionalDto
	r.Set("three_dimensional_dto", _threeDimensionalDto)
	return nil
}

// GetThreeDimensionalDto ThreeDimensionalDto Getter
func (r AlibabaAlihouseExistinghomeThreedimensionSyncAPIRequest) GetThreeDimensionalDto() *SyncExistingHouseThreeDimensionalDto {
	return r._threeDimensionalDto
}

var poolAlibabaAlihouseExistinghomeThreedimensionSyncAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihouseExistinghomeThreedimensionSyncRequest()
	},
}

// GetAlibabaAlihouseExistinghomeThreedimensionSyncRequest 从 sync.Pool 获取 AlibabaAlihouseExistinghomeThreedimensionSyncAPIRequest
func GetAlibabaAlihouseExistinghomeThreedimensionSyncAPIRequest() *AlibabaAlihouseExistinghomeThreedimensionSyncAPIRequest {
	return poolAlibabaAlihouseExistinghomeThreedimensionSyncAPIRequest.Get().(*AlibabaAlihouseExistinghomeThreedimensionSyncAPIRequest)
}

// ReleaseAlibabaAlihouseExistinghomeThreedimensionSyncAPIRequest 将 AlibabaAlihouseExistinghomeThreedimensionSyncAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihouseExistinghomeThreedimensionSyncAPIRequest(v *AlibabaAlihouseExistinghomeThreedimensionSyncAPIRequest) {
	v.Reset()
	poolAlibabaAlihouseExistinghomeThreedimensionSyncAPIRequest.Put(v)
}
