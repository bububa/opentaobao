package dt

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAdsDataImportAPIRequest 数据导入 API请求
// taobao.ads.data.import
//
// 数据导入
type TaobaoAdsDataImportAPIRequest struct {
	model.Params
	// 请求入参
	_param0 *RequesterDataJobSaveCmd
}

// NewTaobaoAdsDataImportRequest 初始化TaobaoAdsDataImportAPIRequest对象
func NewTaobaoAdsDataImportRequest() *TaobaoAdsDataImportAPIRequest {
	return &TaobaoAdsDataImportAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAdsDataImportAPIRequest) GetApiMethodName() string {
	return "taobao.ads.data.import"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAdsDataImportAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParam0 is Param0 Setter
// 请求入参
func (r *TaobaoAdsDataImportAPIRequest) SetParam0(_param0 *RequesterDataJobSaveCmd) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r TaobaoAdsDataImportAPIRequest) GetParam0() *RequesterDataJobSaveCmd {
	return r._param0
}
