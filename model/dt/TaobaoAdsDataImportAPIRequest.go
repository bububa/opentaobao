package dt

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAdsDataImportAPIRequest) Reset() {
	r._param0 = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAdsDataImportAPIRequest) GetApiMethodName() string {
	return "taobao.ads.data.import"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAdsDataImportAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAdsDataImportAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolTaobaoAdsDataImportAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAdsDataImportRequest()
	},
}

// GetTaobaoAdsDataImportRequest 从 sync.Pool 获取 TaobaoAdsDataImportAPIRequest
func GetTaobaoAdsDataImportAPIRequest() *TaobaoAdsDataImportAPIRequest {
	return poolTaobaoAdsDataImportAPIRequest.Get().(*TaobaoAdsDataImportAPIRequest)
}

// ReleaseTaobaoAdsDataImportAPIRequest 将 TaobaoAdsDataImportAPIRequest 放入 sync.Pool
func ReleaseTaobaoAdsDataImportAPIRequest(v *TaobaoAdsDataImportAPIRequest) {
	v.Reset()
	poolTaobaoAdsDataImportAPIRequest.Put(v)
}
