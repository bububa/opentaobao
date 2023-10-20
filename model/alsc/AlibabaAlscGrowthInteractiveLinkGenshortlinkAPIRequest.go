package alsc

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscGrowthInteractiveLinkGenshortlinkAPIRequest 短链接口 API请求
// alibaba.alsc.growth.interactive.link.genshortlink
//
// 短链接口
type AlibabaAlscGrowthInteractiveLinkGenshortlinkAPIRequest struct {
	model.Params
	// 入参
	_param *GenShortLinkRequest
}

// NewAlibabaAlscGrowthInteractiveLinkGenshortlinkRequest 初始化AlibabaAlscGrowthInteractiveLinkGenshortlinkAPIRequest对象
func NewAlibabaAlscGrowthInteractiveLinkGenshortlinkRequest() *AlibabaAlscGrowthInteractiveLinkGenshortlinkAPIRequest {
	return &AlibabaAlscGrowthInteractiveLinkGenshortlinkAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlscGrowthInteractiveLinkGenshortlinkAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscGrowthInteractiveLinkGenshortlinkAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.growth.interactive.link.genshortlink"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscGrowthInteractiveLinkGenshortlinkAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlscGrowthInteractiveLinkGenshortlinkAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 入参
func (r *AlibabaAlscGrowthInteractiveLinkGenshortlinkAPIRequest) SetParam(_param *GenShortLinkRequest) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaAlscGrowthInteractiveLinkGenshortlinkAPIRequest) GetParam() *GenShortLinkRequest {
	return r._param
}

var poolAlibabaAlscGrowthInteractiveLinkGenshortlinkAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlscGrowthInteractiveLinkGenshortlinkRequest()
	},
}

// GetAlibabaAlscGrowthInteractiveLinkGenshortlinkRequest 从 sync.Pool 获取 AlibabaAlscGrowthInteractiveLinkGenshortlinkAPIRequest
func GetAlibabaAlscGrowthInteractiveLinkGenshortlinkAPIRequest() *AlibabaAlscGrowthInteractiveLinkGenshortlinkAPIRequest {
	return poolAlibabaAlscGrowthInteractiveLinkGenshortlinkAPIRequest.Get().(*AlibabaAlscGrowthInteractiveLinkGenshortlinkAPIRequest)
}

// ReleaseAlibabaAlscGrowthInteractiveLinkGenshortlinkAPIRequest 将 AlibabaAlscGrowthInteractiveLinkGenshortlinkAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlscGrowthInteractiveLinkGenshortlinkAPIRequest(v *AlibabaAlscGrowthInteractiveLinkGenshortlinkAPIRequest) {
	v.Reset()
	poolAlibabaAlscGrowthInteractiveLinkGenshortlinkAPIRequest.Put(v)
}
