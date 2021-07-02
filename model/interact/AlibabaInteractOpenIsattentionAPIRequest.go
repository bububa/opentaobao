package interact

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractOpenIsattentionAPIRequest 判断用户是否收藏某个店铺 API请求
// alibaba.interact.open.isattention
//
// 判断用户是否收藏某个店铺
type AlibabaInteractOpenIsattentionAPIRequest struct {
	model.Params
	// 1
	_param0 int64
}

// NewAlibabaInteractOpenIsattentionRequest 初始化AlibabaInteractOpenIsattentionAPIRequest对象
func NewAlibabaInteractOpenIsattentionRequest() *AlibabaInteractOpenIsattentionAPIRequest {
	return &AlibabaInteractOpenIsattentionAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaInteractOpenIsattentionAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.open.isattention"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaInteractOpenIsattentionAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Param0 Setter
// 1
func (r *AlibabaInteractOpenIsattentionAPIRequest) SetParam0(_param0 int64) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// Get Param0 Getter
func (r AlibabaInteractOpenIsattentionAPIRequest) GetParam0() int64 {
	return r._param0
}
