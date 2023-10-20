package simba

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaSerchcrowdBatchDeleteAPIRequest 单品搜索人群批量取消溢价 API请求
// taobao.simba.serchcrowd.batch.delete
//
// 删除单品搜索人群溢价功能
type TaobaoSimbaSerchcrowdBatchDeleteAPIRequest struct {
	model.Params
	// 需要删除的人群id
	_adgroupCrowdIds []string
	// 子帐号nick
	_subNick string
	// 被操作者的淘宝昵称
	_nick string
}

// NewTaobaoSimbaSerchcrowdBatchDeleteRequest 初始化TaobaoSimbaSerchcrowdBatchDeleteAPIRequest对象
func NewTaobaoSimbaSerchcrowdBatchDeleteRequest() *TaobaoSimbaSerchcrowdBatchDeleteAPIRequest {
	return &TaobaoSimbaSerchcrowdBatchDeleteAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSimbaSerchcrowdBatchDeleteAPIRequest) Reset() {
	r._adgroupCrowdIds = r._adgroupCrowdIds[:0]
	r._subNick = ""
	r._nick = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaSerchcrowdBatchDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.simba.serchcrowd.batch.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaSerchcrowdBatchDeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSimbaSerchcrowdBatchDeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAdgroupCrowdIds is AdgroupCrowdIds Setter
// 需要删除的人群id
func (r *TaobaoSimbaSerchcrowdBatchDeleteAPIRequest) SetAdgroupCrowdIds(_adgroupCrowdIds []string) error {
	r._adgroupCrowdIds = _adgroupCrowdIds
	r.Set("adgroup_crowd_ids", _adgroupCrowdIds)
	return nil
}

// GetAdgroupCrowdIds AdgroupCrowdIds Getter
func (r TaobaoSimbaSerchcrowdBatchDeleteAPIRequest) GetAdgroupCrowdIds() []string {
	return r._adgroupCrowdIds
}

// SetSubNick is SubNick Setter
// 子帐号nick
func (r *TaobaoSimbaSerchcrowdBatchDeleteAPIRequest) SetSubNick(_subNick string) error {
	r._subNick = _subNick
	r.Set("sub_nick", _subNick)
	return nil
}

// GetSubNick SubNick Getter
func (r TaobaoSimbaSerchcrowdBatchDeleteAPIRequest) GetSubNick() string {
	return r._subNick
}

// SetNick is Nick Setter
// 被操作者的淘宝昵称
func (r *TaobaoSimbaSerchcrowdBatchDeleteAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaoSimbaSerchcrowdBatchDeleteAPIRequest) GetNick() string {
	return r._nick
}

var poolTaobaoSimbaSerchcrowdBatchDeleteAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSimbaSerchcrowdBatchDeleteRequest()
	},
}

// GetTaobaoSimbaSerchcrowdBatchDeleteRequest 从 sync.Pool 获取 TaobaoSimbaSerchcrowdBatchDeleteAPIRequest
func GetTaobaoSimbaSerchcrowdBatchDeleteAPIRequest() *TaobaoSimbaSerchcrowdBatchDeleteAPIRequest {
	return poolTaobaoSimbaSerchcrowdBatchDeleteAPIRequest.Get().(*TaobaoSimbaSerchcrowdBatchDeleteAPIRequest)
}

// ReleaseTaobaoSimbaSerchcrowdBatchDeleteAPIRequest 将 TaobaoSimbaSerchcrowdBatchDeleteAPIRequest 放入 sync.Pool
func ReleaseTaobaoSimbaSerchcrowdBatchDeleteAPIRequest(v *TaobaoSimbaSerchcrowdBatchDeleteAPIRequest) {
	v.Reset()
	poolTaobaoSimbaSerchcrowdBatchDeleteAPIRequest.Put(v)
}
