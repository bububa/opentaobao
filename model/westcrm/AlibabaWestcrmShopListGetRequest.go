package westcrm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取商圈商户信息列表 API请求
alibaba.westcrm.shop.list.get

获取商圈商户信息列表
*/
type AlibabaWestcrmShopListGetAPIRequest struct {
    model.Params
    // 园区id
    _campusId   int64
}

// 初始化AlibabaWestcrmShopListGetAPIRequest对象
func NewAlibabaWestcrmShopListGetRequest() *AlibabaWestcrmShopListGetAPIRequest{
    return &AlibabaWestcrmShopListGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWestcrmShopListGetAPIRequest) GetApiMethodName() string {
    return "alibaba.westcrm.shop.list.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWestcrmShopListGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CampusId Setter
// 园区id
func (r *AlibabaWestcrmShopListGetAPIRequest) SetCampusId(_campusId int64) error {
    r._campusId = _campusId
    r.Set("campus_id", _campusId)
    return nil
}

// CampusId Getter
func (r AlibabaWestcrmShopListGetAPIRequest) GetCampusId() int64 {
    return r._campusId
}
