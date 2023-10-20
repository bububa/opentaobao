package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidletenderbtobitemuploadAPIRequest 暗拍发布/编辑B2B商品 API请求
// alibaba.idle.tender.btob.item.upload
//
// 暗拍发布/编辑B2B商品
type AlibabaidletenderbtobitemuploadAPIRequest struct {
	model.Params
	// 参数
	_param0 *TenderItemUploadCmd
}

// NewAlibabaidletenderbtobitemuploadRequest 初始化AlibabaidletenderbtobitemuploadAPIRequest对象
func NewAlibabaidletenderbtobitemuploadRequest() *AlibabaidletenderbtobitemuploadAPIRequest {
	return &AlibabaidletenderbtobitemuploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaidletenderbtobitemuploadAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.tender.btob.item.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaidletenderbtobitemuploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaidletenderbtobitemuploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 参数
func (r *AlibabaidletenderbtobitemuploadAPIRequest) SetParam0(_param0 *TenderItemUploadCmd) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaidletenderbtobitemuploadAPIRequest) GetParam0() *TenderItemUploadCmd {
	return r._param0
}
