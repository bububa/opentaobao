package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBaikeImportZhubaoDataAPIRequest 导入数据到商品百科服务 API请求
// taobao.baike.import.zhubao.data
//
// 用于接入外部数据录入到商品百科中
type TaobaoBaikeImportZhubaoDataAPIRequest struct {
	model.Params
	// 约定的Json数据
	_dataJsonStr string
}

// NewTaobaoBaikeImportZhubaoDataRequest 初始化TaobaoBaikeImportZhubaoDataAPIRequest对象
func NewTaobaoBaikeImportZhubaoDataRequest() *TaobaoBaikeImportZhubaoDataAPIRequest {
	return &TaobaoBaikeImportZhubaoDataAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoBaikeImportZhubaoDataAPIRequest) GetApiMethodName() string {
	return "taobao.baike.import.zhubao.data"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoBaikeImportZhubaoDataAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoBaikeImportZhubaoDataAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDataJsonStr is DataJsonStr Setter
// 约定的Json数据
func (r *TaobaoBaikeImportZhubaoDataAPIRequest) SetDataJsonStr(_dataJsonStr string) error {
	r._dataJsonStr = _dataJsonStr
	r.Set("data_json_str", _dataJsonStr)
	return nil
}

// GetDataJsonStr DataJsonStr Getter
func (r TaobaoBaikeImportZhubaoDataAPIRequest) GetDataJsonStr() string {
	return r._dataJsonStr
}
