package simba

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaKeywordFindbyadgroupidAPIRequest 根据推广单元id获取关键词 API请求
// taobao.simba.keyword.findbyadgroupid
//
// 根据一个关键词Id列表取得一组关键词
type TaobaoSimbaKeywordFindbyadgroupidAPIRequest struct {
	model.Params
	// 推广单元id
	_adgroupId int64
}

// NewTaobaoSimbaKeywordFindbyadgroupidRequest 初始化TaobaoSimbaKeywordFindbyadgroupidAPIRequest对象
func NewTaobaoSimbaKeywordFindbyadgroupidRequest() *TaobaoSimbaKeywordFindbyadgroupidAPIRequest {
	return &TaobaoSimbaKeywordFindbyadgroupidAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSimbaKeywordFindbyadgroupidAPIRequest) Reset() {
	r._adgroupId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaKeywordFindbyadgroupidAPIRequest) GetApiMethodName() string {
	return "taobao.simba.keyword.findbyadgroupid"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaKeywordFindbyadgroupidAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSimbaKeywordFindbyadgroupidAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAdgroupId is AdgroupId Setter
// 推广单元id
func (r *TaobaoSimbaKeywordFindbyadgroupidAPIRequest) SetAdgroupId(_adgroupId int64) error {
	r._adgroupId = _adgroupId
	r.Set("adgroup_id", _adgroupId)
	return nil
}

// GetAdgroupId AdgroupId Getter
func (r TaobaoSimbaKeywordFindbyadgroupidAPIRequest) GetAdgroupId() int64 {
	return r._adgroupId
}

var poolTaobaoSimbaKeywordFindbyadgroupidAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSimbaKeywordFindbyadgroupidRequest()
	},
}

// GetTaobaoSimbaKeywordFindbyadgroupidRequest 从 sync.Pool 获取 TaobaoSimbaKeywordFindbyadgroupidAPIRequest
func GetTaobaoSimbaKeywordFindbyadgroupidAPIRequest() *TaobaoSimbaKeywordFindbyadgroupidAPIRequest {
	return poolTaobaoSimbaKeywordFindbyadgroupidAPIRequest.Get().(*TaobaoSimbaKeywordFindbyadgroupidAPIRequest)
}

// ReleaseTaobaoSimbaKeywordFindbyadgroupidAPIRequest 将 TaobaoSimbaKeywordFindbyadgroupidAPIRequest 放入 sync.Pool
func ReleaseTaobaoSimbaKeywordFindbyadgroupidAPIRequest(v *TaobaoSimbaKeywordFindbyadgroupidAPIRequest) {
	v.Reset()
	poolTaobaoSimbaKeywordFindbyadgroupidAPIRequest.Put(v)
}
