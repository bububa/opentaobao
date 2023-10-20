package nrt

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallnrtassetauthorizationdeleteAPIRequest 移除资产数据权限授权关系 API请求
// tmall.nrt.asset.authorization.delete
//
// 移除资产数据权限授权关系
type TmallnrtassetauthorizationdeleteAPIRequest struct {
	model.Params
	// 请求对象
	_topAssetDataAuthReqDto *TopAssetDataAuthReqDto
}

// NewTmallnrtassetauthorizationdeleteRequest 初始化TmallnrtassetauthorizationdeleteAPIRequest对象
func NewTmallnrtassetauthorizationdeleteRequest() *TmallnrtassetauthorizationdeleteAPIRequest {
	return &TmallnrtassetauthorizationdeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallnrtassetauthorizationdeleteAPIRequest) GetApiMethodName() string {
	return "tmall.nrt.asset.authorization.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallnrtassetauthorizationdeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallnrtassetauthorizationdeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopAssetDataAuthReqDto is TopAssetDataAuthReqDto Setter
// 请求对象
func (r *TmallnrtassetauthorizationdeleteAPIRequest) SetTopAssetDataAuthReqDto(_topAssetDataAuthReqDto *TopAssetDataAuthReqDto) error {
	r._topAssetDataAuthReqDto = _topAssetDataAuthReqDto
	r.Set("top_asset_data_auth_req_dto", _topAssetDataAuthReqDto)
	return nil
}

// GetTopAssetDataAuthReqDto TopAssetDataAuthReqDto Getter
func (r TmallnrtassetauthorizationdeleteAPIRequest) GetTopAssetDataAuthReqDto() *TopAssetDataAuthReqDto {
	return r._topAssetDataAuthReqDto
}
