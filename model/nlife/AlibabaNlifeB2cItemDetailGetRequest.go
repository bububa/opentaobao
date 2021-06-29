package nlife

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
b2c码详情查询 API请求
alibaba.nlife.b2c.item.detail.get

根据零售+平台生成的唯一码获取对应详情
*/
type AlibabaNlifeB2cItemDetailGetRequest struct {
    model.Params
    // 商家入驻门店在零售+平台的ID
    storeId   string
    // 零售+平台生成的唯一码或条码
    uniqueCode   string
}

// 初始化AlibabaNlifeB2cItemDetailGetRequest对象
func NewAlibabaNlifeB2cItemDetailGetRequest() *AlibabaNlifeB2cItemDetailGetRequest{
    return &AlibabaNlifeB2cItemDetailGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaNlifeB2cItemDetailGetRequest) GetApiMethodName() string {
    return "alibaba.nlife.b2c.item.detail.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaNlifeB2cItemDetailGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StoreId Setter
// 商家入驻门店在零售+平台的ID
func (r *AlibabaNlifeB2cItemDetailGetRequest) SetStoreId(storeId string) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

// StoreId Getter
func (r AlibabaNlifeB2cItemDetailGetRequest) GetStoreId() string {
    return r.storeId
}
// UniqueCode Setter
// 零售+平台生成的唯一码或条码
func (r *AlibabaNlifeB2cItemDetailGetRequest) SetUniqueCode(uniqueCode string) error {
    r.uniqueCode = uniqueCode
    r.Set("unique_code", uniqueCode)
    return nil
}

// UniqueCode Getter
func (r AlibabaNlifeB2cItemDetailGetRequest) GetUniqueCode() string {
    return r.uniqueCode
}
