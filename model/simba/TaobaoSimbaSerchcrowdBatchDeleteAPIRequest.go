package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaSerchcrowdBatchDeleteAPIRequest 单品搜索人群批量取消溢价 API请求
// taobao.simba.serchcrowd.batch.delete
//
// 删除单品搜索人群溢价功能
type TaobaoSimbaSerchcrowdBatchDeleteAPIRequest struct {
	model.Params
	// 需要删除的人群id
	_adgroupCrowdIds []int64
	// 被操作者的淘宝昵称
	_nick string
	// 子帐号nick
	_subNick string
}

// NewTaobaoSimbaSerchcrowdBatchDeleteRequest 初始化TaobaoSimbaSerchcrowdBatchDeleteAPIRequest对象
func NewTaobaoSimbaSerchcrowdBatchDeleteRequest() *TaobaoSimbaSerchcrowdBatchDeleteAPIRequest {
	return &TaobaoSimbaSerchcrowdBatchDeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaSerchcrowdBatchDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.simba.serchcrowd.batch.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaSerchcrowdBatchDeleteAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetAdgroupCrowdIds is AdgroupCrowdIds Setter
// 需要删除的人群id
func (r *TaobaoSimbaSerchcrowdBatchDeleteAPIRequest) SetAdgroupCrowdIds(_adgroupCrowdIds []int64) error {
	r._adgroupCrowdIds = _adgroupCrowdIds
	r.Set("adgroup_crowd_ids", _adgroupCrowdIds)
	return nil
}

// GetAdgroupCrowdIds AdgroupCrowdIds Getter
func (r TaobaoSimbaSerchcrowdBatchDeleteAPIRequest) GetAdgroupCrowdIds() []int64 {
	return r._adgroupCrowdIds
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
