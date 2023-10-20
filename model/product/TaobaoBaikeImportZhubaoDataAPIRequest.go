package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaobaikeimportzhubaodataAPIRequest 导入数据到商品百科服务 API请求
// taobao.baike.import.zhubao.data
//
// 用于接入外部数据录入到商品百科中
type TaobaobaikeimportzhubaodataAPIRequest struct {
	model.Params
	// 约定的Json数据
	_dataJsonStr string
}

// NewTaobaobaikeimportzhubaodataRequest 初始化TaobaobaikeimportzhubaodataAPIRequest对象
func NewTaobaobaikeimportzhubaodataRequest() *TaobaobaikeimportzhubaodataAPIRequest {
	return &TaobaobaikeimportzhubaodataAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaobaikeimportzhubaodataAPIRequest) GetApiMethodName() string {
	return "taobao.baike.import.zhubao.data"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaobaikeimportzhubaodataAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaobaikeimportzhubaodataAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDataJsonStr is DataJsonStr Setter
// 约定的Json数据
func (r *TaobaobaikeimportzhubaodataAPIRequest) SetDataJsonStr(_dataJsonStr string) error {
	r._dataJsonStr = _dataJsonStr
	r.Set("data_json_str", _dataJsonStr)
	return nil
}

// GetDataJsonStr DataJsonStr Getter
func (r TaobaobaikeimportzhubaodataAPIRequest) GetDataJsonStr() string {
	return r._dataJsonStr
}
