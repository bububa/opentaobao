package moscm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批量录入商品信息 APIRequest
alibaba.mos.goods.bulkinputcspu

用于商品信息的批量导入到银泰商品中台
*/
type AlibabaMosGoodsBulkinputcspuRequest struct {
    model.Params

    // 录入商品信息列表（最大列表长度：20）
    cspuInputDtoList   []CspuInputDto 

}

func NewAlibabaMosGoodsBulkinputcspuRequest() *AlibabaMosGoodsBulkinputcspuRequest{
    return &AlibabaMosGoodsBulkinputcspuRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaMosGoodsBulkinputcspuRequest) GetApiMethodName() string {
    return "alibaba.mos.goods.bulkinputcspu"
}

func (r AlibabaMosGoodsBulkinputcspuRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaMosGoodsBulkinputcspuRequest) SetCspuInputDtoList(cspuInputDtoList []CspuInputDto) error {
    r.cspuInputDtoList = cspuInputDtoList
    r.Set("cspu_input_dto_list", cspuInputDtoList)
    return nil
}

func (r AlibabaMosGoodsBulkinputcspuRequest) GetCspuInputDtoList() []CspuInputDto {
    return r.cspuInputDtoList
}

