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
type AlibabaMosCommdyPosmerchandiseGetmerchandiseRequest struct {
    model.Params
    // 查询参数列表
    _posMerchandiseList   []QueryMerchandiseDto
}

// 初始化AlibabaMosCommdyPosmerchandiseGetmerchandiseRequest对象
func NewAlibabaMosCommdyPosmerchandiseGetmerchandiseRequest() *AlibabaMosCommdyPosmerchandiseGetmerchandiseRequest{
    return &AlibabaMosCommdyPosmerchandiseGetmerchandiseRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaMosCommdyPosmerchandiseGetmerchandiseRequest) GetApiMethodName() string {
    return "alibaba.mos.commdy.posmerchandise.getmerchandise"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaMosCommdyPosmerchandiseGetmerchandiseRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PosMerchandiseList Setter
// 查询参数列表
func (r *AlibabaMosCommdyPosmerchandiseGetmerchandiseRequest) SetPosMerchandiseList(_posMerchandiseList []QueryMerchandiseDto) error {
    r._posMerchandiseList = _posMerchandiseList
    r.Set("pos_merchandise_list", _posMerchandiseList)
    return nil
}

// PosMerchandiseList Getter
func (r AlibabaMosCommdyPosmerchandiseGetmerchandiseRequest) GetPosMerchandiseList() []QueryMerchandiseDto {
    return r._posMerchandiseList
}
