package dt

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAdsDataQueryAPIRequest 导入数据查询 API请求
// taobao.ads.data.query
//
// 导入数据查询
type TaobaoAdsDataQueryAPIRequest struct {
	model.Params
	// 查询参数
	_param0 *TaskDataImportSeqQry
}

// NewTaobaoAdsDataQueryRequest 初始化TaobaoAdsDataQueryAPIRequest对象
func NewTaobaoAdsDataQueryRequest() *TaobaoAdsDataQueryAPIRequest {
	return &TaobaoAdsDataQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAdsDataQueryAPIRequest) GetApiMethodName() string {
	return "taobao.ads.data.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAdsDataQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParam0 is Param0 Setter
// 查询参数
func (r *TaobaoAdsDataQueryAPIRequest) SetParam0(_param0 *TaskDataImportSeqQry) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r TaobaoAdsDataQueryAPIRequest) GetParam0() *TaskDataImportSeqQry {
	return r._param0
}
