package lstbm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
导入品牌商自有门店 APIRequest
alibaba.lst.bm.store.add

导入品牌商自有门店
*/
type AlibabaLstBmStoreAddRequest struct {
    model.Params

    // 门店数据模型
    openStoreDto   *LstTopOpenStoreDto 

}

func NewAlibabaLstBmStoreAddRequest() *AlibabaLstBmStoreAddRequest{
    return &AlibabaLstBmStoreAddRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaLstBmStoreAddRequest) GetApiMethodName() string {
    return "alibaba.lst.bm.store.add"
}

func (r AlibabaLstBmStoreAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaLstBmStoreAddRequest) SetOpenStoreDto(openStoreDto *LstTopOpenStoreDto) error {
    r.openStoreDto = openStoreDto
    r.Set("open_store_dto", openStoreDto)
    return nil
}

func (r AlibabaLstBmStoreAddRequest) GetOpenStoreDto() *LstTopOpenStoreDto {
    return r.openStoreDto
}

