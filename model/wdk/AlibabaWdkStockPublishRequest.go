package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
五道口库存发布接口（针对外部渠道） APIRequest
alibaba.wdk.stock.publish

五道口库存发布接口（针对外部渠道）
*/
type AlibabaWdkStockPublishRequest struct {
    model.Params

    // 批量参数
    batchStockPublishDto   *BatchStockPublishDto 

}

func NewAlibabaWdkStockPublishRequest() *AlibabaWdkStockPublishRequest{
    return &AlibabaWdkStockPublishRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkStockPublishRequest) GetApiMethodName() string {
    return "alibaba.wdk.stock.publish"
}

func (r AlibabaWdkStockPublishRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkStockPublishRequest) SetBatchStockPublishDto(batchStockPublishDto *BatchStockPublishDto) error {
    r.batchStockPublishDto = batchStockPublishDto
    r.Set("batch_stock_publish_dto", batchStockPublishDto)
    return nil
}

func (r AlibabaWdkStockPublishRequest) GetBatchStockPublishDto() *BatchStockPublishDto {
    return r.batchStockPublishDto
}

