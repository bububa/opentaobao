package lstbm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
修改品牌商自有门店数据 API请求
alibaba.lst.bm.store.update

修改品牌商自有门店数据
*/
type AlibabaLstBmStoreUpdateRequest struct {
    model.Params
    // 门店数据模型
    _openStoreDto   *LstTopOpenStoreDto
}

// 初始化AlibabaLstBmStoreUpdateRequest对象
func NewAlibabaLstBmStoreUpdateRequest() *AlibabaLstBmStoreUpdateRequest{
    return &AlibabaLstBmStoreUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaLstBmStoreUpdateRequest) GetApiMethodName() string {
    return "alibaba.lst.bm.store.update"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaLstBmStoreUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OpenStoreDto Setter
// 门店数据模型
func (r *AlibabaLstBmStoreUpdateRequest) SetOpenStoreDto(_openStoreDto *LstTopOpenStoreDto) error {
    r._openStoreDto = _openStoreDto
    r.Set("open_store_dto", _openStoreDto)
    return nil
}

// OpenStoreDto Getter
func (r AlibabaLstBmStoreUpdateRequest) GetOpenStoreDto() *LstTopOpenStoreDto {
    return r._openStoreDto
}
