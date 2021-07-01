package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaScbpProductGroupGetAPIRequest
查询指定产品分组的下一层子分组 API请求
alibaba.scbp.product.group.get

查询指定产品分组的下一层子分组 */
type AlibabaScbpProductGroupGetAPIRequest struct {
	model.Params
	// 产品分组标识，null表示查询第一层分组
	_groupId string
}

// NewAlibabaScbpProductGroupGetRequest 初始化AlibabaScbpProductGroupGetAPIRequest对象
func NewAlibabaScbpProductGroupGetRequest() *AlibabaScbpProductGroupGetAPIRequest {
	return &AlibabaScbpProductGroupGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpProductGroupGetAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.product.group.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpProductGroupGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is GroupId Setter
// 产品分组标识，null表示查询第一层分组
func (r *AlibabaScbpProductGroupGetAPIRequest) SetGroupId(_groupId string) error {
	r._groupId = _groupId
	r.Set("group_id", _groupId)
	return nil
}

// Get GroupId Getter
func (r AlibabaScbpProductGroupGetAPIRequest) GetGroupId() string {
	return r._groupId
}
