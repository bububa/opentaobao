package nrpos

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
去前置机商品在线查询 APIRequest
alibaba.mos.commdy.posmerchandise.getmerchandise

去前置机商品在线查询接口
*/
type AlibabaMosCommdyPosmerchandiseGetmerchandiseRequest struct {
    model.Params

    // 查询参数列表
    posMerchandiseList   []QueryMerchandiseDto 

}

func NewAlibabaMosCommdyPosmerchandiseGetmerchandiseRequest() *AlibabaMosCommdyPosmerchandiseGetmerchandiseRequest{
    return &AlibabaMosCommdyPosmerchandiseGetmerchandiseRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaMosCommdyPosmerchandiseGetmerchandiseRequest) GetApiMethodName() string {
    return "alibaba.mos.commdy.posmerchandise.getmerchandise"
}

func (r AlibabaMosCommdyPosmerchandiseGetmerchandiseRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaMosCommdyPosmerchandiseGetmerchandiseRequest) SetPosMerchandiseList(posMerchandiseList []QueryMerchandiseDto) error {
    r.posMerchandiseList = posMerchandiseList
    r.Set("pos_merchandise_list", posMerchandiseList)
    return nil
}

func (r AlibabaMosCommdyPosmerchandiseGetmerchandiseRequest) GetPosMerchandiseList() []QueryMerchandiseDto {
    return r.posMerchandiseList
}

