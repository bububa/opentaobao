package lstbm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
修改品牌商自有门店数据 APIRequest
alibaba.lst.bm.store.update

修改品牌商自有门店数据
*/
type AlibabaLstBmStoreUpdateRequest struct {
    model.Params

    // 门店数据模型
    openStoreDto   *LstTopOpenStoreDto 

}

func NewAlibabaLstBmStoreUpdateRequest() *AlibabaLstBmStoreUpdateRequest{
    return &AlibabaLstBmStoreUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaLstBmStoreUpdateRequest) GetApiMethodName() string {
    return "alibaba.lst.bm.store.update"
}

func (r AlibabaLstBmStoreUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaLstBmStoreUpdateRequest) SetOpenStoreDto(openStoreDto *LstTopOpenStoreDto) error {
    r.openStoreDto = openStoreDto
    r.Set("open_store_dto", openStoreDto)
    return nil
}

func (r AlibabaLstBmStoreUpdateRequest) GetOpenStoreDto() *LstTopOpenStoreDto {
    return r.openStoreDto
}

