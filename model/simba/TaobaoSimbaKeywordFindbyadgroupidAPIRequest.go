package simba

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaKeywordFindbyadgroupidAPIRequest) GetApiMethodName() string {
	return "taobao.simba.keyword.findbyadgroupid"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaKeywordFindbyadgroupidAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
