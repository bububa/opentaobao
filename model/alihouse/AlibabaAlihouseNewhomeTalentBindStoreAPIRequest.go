package alihouse

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeTalentBindStoreAPIRequest 达人号门店关系绑定 API请求
// alibaba.alihouse.newhome.talent.bind.store
//
// 达人号门店关系绑定
type AlibabaAlihouseNewhomeTalentBindStoreAPIRequest struct {
	model.Params
	// aaa
	_storeTalentList *StoreTalentDto
}

// NewAlibabaAlihouseNewhomeTalentBindStoreRequest 初始化AlibabaAlihouseNewhomeTalentBindStoreAPIRequest对象
func NewAlibabaAlihouseNewhomeTalentBindStoreRequest() *AlibabaAlihouseNewhomeTalentBindStoreAPIRequest {
	return &AlibabaAlihouseNewhomeTalentBindStoreAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihouseNewhomeTalentBindStoreAPIRequest) Reset() {
	r._storeTalentList = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeTalentBindStoreAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.talent.bind.store"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeTalentBindStoreAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseNewhomeTalentBindStoreAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStoreTalentList is StoreTalentList Setter
// aaa
func (r *AlibabaAlihouseNewhomeTalentBindStoreAPIRequest) SetStoreTalentList(_storeTalentList *StoreTalentDto) error {
	r._storeTalentList = _storeTalentList
	r.Set("store_talent_list", _storeTalentList)
	return nil
}

// GetStoreTalentList StoreTalentList Getter
func (r AlibabaAlihouseNewhomeTalentBindStoreAPIRequest) GetStoreTalentList() *StoreTalentDto {
	return r._storeTalentList
}

var poolAlibabaAlihouseNewhomeTalentBindStoreAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihouseNewhomeTalentBindStoreRequest()
	},
}

// GetAlibabaAlihouseNewhomeTalentBindStoreRequest 从 sync.Pool 获取 AlibabaAlihouseNewhomeTalentBindStoreAPIRequest
func GetAlibabaAlihouseNewhomeTalentBindStoreAPIRequest() *AlibabaAlihouseNewhomeTalentBindStoreAPIRequest {
	return poolAlibabaAlihouseNewhomeTalentBindStoreAPIRequest.Get().(*AlibabaAlihouseNewhomeTalentBindStoreAPIRequest)
}

// ReleaseAlibabaAlihouseNewhomeTalentBindStoreAPIRequest 将 AlibabaAlihouseNewhomeTalentBindStoreAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihouseNewhomeTalentBindStoreAPIRequest(v *AlibabaAlihouseNewhomeTalentBindStoreAPIRequest) {
	v.Reset()
	poolAlibabaAlihouseNewhomeTalentBindStoreAPIRequest.Put(v)
}
