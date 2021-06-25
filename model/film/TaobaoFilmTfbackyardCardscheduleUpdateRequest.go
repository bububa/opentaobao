package film

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
CGV影城卡排期数据传输 APIRequest
taobao.film.tfbackyard.cardschedule.update

cgv影城卡排期价格数据传输API
*/
type TaobaoFilmTfbackyardCardscheduleUpdateRequest struct {
    model.Params

    // CGV影城卡价格数据
    jsonData   string 

}

func NewTaobaoFilmTfbackyardCardscheduleUpdateRequest() *TaobaoFilmTfbackyardCardscheduleUpdateRequest{
    return &TaobaoFilmTfbackyardCardscheduleUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoFilmTfbackyardCardscheduleUpdateRequest) GetApiMethodName() string {
    return "taobao.film.tfbackyard.cardschedule.update"
}

func (r TaobaoFilmTfbackyardCardscheduleUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoFilmTfbackyardCardscheduleUpdateRequest) SetJsonData(jsonData string) error {
    r.jsonData = jsonData
    r.Set("json_data", jsonData)
    return nil
}

func (r TaobaoFilmTfbackyardCardscheduleUpdateRequest) GetJsonData() string {
    return r.jsonData
}

