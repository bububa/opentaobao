package lstbm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
导入品牌商自有门店 API请求
alibaba.lst.bm.store.add

导入品牌商自有门店
*/
type AlibabaLstBmStoreAddRequest struct {
    model.Params
    // 门店数据模型
    _openStoreDto   *LstTopOpenStoreDTO
}

// 初始化AlibabaLstBmStoreAddRequest对象
func NewAlibabaLstBmStoreAddRequest() *AlibabaLstBmStoreAddRequest{
    return &AlibabaLstBmStoreAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaLstBmStoreAddRequest) GetApiMethodName() string {
    return "alibaba.lst.bm.store.add"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaLstBmStoreAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OpenStoreDto Setter
// 门店数据模型
func (r *AlibabaLstBmStoreAddRequest) SetOpenStoreDto(_openStoreDto *LstTopOpenStoreDTO) error {
    r._openStoreDto = _openStoreDto
    r.Set("open_store_dto", _openStoreDto)
    return nil
}

// OpenStoreDto Getter
func (r AlibabaLstBmStoreAddRequest) GetOpenStoreDto() *LstTopOpenStoreDTO {
    return r._openStoreDto
}
