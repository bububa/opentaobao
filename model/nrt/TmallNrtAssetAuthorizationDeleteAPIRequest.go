package nrt

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallNrtAssetAuthorizationDeleteAPIRequest 移除资产数据权限授权关系 API请求
// tmall.nrt.asset.authorization.delete
//
// 移除资产数据权限授权关系
type TmallNrtAssetAuthorizationDeleteAPIRequest struct {
	model.Params
	// 请求对象
	_topAssetDataAuthReqDto *TopAssetDataAuthReqDto
}

// NewTmallNrtAssetAuthorizationDeleteRequest 初始化TmallNrtAssetAuthorizationDeleteAPIRequest对象
func NewTmallNrtAssetAuthorizationDeleteRequest() *TmallNrtAssetAuthorizationDeleteAPIRequest {
	return &TmallNrtAssetAuthorizationDeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallNrtAssetAuthorizationDeleteAPIRequest) GetApiMethodName() string {
	return "tmall.nrt.asset.authorization.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallNrtAssetAuthorizationDeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallNrtAssetAuthorizationDeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopAssetDataAuthReqDto is TopAssetDataAuthReqDto Setter
// 请求对象
func (r *TmallNrtAssetAuthorizationDeleteAPIRequest) SetTopAssetDataAuthReqDto(_topAssetDataAuthReqDto *TopAssetDataAuthReqDto) error {
	r._topAssetDataAuthReqDto = _topAssetDataAuthReqDto
	r.Set("top_asset_data_auth_req_dto", _topAssetDataAuthReqDto)
	return nil
}

// GetTopAssetDataAuthReqDto TopAssetDataAuthReqDto Getter
func (r TmallNrtAssetAuthorizationDeleteAPIRequest) GetTopAssetDataAuthReqDto() *TopAssetDataAuthReqDto {
	return r._topAssetDataAuthReqDto
}
