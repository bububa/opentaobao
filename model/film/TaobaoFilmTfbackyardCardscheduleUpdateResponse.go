package film

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
CGV影城卡排期数据传输 API返回值 
taobao.film.tfbackyard.cardschedule.update

cgv影城卡排期价格数据传输API
*/
type TaobaoFilmTfbackyardCardscheduleUpdateAPIResponse struct {
    model.CommonResponse
    TaobaoFilmTfbackyardCardscheduleUpdateResponse
}

// CGV影城卡排期数据传输 成功返回结果
type TaobaoFilmTfbackyardCardscheduleUpdateResponse struct {
    XMLName xml.Name `xml:"film_tfbackyard_cardschedule_update_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   string `json:"result,omitempty" xml:"result,omitempty"`
}
