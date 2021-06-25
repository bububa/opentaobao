package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
导入数据到商品百科服务 APIRequest
taobao.baike.import.zhubao.data

用于接入外部数据录入到商品百科中
*/
type TaobaoBaikeImportZhubaoDataRequest struct {
    model.Params

    // 约定的Json数据
    dataJsonStr   string 

}

func NewTaobaoBaikeImportZhubaoDataRequest() *TaobaoBaikeImportZhubaoDataRequest{
    return &TaobaoBaikeImportZhubaoDataRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoBaikeImportZhubaoDataRequest) GetApiMethodName() string {
    return "taobao.baike.import.zhubao.data"
}

func (r TaobaoBaikeImportZhubaoDataRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoBaikeImportZhubaoDataRequest) SetDataJsonStr(dataJsonStr string) error {
    r.dataJsonStr = dataJsonStr
    r.Set("data_json_str", dataJsonStr)
    return nil
}

func (r TaobaoBaikeImportZhubaoDataRequest) GetDataJsonStr() string {
    return r.dataJsonStr
}

