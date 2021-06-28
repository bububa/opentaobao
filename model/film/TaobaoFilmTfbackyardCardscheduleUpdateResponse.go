package film

import (
    "github.com/bububa/opentaobao/model"
)

/* 
CGV影城卡排期数据传输 APIResponse
taobao.film.tfbackyard.cardschedule.update

cgv影城卡排期价格数据传输API
*/
type TaobaoFilmTfbackyardCardscheduleUpdateAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoFilmTfbackyardCardscheduleUpdateResponse `json:"film_tfbackyard_cardschedule_update_response,omitempty"` 
    TaobaoFilmTfbackyardCardscheduleUpdateResponse
}

/* model for simplify = false
type TaobaoFilmTfbackyardCardscheduleUpdateResponse struct {

    // result
    
    Result   string `json:"result,omitempty"`
    

}
*/

type TaobaoFilmTfbackyardCardscheduleUpdateResponse struct {

    // result
    Result   string `json:"result,omitempty"`

}
