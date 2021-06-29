package larkiot

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据影城id工作站和macId获取工作站 APIRequest
taobao.lark.pos.basedata.getworkstation

获取单独工作站
*/
type TaobaoLarkPosBasedataGetworkstationRequest struct {
    model.Params

    // 影城cinemaLinkId
    cinemaLinkId   string 

    // 终端编码
    posCode   string 

}

func NewTaobaoLarkPosBasedataGetworkstationRequest() *TaobaoLarkPosBasedataGetworkstationRequest{
    return &TaobaoLarkPosBasedataGetworkstationRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoLarkPosBasedataGetworkstationRequest) GetApiMethodName() string {
    return "taobao.lark.pos.basedata.getworkstation"
}

func (r TaobaoLarkPosBasedataGetworkstationRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoLarkPosBasedataGetworkstationRequest) SetCinemaLinkId(cinemaLinkId string) error {
    r.cinemaLinkId = cinemaLinkId
    r.Set("cinema_link_id", cinemaLinkId)
    return nil
}

func (r TaobaoLarkPosBasedataGetworkstationRequest) GetCinemaLinkId() string {
    return r.cinemaLinkId
}

func (r *TaobaoLarkPosBasedataGetworkstationRequest) SetPosCode(posCode string) error {
    r.posCode = posCode
    r.Set("pos_code", posCode)
    return nil
}

func (r TaobaoLarkPosBasedataGetworkstationRequest) GetPosCode() string {
    return r.posCode
}

