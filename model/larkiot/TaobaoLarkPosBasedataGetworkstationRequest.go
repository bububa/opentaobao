package larkiot

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据影城id工作站和macId获取工作站 API请求
taobao.lark.pos.basedata.getworkstation

获取单独工作站
*/
type TaobaoLarkPosBasedataGetworkstationRequest struct {
    model.Params
    // 影城cinemaLinkId
    _cinemaLinkId   string
    // 终端编码
    _posCode   string
}

// 初始化TaobaoLarkPosBasedataGetworkstationRequest对象
func NewTaobaoLarkPosBasedataGetworkstationRequest() *TaobaoLarkPosBasedataGetworkstationRequest{
    return &TaobaoLarkPosBasedataGetworkstationRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoLarkPosBasedataGetworkstationRequest) GetApiMethodName() string {
    return "taobao.lark.pos.basedata.getworkstation"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoLarkPosBasedataGetworkstationRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CinemaLinkId Setter
// 影城cinemaLinkId
func (r *TaobaoLarkPosBasedataGetworkstationRequest) SetCinemaLinkId(_cinemaLinkId string) error {
    r._cinemaLinkId = _cinemaLinkId
    r.Set("cinema_link_id", _cinemaLinkId)
    return nil
}

// CinemaLinkId Getter
func (r TaobaoLarkPosBasedataGetworkstationRequest) GetCinemaLinkId() string {
    return r._cinemaLinkId
}
// PosCode Setter
// 终端编码
func (r *TaobaoLarkPosBasedataGetworkstationRequest) SetPosCode(_posCode string) error {
    r._posCode = _posCode
    r.Set("pos_code", _posCode)
    return nil
}

// PosCode Getter
func (r TaobaoLarkPosBasedataGetworkstationRequest) GetPosCode() string {
    return r._posCode
}
