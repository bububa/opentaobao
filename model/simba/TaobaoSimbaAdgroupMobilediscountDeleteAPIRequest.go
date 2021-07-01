package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaAdgroupMobilediscountDeleteAPIRequest
批量删除adgroup的移动溢价 API请求
taobao.simba.adgroup.mobilediscount.delete

批量删除adgroup的移动溢价 */
type TaobaoSimbaAdgroupMobilediscountDeleteAPIRequest struct {
	model.Params
	// 昵称
	_nick string
	// adgroup主键数组（批量最多支持200个）
	_adgroupIds []int64
}

// NewTaobaoSimbaAdgroupMobilediscountDeleteRequest 初始化TaobaoSimbaAdgroupMobilediscountDeleteAPIRequest对象
func NewTaobaoSimbaAdgroupMobilediscountDeleteRequest() *TaobaoSimbaAdgroupMobilediscountDeleteAPIRequest {
	return &TaobaoSimbaAdgroupMobilediscountDeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaAdgroupMobilediscountDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.simba.adgroup.mobilediscount.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaAdgroupMobilediscountDeleteAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Nick Setter
// 昵称
func (r *TaobaoSimbaAdgroupMobilediscountDeleteAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// Get Nick Getter
func (r TaobaoSimbaAdgroupMobilediscountDeleteAPIRequest) GetNick() string {
	return r._nick
}

// Set is AdgroupIds Setter
// adgroup主键数组（批量最多支持200个）
func (r *TaobaoSimbaAdgroupMobilediscountDeleteAPIRequest) SetAdgroupIds(_adgroupIds []int64) error {
	r._adgroupIds = _adgroupIds
	r.Set("adgroup_ids", _adgroupIds)
	return nil
}

// Get AdgroupIds Getter
func (r TaobaoSimbaAdgroupMobilediscountDeleteAPIRequest) GetAdgroupIds() []int64 {
	return r._adgroupIds
}
