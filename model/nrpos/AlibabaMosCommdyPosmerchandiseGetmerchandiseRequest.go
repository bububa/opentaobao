package nrpos

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
去前置机商品在线查询 API请求
alibaba.mos.commdy.posmerchandise.getmerchandise

去前置机商品在线查询接口
*/
type AlibabaMosCommdyPosmerchandiseGetmerchandiseAPIRequest struct {
    model.Params
    // 查询参数列表
    _posMerchandiseList   []QueryMerchandiseDTO
}

// 初始化AlibabaMosCommdyPosmerchandiseGetmerchandiseAPIRequest对象
func NewAlibabaMosCommdyPosmerchandiseGetmerchandiseRequest() *AlibabaMosCommdyPosmerchandiseGetmerchandiseAPIRequest{
    return &AlibabaMosCommdyPosmerchandiseGetmerchandiseAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaMosCommdyPosmerchandiseGetmerchandiseAPIRequest) GetApiMethodName() string {
    return "alibaba.mos.commdy.posmerchandise.getmerchandise"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaMosCommdyPosmerchandiseGetmerchandiseAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PosMerchandiseList Setter
// 查询参数列表
func (r *AlibabaMosCommdyPosmerchandiseGetmerchandiseAPIRequest) SetPosMerchandiseList(_posMerchandiseList []QueryMerchandiseDTO) error {
    r._posMerchandiseList = _posMerchandiseList
    r.Set("pos_merchandise_list", _posMerchandiseList)
    return nil
}

// PosMerchandiseList Getter
func (r AlibabaMosCommdyPosmerchandiseGetmerchandiseAPIRequest) GetPosMerchandiseList() []QueryMerchandiseDTO {
    return r._posMerchandiseList
}
