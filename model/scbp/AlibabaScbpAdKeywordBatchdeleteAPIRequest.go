package scbp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdKeywordBatchdeleteAPIRequest 外贸直通车批量删除关键词 API请求
// alibaba.scbp.ad.keyword.batchdelete
//
// 外贸直通车批量删除关键词
type AlibabaScbpAdKeywordBatchdeleteAPIRequest struct {
	model.Params
	// 关键词Id列表
	_keywordIdList []string
}

// NewAlibabaScbpAdKeywordBatchdeleteRequest 初始化AlibabaScbpAdKeywordBatchdeleteAPIRequest对象
func NewAlibabaScbpAdKeywordBatchdeleteRequest() *AlibabaScbpAdKeywordBatchdeleteAPIRequest {
	return &AlibabaScbpAdKeywordBatchdeleteAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaScbpAdKeywordBatchdeleteAPIRequest) Reset() {
	r._keywordIdList = r._keywordIdList[:0]
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdKeywordBatchdeleteAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.keyword.batchdelete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdKeywordBatchdeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaScbpAdKeywordBatchdeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetKeywordIdList is KeywordIdList Setter
// 关键词Id列表
func (r *AlibabaScbpAdKeywordBatchdeleteAPIRequest) SetKeywordIdList(_keywordIdList []string) error {
	r._keywordIdList = _keywordIdList
	r.Set("keyword_id_list", _keywordIdList)
	return nil
}

// GetKeywordIdList KeywordIdList Getter
func (r AlibabaScbpAdKeywordBatchdeleteAPIRequest) GetKeywordIdList() []string {
	return r._keywordIdList
}

var poolAlibabaScbpAdKeywordBatchdeleteAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaScbpAdKeywordBatchdeleteRequest()
	},
}

// GetAlibabaScbpAdKeywordBatchdeleteRequest 从 sync.Pool 获取 AlibabaScbpAdKeywordBatchdeleteAPIRequest
func GetAlibabaScbpAdKeywordBatchdeleteAPIRequest() *AlibabaScbpAdKeywordBatchdeleteAPIRequest {
	return poolAlibabaScbpAdKeywordBatchdeleteAPIRequest.Get().(*AlibabaScbpAdKeywordBatchdeleteAPIRequest)
}

// ReleaseAlibabaScbpAdKeywordBatchdeleteAPIRequest 将 AlibabaScbpAdKeywordBatchdeleteAPIRequest 放入 sync.Pool
func ReleaseAlibabaScbpAdKeywordBatchdeleteAPIRequest(v *AlibabaScbpAdKeywordBatchdeleteAPIRequest) {
	v.Reset()
	poolAlibabaScbpAdKeywordBatchdeleteAPIRequest.Put(v)
}
