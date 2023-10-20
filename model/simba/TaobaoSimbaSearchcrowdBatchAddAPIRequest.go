package simba

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaSearchcrowdBatchAddAPIRequest 推广单元增加搜索人群 API请求
// taobao.simba.searchcrowd.batch.add
//
// 推广单元新增搜索人群
type TaobaoSimbaSearchcrowdBatchAddAPIRequest struct {
	model.Params
	// 被操作者的淘宝昵称
	_nick string
	// 新增人群信息,批量接口,入参为list,溢价(discount)范围为[105,400]
	_adgroupTargetingTags string
	// 推广单元id
	_adgroupId int64
}

// NewTaobaoSimbaSearchcrowdBatchAddRequest 初始化TaobaoSimbaSearchcrowdBatchAddAPIRequest对象
func NewTaobaoSimbaSearchcrowdBatchAddRequest() *TaobaoSimbaSearchcrowdBatchAddAPIRequest {
	return &TaobaoSimbaSearchcrowdBatchAddAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSimbaSearchcrowdBatchAddAPIRequest) Reset() {
	r._nick = ""
	r._adgroupTargetingTags = ""
	r._adgroupId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaSearchcrowdBatchAddAPIRequest) GetApiMethodName() string {
	return "taobao.simba.searchcrowd.batch.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaSearchcrowdBatchAddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSimbaSearchcrowdBatchAddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 被操作者的淘宝昵称
func (r *TaobaoSimbaSearchcrowdBatchAddAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaoSimbaSearchcrowdBatchAddAPIRequest) GetNick() string {
	return r._nick
}

// SetAdgroupTargetingTags is AdgroupTargetingTags Setter
// 新增人群信息,批量接口,入参为list,溢价(discount)范围为[105,400]
func (r *TaobaoSimbaSearchcrowdBatchAddAPIRequest) SetAdgroupTargetingTags(_adgroupTargetingTags string) error {
	r._adgroupTargetingTags = _adgroupTargetingTags
	r.Set("adgroup_targeting_tags", _adgroupTargetingTags)
	return nil
}

// GetAdgroupTargetingTags AdgroupTargetingTags Getter
func (r TaobaoSimbaSearchcrowdBatchAddAPIRequest) GetAdgroupTargetingTags() string {
	return r._adgroupTargetingTags
}

// SetAdgroupId is AdgroupId Setter
// 推广单元id
func (r *TaobaoSimbaSearchcrowdBatchAddAPIRequest) SetAdgroupId(_adgroupId int64) error {
	r._adgroupId = _adgroupId
	r.Set("adgroup_id", _adgroupId)
	return nil
}

// GetAdgroupId AdgroupId Getter
func (r TaobaoSimbaSearchcrowdBatchAddAPIRequest) GetAdgroupId() int64 {
	return r._adgroupId
}

var poolTaobaoSimbaSearchcrowdBatchAddAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSimbaSearchcrowdBatchAddRequest()
	},
}

// GetTaobaoSimbaSearchcrowdBatchAddRequest 从 sync.Pool 获取 TaobaoSimbaSearchcrowdBatchAddAPIRequest
func GetTaobaoSimbaSearchcrowdBatchAddAPIRequest() *TaobaoSimbaSearchcrowdBatchAddAPIRequest {
	return poolTaobaoSimbaSearchcrowdBatchAddAPIRequest.Get().(*TaobaoSimbaSearchcrowdBatchAddAPIRequest)
}

// ReleaseTaobaoSimbaSearchcrowdBatchAddAPIRequest 将 TaobaoSimbaSearchcrowdBatchAddAPIRequest 放入 sync.Pool
func ReleaseTaobaoSimbaSearchcrowdBatchAddAPIRequest(v *TaobaoSimbaSearchcrowdBatchAddAPIRequest) {
	v.Reset()
	poolTaobaoSimbaSearchcrowdBatchAddAPIRequest.Put(v)
}
