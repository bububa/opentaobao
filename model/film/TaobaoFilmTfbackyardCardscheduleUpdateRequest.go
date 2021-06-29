package film

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
CGV影城卡排期数据传输 API请求
taobao.film.tfbackyard.cardschedule.update

cgv影城卡排期价格数据传输API
*/
type TaobaoFilmTfbackyardCardscheduleUpdateRequest struct {
    model.Params
    // CGV影城卡价格数据
    jsonData   string
}

// 初始化TaobaoFilmTfbackyardCardscheduleUpdateRequest对象
func NewTaobaoFilmTfbackyardCardscheduleUpdateRequest() *TaobaoFilmTfbackyardCardscheduleUpdateRequest{
    return &TaobaoFilmTfbackyardCardscheduleUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFilmTfbackyardCardscheduleUpdateRequest) GetApiMethodName() string {
    return "taobao.film.tfbackyard.cardschedule.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFilmTfbackyardCardscheduleUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// JsonData Setter
// CGV影城卡价格数据
func (r *TaobaoFilmTfbackyardCardscheduleUpdateRequest) SetJsonData(jsonData string) error {
    r.jsonData = jsonData
    r.Set("json_data", jsonData)
    return nil
}

// JsonData Getter
func (r TaobaoFilmTfbackyardCardscheduleUpdateRequest) GetJsonData() string {
    return r.jsonData
}
