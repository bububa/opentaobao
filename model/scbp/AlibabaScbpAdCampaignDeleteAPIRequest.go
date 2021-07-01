package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaScbpAdCampaignDeleteAPIRequest
删除计划 API请求
alibaba.scbp.ad.campaign.delete

删除计划 */
type AlibabaScbpAdCampaignDeleteAPIRequest struct {
	model.Params
	// 操作对象
	_batchOperation *CampaignBatchOperationDto
	// 用户信息
	_topContext *TopContextDto
}

// NewAlibabaScbpAdCampaignDeleteRequest 初始化AlibabaScbpAdCampaignDeleteAPIRequest对象
func NewAlibabaScbpAdCampaignDeleteRequest() *AlibabaScbpAdCampaignDeleteAPIRequest {
	return &AlibabaScbpAdCampaignDeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdCampaignDeleteAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.campaign.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdCampaignDeleteAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is BatchOperation Setter
// 操作对象
func (r *AlibabaScbpAdCampaignDeleteAPIRequest) SetBatchOperation(_batchOperation *CampaignBatchOperationDto) error {
	r._batchOperation = _batchOperation
	r.Set("batch_operation", _batchOperation)
	return nil
}

// Get BatchOperation Getter
func (r AlibabaScbpAdCampaignDeleteAPIRequest) GetBatchOperation() *CampaignBatchOperationDto {
	return r._batchOperation
}

// Set is TopContext Setter
// 用户信息
func (r *AlibabaScbpAdCampaignDeleteAPIRequest) SetTopContext(_topContext *TopContextDto) error {
	r._topContext = _topContext
	r.Set("top_context", _topContext)
	return nil
}

// Get TopContext Getter
func (r AlibabaScbpAdCampaignDeleteAPIRequest) GetTopContext() *TopContextDto {
	return r._topContext
}
