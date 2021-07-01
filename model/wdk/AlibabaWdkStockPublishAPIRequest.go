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
type AlibabaWdkStockPublishAPIRequest struct {
    model.Params
    // 批量参数
    _batchStockPublishDto   *BatchStockPublishDto
}

// 初始化AlibabaWdkStockPublishAPIRequest对象
func NewAlibabaWdkStockPublishRequest() *AlibabaWdkStockPublishAPIRequest{
    return &AlibabaWdkStockPublishAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkStockPublishAPIRequest) GetApiMethodName() string {
    return "alibaba.wdk.stock.publish"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkStockPublishAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BatchStockPublishDto Setter
// 批量参数
func (r *AlibabaWdkStockPublishAPIRequest) SetBatchStockPublishDto(_batchStockPublishDto *BatchStockPublishDto) error {
    r._batchStockPublishDto = _batchStockPublishDto
    r.Set("batch_stock_publish_dto", _batchStockPublishDto)
    return nil
}

// BatchStockPublishDto Getter
func (r AlibabaWdkStockPublishAPIRequest) GetBatchStockPublishDto() *BatchStockPublishDto {
    return r._batchStockPublishDto
}
