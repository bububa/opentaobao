package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingOpenDataRelationQueryAPIRequest 数据关联关系查询 API请求
// alibaba.wdk.marketing.open.data.relation.query
//
// 数据关联关系查询
type AlibabaWdkMarketingOpenDataRelationQueryAPIRequest struct {
	model.Params
	// 外部数据Id
	_outDataIds []string
	// 数据类型：WDK_MARKET:五道口营销
	_bizCode string
	// 数据子类型：ACTIVITY:营销活动数据
	_subBizCode string
}

// NewAlibabaWdkMarketingOpenDataRelationQueryRequest 初始化AlibabaWdkMarketingOpenDataRelationQueryAPIRequest对象
func NewAlibabaWdkMarketingOpenDataRelationQueryRequest() *AlibabaWdkMarketingOpenDataRelationQueryAPIRequest {
	return &AlibabaWdkMarketingOpenDataRelationQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingOpenDataRelationQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.open.data.relation.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingOpenDataRelationQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkMarketingOpenDataRelationQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOutDataIds is OutDataIds Setter
// 外部数据Id
func (r *AlibabaWdkMarketingOpenDataRelationQueryAPIRequest) SetOutDataIds(_outDataIds []string) error {
	r._outDataIds = _outDataIds
	r.Set("out_data_ids", _outDataIds)
	return nil
}

// GetOutDataIds OutDataIds Getter
func (r AlibabaWdkMarketingOpenDataRelationQueryAPIRequest) GetOutDataIds() []string {
	return r._outDataIds
}

// SetBizCode is BizCode Setter
// 数据类型：WDK_MARKET:五道口营销
func (r *AlibabaWdkMarketingOpenDataRelationQueryAPIRequest) SetBizCode(_bizCode string) error {
	r._bizCode = _bizCode
	r.Set("biz_code", _bizCode)
	return nil
}

// GetBizCode BizCode Getter
func (r AlibabaWdkMarketingOpenDataRelationQueryAPIRequest) GetBizCode() string {
	return r._bizCode
}

// SetSubBizCode is SubBizCode Setter
// 数据子类型：ACTIVITY:营销活动数据
func (r *AlibabaWdkMarketingOpenDataRelationQueryAPIRequest) SetSubBizCode(_subBizCode string) error {
	r._subBizCode = _subBizCode
	r.Set("sub_biz_code", _subBizCode)
	return nil
}

// GetSubBizCode SubBizCode Getter
func (r AlibabaWdkMarketingOpenDataRelationQueryAPIRequest) GetSubBizCode() string {
	return r._subBizCode
}
