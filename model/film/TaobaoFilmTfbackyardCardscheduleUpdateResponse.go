package film

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
CGV影城卡排期数据传输 APIResponse
taobao.film.tfbackyard.cardschedule.update

cgv影城卡排期价格数据传输API
*/
type TaobaoFilmTfbackyardCardscheduleUpdateAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"film_tfbackyard_cardschedule_update_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   string `json:"result,omitempty" xml:"