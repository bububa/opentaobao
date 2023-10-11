package nrt

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallNrtAssetAuthorizationAddAPIRequest 增加数据权限授权 API请求
// tmall.nrt.asset.authorization.add
//
// 增加数据权限授权
type TmallNrtAssetAuthorizationAddAPIRequest struct {
	model.Params
	// 添加数据权限请求
	_topAssetDataAuthReqDto *TopAssetDataAuthReqDto
}

// NewTmallNrtAssetAuthorizationAddRequest 初始化TmallNrtAssetAuthorizationAddAPIRequest对象
func NewTmallNrtAssetAuthorizationAddRequest() *TmallNrtAssetAuthorizationAddAPIRequest {
	return &TmallNrtAssetAuthorizationAddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallNrtAssetAuthorizationAddAPIRequest) GetApiMethodName() string {
	return "tmall.nrt.asset.authorization.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallNrtAssetAuthorizationAddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallNrtAssetAuthorizationAddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopAssetDataAuthReqDto is TopAssetDataAuthReqDto Setter
// 添加数据权限请求
func (r *TmallNrtAssetAuthorizationAddAPIRequest) SetTopAssetDataAuthReqDto(_topAssetDataAuthReqDto *TopAssetDataAuthReqDto) error {
	r._topAssetDataAuthReqDto = _topAssetDataAuthReqDto
	r.Set("top_asset_data_auth_req_dto", _topAssetDataAuthReqDto)
	return nil
}

// GetTopAssetDataAuthReqDto TopAssetDataAuthReqDto Getter
func (r TmallNrtAssetAuthorizationAddAPIRequest) GetTopAssetDataAuthReqDto() *TopAssetDataAuthReqDto {
	return r._topAssetDataAuthReqDto
}
