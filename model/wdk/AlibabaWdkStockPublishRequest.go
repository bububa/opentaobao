package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
五道口库存发布接口（针对外部渠道） API请求
alibaba.wdk.stock.publish

五道口库存发布接口（针对外部渠道）
*/
type AlibabaWdkStockPublishRequest struct {
    model.Params
    // 批量参数
    _batchStockPublishDto   *BatchStockPublishDTO
}

// 初始化AlibabaWdkStockPublishRequest对象
func NewAlibabaWdkStockPublishRequest() *AlibabaWdkStockPublishRequest{
    return &AlibabaWdkStockPublishRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkStockPublishRequest) GetApiMethodName() string {
    return "alibaba.wdk.stock.publish"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkStockPublishRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BatchStockPublishDto Setter
// 批量参数
func (r *AlibabaWdkStockPublishRequest) SetBatchStockPublishDto(_batchStockPublishDto *BatchStockPublishDTO) error {
    r._batchStockPublishDto = _batchStockPublishDto
    r.Set("batch_stock_publish_dto", _batchStockPublishDto)
    return nil
}

// BatchStockPublishDto Getter
func (r AlibabaWdkStockPublishRequest) GetBatchStockPublishDto() *BatchStockPublishDTO {
    return r._batchStockPublishDto
}
