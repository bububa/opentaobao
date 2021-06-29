package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
导入数据到商品百科服务 API请求
taobao.baike.import.zhubao.data

用于接入外部数据录入到商品百科中
*/
type TaobaoBaikeImportZhubaoDataRequest struct {
    model.Params
    // 约定的Json数据
    _dataJsonStr   string
}

// 初始化TaobaoBaikeImportZhubaoDataRequest对象
func NewTaobaoBaikeImportZhubaoDataRequest() *TaobaoBaikeImportZhubaoDataRequest{
    return &TaobaoBaikeImportZhubaoDataRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoBaikeImportZhubaoDataRequest) GetApiMethodName() string {
    return "taobao.baike.import.zhubao.data"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoBaikeImportZhubaoDataRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DataJsonStr Setter
// 约定的Json数据
func (r *TaobaoBaikeImportZhubaoDataRequest) SetDataJsonStr(_dataJsonStr string) error {
    r._dataJsonStr = _dataJsonStr
    r.Set("data_json_str", _dataJsonStr)
    return nil
}

// DataJsonStr Getter
func (r TaobaoBaikeImportZhubaoDataRequest) GetDataJsonStr() string {
    return r._dataJsonStr
}
