package nrt

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallnrtassetauthorizationaddAPIRequest 增加数据权限授权 API请求
// tmall.nrt.asset.authorization.add
//
// 增加数据权限授权
type TmallnrtassetauthorizationaddAPIRequest struct {
	model.Params
	// 添加数据权限请求
	_topAssetDataAuthReqDto *TopAssetDataAuthReqDto
}

// NewTmallnrtassetauthorizationaddRequest 初始化TmallnrtassetauthorizationaddAPIRequest对象
func NewTmallnrtassetauthorizationaddRequest() *TmallnrtassetauthorizationaddAPIRequest {
	return &TmallnrtassetauthorizationaddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallnrtassetauthorizationaddAPIRequest) GetApiMethodName() string {
	return "tmall.nrt.asset.authorization.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallnrtassetauthorizationaddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallnrtassetauthorizationaddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopAssetDataAuthReqDto is TopAssetDataAuthReqDto Setter
// 添加数据权限请求
func (r *TmallnrtassetauthorizationaddAPIRequest) SetTopAssetDataAuthReqDto(_topAssetDataAuthReqDto *TopAssetDataAuthReqDto) error {
	r._topAssetDataAuthReqDto = _topAssetDataAuthReqDto
	r.Set("top_asset_data_auth_req_dto", _topAssetDataAuthReqDto)
	return nil
}

// GetTopAssetDataAuthReqDto TopAssetDataAuthReqDto Getter
func (r TmallnrtassetauthorizationaddAPIRequest) GetTopAssetDataAuthReqDto() *TopAssetDataAuthReqDto {
	return r._topAssetDataAuthReqDto
}
